setup_file() {
    # Using a custom namespace to avoid side effects with other tests
    kubectl create ns qdrant-multinode-apikey
    # No --wait: non-zero peers may crash once if they start before the bootstrap peer is reachable,
    # which makes helm's wait fail, while the rollout recovers after the restart.
    helm upgrade --install qdrant charts/qdrant -n qdrant-multinode-apikey --set replicaCount=2 --set apiKey=foobar
    kubectl rollout status statefulset qdrant -n qdrant-multinode-apikey --timeout 3m
}

@test "all peers join the cluster" {
    for i in $(seq 1 30); do
        run kubectl exec -n default curl -- curl -s http://qdrant.qdrant-multinode-apikey:6333/cluster -H 'api-key: foobar' --fail-with-body
        peers=$(echo "${output}" | grep -o '"uri"' | wc -l)
        if [ "${peers}" -eq 2 ]; then
            return 0
        fi
        sleep 2
    done
    echo "Expected 2 peers, got ${peers}: ${output}"
    return 1
}

@test "replicated collection works with internal auth" {
    run kubectl exec -n default curl -- curl -s -X PUT http://qdrant.qdrant-multinode-apikey:6333/collections/test_collection -H 'api-key: foobar' -H 'Content-Type: application/json' --fail-with-body \
        --data-raw '{"vectors": {"size": 4, "distance": "Dot"}, "shard_number": 2, "replication_factor": 2}'
    [ $status -eq 0 ]
    [[ "${output}" =~ .*\"status\":\"ok\".* ]]

    # Wait until all shard replicas are active on both peers
    for i in $(seq 1 30); do
        run kubectl exec -n default curl -- curl -s http://qdrant-0.qdrant-headless.qdrant-multinode-apikey:6333/collections/test_collection/cluster -H 'api-key: foobar' --fail-with-body
        active=$(echo "${output}" | grep -o '"state":"Active"' | wc -l)
        if [ "${active}" -eq 4 ]; then
            break
        fi
        sleep 2
    done
    [ "${active}" -eq 4 ]

    # Writes are forwarded to the replica on the other peer via the internal p2p API
    run kubectl exec -n default curl -- curl -s -X PUT "http://qdrant.qdrant-multinode-apikey:6333/collections/test_collection/points?wait=true" -H 'api-key: foobar' -H 'Content-Type: application/json' --fail-with-body \
        --data-raw '{"points": [{"id": 1, "vector": [0.05, 0.61, 0.76, 0.74]}, {"id": 2, "vector": [0.19, 0.81, 0.75, 0.11]}, {"id": 3, "vector": [0.36, 0.55, 0.47, 0.94]}]}'
    [ $status -eq 0 ]
    [[ "${output}" =~ .*\"status\":\"completed\".* ]]

    # Every peer holds a full copy and must see all points
    for pod in qdrant-0 qdrant-1; do
        run kubectl exec -n default curl -- curl -s -X POST "http://${pod}.qdrant-headless.qdrant-multinode-apikey:6333/collections/test_collection/points/count" -H 'api-key: foobar' -H 'Content-Type: application/json' --fail-with-body \
            --data-raw '{"exact": true}'
        [ $status -eq 0 ]
        [[ "${output}" =~ .*\"count\":3.* ]]
    done
}

@test "no warnings in logs" {
    for pod in qdrant-0 qdrant-1; do
        run kubectl logs -n qdrant-multinode-apikey ${pod}
        [ $status -eq 0 ]
        # Ignore transient connection errors while the other peer is still starting up
        warnings=$(echo "${output}" | grep -E "WARN|ERROR" | grep -v "Failed to send message to .* error: transport error" || true)
        if [[ -n "${warnings}" ]]; then
            echo "Found warning output in ${pod}:"
            echo "${warnings}"
            return 1
        fi
    done
}

@test "helm test - multi node with api key" {
    run helm test qdrant -n qdrant-multinode-apikey --logs
    [ $status -eq 0 ]
}

teardown_file() {
    helm uninstall qdrant -n qdrant-multinode-apikey
    kubectl delete ns qdrant-multinode-apikey
}

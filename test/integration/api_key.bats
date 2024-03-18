setup_file() {
    helm upgrade --install qdrant charts/qdrant --set apiKey=foobar -n qdrant-helm-integration --wait
    kubectl rollout status statefulset qdrant -n qdrant-helm-integration
}

@test "api key authentication works" {
    run kubectl exec -n default curl -- curl -s http://qdrant.qdrant-helm-integration:6333/collections -H 'api-key: foobar' --fail-with-body
    [ $status -eq 0 ]
    [[ "${output}" =~ .*\"status\":\"ok\".* ]]
}

@test "api key authentication fails with key" {
    run kubectl exec -n default curl -- curl -s http://qdrant.qdrant-helm-integration:6333/collections
    [ "${output}" = "Write access denied" ]
}

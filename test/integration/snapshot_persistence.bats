QDRANT_URL="http://qdrant.qdrant-helm-integration:6333"

setup_file() {
    # Must delete since this test modifies immutable fields of the STS
    helm delete qdrant -n qdrant-helm-integration || true
    kubectl wait --for=delete pod/qdrant-0 -n qdrant-helm-integration --timeout=300s
    helm upgrade --install qdrant charts/qdrant -n qdrant-helm-integration -f test/integration/assets/snapshot-persistence-values.yaml --wait
    kubectl rollout status statefulset qdrant -n qdrant-helm-integration
}


@test "snapshot persistence works" {
    # Create the test data and create a snapshot
    helm test qdrant -n qdrant-helm-integration --logs
    SNAPSHOT_NAME=$(kubectl exec -n default curl -- curl -X POST "$QDRANT_URL/collections/test_collection/snapshots" | grep -o 'test_collection.*\.snapshot')

    # Delete the collection
    kubectl exec -n default curl -- curl -X DELETE "$QDRANT_URL/collections/test_collection?wait"

    # Restart the cluster
    kubectl rollout restart statefulset qdrant -n qdrant-helm-integration
    kubectl rollout status statefulset qdrant -n qdrant-helm-integration
    # Ensure the snapshot survived the reboot by trying to restore from it
    helm upgrade qdrant charts/qdrant -n qdrant-helm-integration -f test/integration/assets/snapshot-values.yaml -f test/integration/assets/snapshot-persistence-values.yaml --set snapshotRestoration.snapshots="{/qdrant/snapshots/test_collection/$SNAPSHOT_NAME\:test_collection}" --wait
    kubectl rollout status statefulset qdrant -n qdrant-helm-integration

    # Ensure the restoration worked
    run kubectl exec -n default curl -- curl -s $QDRANT_URL/collections/test_collection/points/6 --fail-with-body
    [ $status -eq 0 ]
    [[ "${output}" =~ .*\"Mumbai\".* ]]
}

teardown_file() {
    helm delete qdrant -n qdrant-helm-integration || true
    kubectl wait --for=delete pod/qdrant-0 -n qdrant-helm-integration --timeout=300s
}

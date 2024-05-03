# Inspired by https://github.com/qdrant/qdrant/blob/v1.6.1/tests/snapshots/snapshots-recovery.sh

QDRANT_URL="http://qdrant.qdrant-helm-integration:6333"

setup_file() {
    # Create the PVC, then helm upgrade with no snapshots so that the PVC gets mounted
    kubectl apply -f test/integration/assets/snapshot-pvc.yaml -n qdrant-helm-integration
    helm upgrade --install qdrant charts/qdrant -n qdrant-helm-integration -f test/integration/assets/snapshot-values.yaml --set snapshotRestoration.snapshots=null --wait
    kubectl rollout status statefulset qdrant -n qdrant-helm-integration
}


@test "snapshot restoration works" {
    # Create the test data and create a snapshot
    helm test qdrant -n qdrant-helm-integration --logs
    SNAPSHOT_NAME=$(kubectl exec -n default curl -- curl -X POST "$QDRANT_URL/collections/test_collection/snapshots" | grep -o 'test_collection.*\.snapshot')

    # Delete the collection, then restore
    kubectl exec -n default curl -- curl -X DELETE "$QDRANT_URL/collections/test_collection?wait"
    helm upgrade --install qdrant charts/qdrant -n qdrant-helm-integration -f test/integration/assets/snapshot-values.yaml --set snapshotRestoration.snapshots="{/qdrant/snapshot-restoration/test_collection/$SNAPSHOT_NAME\:test_collection}" --wait
    kubectl rollout status statefulset qdrant -n qdrant-helm-integration

    run kubectl exec -n default curl -- curl -s $QDRANT_URL/collections/test_collection/points/6 --fail-with-body
    [ $status -eq 0 ]
    [[ "${output}" =~ .*\"Mumbai\".* ]]
}

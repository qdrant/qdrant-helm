setup_file() {
    helm upgrade --install qdrant charts/qdrant -n qdrant-helm-integration --wait --set image.useUnprivilegedImage=true
    kubectl rollout status statefulset qdrant -n qdrant-helm-integration
}

@test "helm test - with unprivileged image" {
    run helm test qdrant -n qdrant-helm-integration --logs
    [ $status -eq 0 ]
}

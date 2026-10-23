setup_file() {
    helm upgrade --install qdrant charts/qdrant -n qdrant-helm-integration --wait --set chartTests.dbInteraction.containerSecurityContext.readOnlyRootFilesystem=true
}

@test "helm test - with read-only root filesystem" {
    run helm test qdrant -n qdrant-helm-integration --logs
    [ $status -eq 0 ]
}

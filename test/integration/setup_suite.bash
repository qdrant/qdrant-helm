setup_suite() {
    kind create cluster -n qdrant-helm-integration
}

teardown_suite() {
    kind delete cluster -n qdrant-helm-integration
}

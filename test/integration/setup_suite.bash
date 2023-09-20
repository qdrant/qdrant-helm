setup_suite() {
    kind create cluster -n qdrant-helm-integration
    kubectl create serviceaccount default -n default
    kubectl -n default run curl --image=curlimages/curl --command -- sh -c "sleep 3600"
    kubectl wait --for=condition=Ready pod/curl -n default --timeout=300s
    kubectl create namespace qdrant-helm-integration
    kubectl create serviceaccount default -n qdrant-helm-integration || true
}

teardown_suite() {
    kind delete cluster -n qdrant-helm-integration
}

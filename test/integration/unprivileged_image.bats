setup_file() {
    kubectl create namespace qdrant-helm-integration
    kubectl create serviceaccount default -n qdrant-helm-integration || true
    helm install qdrant charts/qdrant -n qdrant-helm-integration --wait --set image.useUnprivilegedImage=true
    kubectl rollout status statefulset qdrant -n qdrant-helm-integration
    sleep 10
}

teardown_file() {
    helm uninstall qdrant -n qdrant-helm-integration
    kubectl delete serviceaccount default -n qdrant-helm-integration
    kubectl delete namespace qdrant-helm-integration
}

@test "helm test - with unprivileged image" {
    run helm test qdrant -n qdrant-helm-integration --logs
    [ $status -eq 0 ]
}

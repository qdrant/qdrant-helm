setup_file() {
    kubectl create namespace qdrant-helm-integration
    kubectl create serviceaccount default -n qdrant-helm-integration || true
    helm install qdrant charts/qdrant --set apiKey=foobar -n qdrant-helm-integration --wait
    kubectl rollout status statefulset qdrant -n qdrant-helm-integration
}

teardown_file() {
    helm uninstall qdrant -n qdrant-helm-integration
    kubectl delete serviceaccount default -n qdrant-helm-integration
    kubectl delete namespace qdrant-helm-integration
}

@test "api key authentication works" {
    run kubectl exec -it -n default curl -- curl http://qdrant.qdrant-helm-integration:6333/collections -H 'api-key: foobar'
    [ $status -eq 0 ]
    [ "${output}" != "Invalid api-key" ]
}

@test "api key authentication fails with key" {
    run kubectl exec -it -n default curl -- curl http://qdrant.qdrant-helm-integration:6333/collections
    [ "${output}" = "Invalid api-key" ]
}

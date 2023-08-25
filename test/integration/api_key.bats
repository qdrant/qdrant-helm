setup_file() {
    kubectl create namespace qdrant-helm-integration
    kubectl create serviceaccount default -n qdrant-helm-integration || true
    helm install qdrant charts/qdrant --set apiKey=foobar -n qdrant-helm-integration
}

teardown_file() {
    helm uninstall qdrant -n qdrant-helm-integration
    kubectl delete serviceaccount default -n qdrant-helm-integration
    kubectl delete namespace qdrant-helm-integration
}

@test "api key authentication works" {
    run kubectl -n qdrant-helm-integration run --rm -i -t api-key-test-works --image=curlimages/curl -- http://qdrant:6333/collections -H 'api-key: foobar'
    [ $status -eq 0 ]
}

@test "api key authentication fails with key" {
    run kubectl -n qdrant-helm-integration run --rm -i -t api-key-test-fails --image=curlimages/curl -- http://qdrant:6333/collections
    [ $status -ne 0 ]
}

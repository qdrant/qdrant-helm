setup_file() {
    kubectl create namespace qdrant-helm-integration
    kubectl create serviceaccount default -n qdrant-helm-integration || true
    helm install qdrant charts/qdrant --set apiKey=true -n qdrant-helm-integration
}

teardown_file() {
    helm uninstall qdrant -n qdrant-helm-integration
    kubectl delete serviceaccount default -n qdrant-helm-integration
    kubectl delete namespace qdrant-helm-integration
}

@test "random api key works" {
    apiKey=$(kubectl -n qdrant-helm-integration get secret qdrant-apikey -o jsonpath="{.data.api-key}" | base64 --decode)
    run kubectl -n qdrant-helm-integration run --rm -i -t api-key-test-works --image=curlimages/curl -- http://qdrant:6333/collections -H "api-key: ${apiKey}"
    [ $status -eq 0 ]
}

@test "random api key stays the same after upgrade" {
    apiKeyBeforeUpgrade=$(kubectl -n qdrant-helm-integration get secret qdrant-apikey -o jsonpath="{.data.api-key}" | base64 --decode)
    helm upgrade qdrant charts/qdrant --set apiKey=true -n qdrant-helm-integration
    apiKeyAfterUpgrade=$(kubectl -n qdrant-helm-integration get secret qdrant-apikey -o jsonpath="{.data.api-key}" | base64 --decode)
    [ "${apiKeyBeforeUpgrade}" = "${apiKeyAfterUpgrade}" ]
}

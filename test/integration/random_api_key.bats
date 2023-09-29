setup_file() {
    helm upgrade --install qdrant charts/qdrant --set apiKey=true -n qdrant-helm-integration --wait
    kubectl rollout status statefulset qdrant -n qdrant-helm-integration
}

@test "random api key works" {
    apiKey=$(kubectl -n qdrant-helm-integration get secret qdrant-apikey -o jsonpath="{.data.api-key}" | base64 --decode)
    [ ${#apiKey} -eq 32 ]
    run kubectl exec -n default curl -- curl -s http://qdrant.qdrant-helm-integration:6333/collections -H "api-key: ${apiKey}" --fail-with-body
    [ $status -eq 0 ]
    [[ "${output}" =~ .*\"status\":\"ok\".* ]]
}

@test "random api key stays the same after upgrade" {
    apiKeyBeforeUpgrade=$(kubectl -n qdrant-helm-integration get secret qdrant-apikey -o jsonpath="{.data.api-key}" | base64 --decode)
    helm upgrade qdrant charts/qdrant --set apiKey=true -n qdrant-helm-integration --wait
    kubectl rollout status statefulset qdrant -n qdrant-helm-integration
    apiKeyAfterUpgrade=$(kubectl -n qdrant-helm-integration get secret qdrant-apikey -o jsonpath="{.data.api-key}" | base64 --decode)
    [ "${apiKeyBeforeUpgrade}" = "${apiKeyAfterUpgrade}" ]
    [ ${#apiKeyBeforeUpgrade} -eq 32 ]
    [ ${#apiKeyAfterUpgrade} -eq 32 ]
}

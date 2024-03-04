setup_file() {
    helm upgrade --install qdrant charts/qdrant --set readOnlyApiKey=barbaz -n qdrant-helm-integration --wait
    kubectl rollout status statefulset qdrant -n qdrant-helm-integration
}

@test "api key authentication works" {
    run kubectl exec -n default curl -- curl -s http://qdrant.qdrant-helm-integration:6333/collections -H 'api-key: barbaz' --fail-with-body
    [ $status -eq 0 ]
    [[ "${output}" =~ .*\"status\":\"ok\".* ]]
}

@test "api key authentication fails with key" {
    run kubectl exec -n default curl -- curl -s http://qdrant.qdrant-helm-integration:6333/collections
    [ "${output}" = "Invalid api-key" ]
}

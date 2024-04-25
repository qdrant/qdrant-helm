setup_file() {
    helm upgrade --install qdrant charts/qdrant --set readOnlyApiKey=barbaz -n qdrant-helm-integration --wait
    kubectl rollout status statefulset qdrant -n qdrant-helm-integration
}

@test "read-only api key authentication works" {
    run kubectl exec -n default curl -- curl -s http://qdrant.qdrant-helm-integration:6333/collections -H 'api-key: barbaz' --fail-with-body
    [ $status -eq 0 ]
    [[ "${output}" =~ .*\"status\":\"ok\".* ]]
}

@test "read-only api key authentication fails with no key" {
    run kubectl exec -n default curl -- curl -s -w " - %{response_code}" http://qdrant.qdrant-helm-integration:6333/collections
    [ "${output}" = "Must provide an API key or an Authorization bearer token - 401" ]
}

@test "read-only api key authentication fails with wrong key" {
    run kubectl exec -n default curl -- curl -s -w " - %{response_code}" http://qdrant.qdrant-helm-integration:6333/collections -H 'api-key: invalid'
    [ "${output}" = "Invalid API key or JWT - 401" ]
}

setup_file() {
    kubectl -n qdrant-helm-integration create secret generic qdrant-external-apikey --from-literal=apiKey=test-api-key --from-literal=readOnlyApiKey=test-read-only-api-key
    helm upgrade --install qdrant charts/qdrant --set apiKey.valueFrom.secretKeyRef.name=qdrant-external-apikey,apiKey.valueFrom.secretKeyRef.key=apiKey,readOnlyApiKey.valueFrom.secretKeyRef.name=qdrant-external-apikey,readOnlyApiKey.valueFrom.secretKeyRef.key=readOnlyApiKey -n qdrant-helm-integration --wait
    kubectl rollout status statefulset qdrant -n qdrant-helm-integration
}

@test "external api key works" {
    run kubectl exec -n default curl -- curl -s http://qdrant.qdrant-helm-integration:6333/collections -H "api-key: test-api-key" --fail-with-body
    [ $status -eq 0 ]
    [[ "${output}" =~ .*\"status\":\"ok\".* ]]
}

@test "external read only api key works" {
    run kubectl exec -n default curl -- curl -s http://qdrant.qdrant-helm-integration:6333/collections -H "api-key: test-read-only-api-key" --fail-with-body
    [ $status -eq 0 ]
    [[ "${output}" =~ .*\"status\":\"ok\".* ]]
}
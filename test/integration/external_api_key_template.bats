setup_file() {
    # Using a custom namespace to avoid side effects with other tests
    kubectl create ns qdrant-external-apikey-template
    kubectl -n qdrant-external-apikey-template create secret generic qdrant-external-apikey --from-literal=apiKey=test-api-key --from-literal=readOnlyApiKey=test-read-only-api-key
    # Render without cluster access and apply the output, like ArgoCD and other GitOps tools do.
    # Helm's lookup function returns nothing in this mode.
    helm template qdrant charts/qdrant -n qdrant-external-apikey-template --skip-tests \
        --set apiKey.valueFrom.secretKeyRef.name=qdrant-external-apikey,apiKey.valueFrom.secretKeyRef.key=apiKey,readOnlyApiKey.valueFrom.secretKeyRef.name=qdrant-external-apikey,readOnlyApiKey.valueFrom.secretKeyRef.key=readOnlyApiKey \
        | kubectl apply -n qdrant-external-apikey-template -f -
    kubectl rollout status statefulset qdrant -n qdrant-external-apikey-template --timeout 3m
}

@test "external api key works when rendered with helm template" {
    run kubectl exec -n default curl -- curl -s http://qdrant.qdrant-external-apikey-template:6333/collections -H "api-key: test-api-key" --fail-with-body
    [ $status -eq 0 ]
    [[ "${output}" =~ .*\"status\":\"ok\".* ]]
}

@test "external read only api key works when rendered with helm template" {
    run kubectl exec -n default curl -- curl -s http://qdrant.qdrant-external-apikey-template:6333/collections -H "api-key: test-read-only-api-key" --fail-with-body
    [ $status -eq 0 ]
    [[ "${output}" =~ .*\"status\":\"ok\".* ]]
}

@test "external api key authentication fails with no key when rendered with helm template" {
    run kubectl exec -n default curl -- curl -s -w " - %{response_code}" http://qdrant.qdrant-external-apikey-template:6333/collections
    [ "${output}" = "Must provide an API key or an Authorization bearer token - 401" ]
}

@test "external api key authentication fails with wrong key when rendered with helm template" {
    run kubectl exec -n default curl -- curl -s -w " - %{response_code}" http://qdrant.qdrant-external-apikey-template:6333/collections -H 'api-key: invalid'
    [ "${output}" = "Invalid API key or JWT - 401" ]
}

teardown_file() {
    kubectl delete ns qdrant-external-apikey-template
}

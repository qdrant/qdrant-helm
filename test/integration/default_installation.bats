setup_file() {
    helm upgrade --install qdrant charts/qdrant -n qdrant-helm-integration --wait
    kubectl rollout status statefulset qdrant -n qdrant-helm-integration
}

@test "helm test - default values" {
    run helm test qdrant -n qdrant-helm-integration --logs
    [ $status -eq 0 ]
}

@test "no startup warnings in logs" {
    run kubectl logs -n qdrant-helm-integration qdrant-0
    [ $status -eq 0 ]
    [[ "${output}" =~ .*INFO.* ]]
    [[ ! "${output}" =~ .*WARN.* ]]
}

@test "no startup errors in logs" {
    run kubectl logs -n qdrant-helm-integration qdrant-0
    [ $status -eq 0 ]
    [[ "${output}" =~ .*INFO.* ]]
    [[ ! "${output}" =~ .*ERR.* ]]
}
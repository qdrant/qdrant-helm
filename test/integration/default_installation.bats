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
    warnings=$(echo "${output}" | grep "WARN" | grep -v "WARN.*qdrant::settings: Running in distributed mode with an API key configured" || true)
    if [[ -n "${warnings}" ]]; then
        echo "Found warning output:"
        echo "${warnings}"
        return 1
    fi
}

@test "no startup errors in logs" {
    run kubectl logs -n qdrant-helm-integration qdrant-0
    [ $status -eq 0 ]
    [[ "${output}" =~ .*INFO.* ]]
    [[ ! "${output}" =~ .*ERR.* ]]
}

@test "SIGTERM signals are propagated to qdrant" {
    run kubectl rollout restart sts/qdrant -n qdrant-helm-integration
    [ $status -eq 0 ]
    # If signals aren't working, this will take >30 seconds and time out
    run kubectl rollout status statefulset qdrant -n qdrant-helm-integration --timeout=15s
    [ $status -eq 0 ]
}

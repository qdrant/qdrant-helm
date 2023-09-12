setup_file() {
    kubectl create namespace qdrant-helm-integration
    kubectl create serviceaccount default -n qdrant-helm-integration || true
    helm install qdrant charts/qdrant -n qdrant-helm-integration --wait
    kubectl rollout status statefulset qdrant -n qdrant-helm-integration
    sleep 5
}

teardown_file() {
    helm uninstall qdrant -n qdrant-helm-integration
    kubectl delete serviceaccount default -n qdrant-helm-integration
    kubectl delete namespace qdrant-helm-integration
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
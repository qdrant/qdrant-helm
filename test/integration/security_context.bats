setup_file() {
    kubectl create namespace qdrant-helm-integration
    kubectl create serviceaccount default -n qdrant-helm-integration || true
}

teardown_file() {
    helm uninstall qdrant -n qdrant-helm-integration
    kubectl delete serviceaccount default -n qdrant-helm-integration
    kubectl delete namespace qdrant-helm-integration
}

@test "update without security context to security context corrects file permissions" {
    helm install qdrant charts/qdrant --set-json 'podSecurityContext=false,containerSecurityContext=false' -n qdrant-helm-integration --wait
    kubectl rollout status statefulset qdrant -n qdrant-helm-integration
    user=$(kubectl exec qdrant-0 -n qdrant-helm-integration -- id -u)
    [ "${user}" = "0" ]
    user=$(kubectl exec qdrant-0 -n qdrant-helm-integration -- whoami)
    [ "${user}" = "root" ]
    helm upgrade --reset-values qdrant charts/qdrant -n qdrant-helm-integration --wait
    kubectl rollout status statefulset qdrant -n qdrant-helm-integration
    user=$(kubectl exec qdrant-0 -n qdrant-helm-integration -- id -u)
    [ "${user}" = "1000" ]

}
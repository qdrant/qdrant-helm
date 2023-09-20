@test "update without security context to security context corrects file permissions" {
    helm upgrade --install qdrant charts/qdrant --set-json 'podSecurityContext=false,containerSecurityContext=false' -n qdrant-helm-integration --wait
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
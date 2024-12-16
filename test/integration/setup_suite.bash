setup_suite() {
    kind create cluster -n qdrant-helm-integration
    kubectl create serviceaccount default -n default || true
    kubectl -n default run curl --image=docker.io/curlimages/curl --command -- sh -c '
    echo "connect-timeout = 5" > $HOME/.curlrc;
    echo "retry = 60" >> $HOME/.curlrc;
    echo "retry-delay = 5" >> $HOME/.curlrc;
    echo "retry-all-errors" >> $HOME/.curlrc;
    sleep 3600'
    kubectl wait --for=condition=Ready pod/curl -n default --timeout=300s
    kubectl create namespace qdrant-helm-integration
    kubectl create serviceaccount default -n qdrant-helm-integration || true
}

teardown_suite() {
    kind delete cluster -n qdrant-helm-integration
}

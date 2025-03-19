setup_file() {
    rm test/integration/assets/ca.* || true
    rm test/integration/assets/tls.* || true
    openssl genrsa -des3 -passout pass:insecure -out test/integration/assets/ca.key 2048
    openssl req -x509 -passin pass:insecure -new -nodes -key test/integration/assets/ca.key -sha256 -days 1825 -out test/integration/assets/ca.pem -subj "/C=DE/ST=Berlin/L=Berlin/O=Qdrant/OU=Cloud/CN=qdrant"
    openssl genrsa -out test/integration/assets/tls.key 2048
    openssl req -new -key test/integration/assets/tls.key -out test/integration/assets/tls.csr -subj "/C=DE/ST=Berlin/L=Berlin/O=Qdrant/OU=Cloud/CN=qdrant.qdrant-helm-integration"
    openssl x509 -req -passin pass:insecure -in test/integration/assets/tls.csr -CA test/integration/assets/ca.pem -CAkey test/integration/assets/ca.key -CAcreateserial -extensions v3_req -extfile test/integration/assets/cert.cfg -out test/integration/assets/tls.crt -days 825 -sha256

    helm del qdrant -n qdrant-helm-integration || true
    kubectl -n qdrant-helm-integration delete secret test-tls || true
    kubectl -n qdrant-helm-integration create secret generic test-tls --from-file=tls.key="test/integration/assets/tls.key" --from-file=tls.crt="test/integration/assets/tls.crt" --from-file=ca.pem="test/integration/assets/ca.pem"
    helm upgrade --install qdrant charts/qdrant -n qdrant-helm-integration --wait -f test/integration/assets/tls-p2p-values.yaml
    kubectl rollout status statefulset qdrant -n qdrant-helm-integration
}

@test "Connection with https" {
    run kubectl exec -n default curl -- curl -k -s https://qdrant.qdrant-helm-integration:6333/collections --fail-with-body
    [ $status -eq 0 ]
    [[ "${output}" =~ .*\"status\":\"ok\".* ]]
}

@test "helm test - with https" {
    run helm test qdrant -n qdrant-helm-integration --logs
    [ $status -eq 0 ]
}

teardown_file() {
    helm del qdrant -n qdrant-helm-integration || true
    kubectl -n qdrant-helm-integration delete secret test-tls || true
}

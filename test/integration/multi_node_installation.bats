@test "Rolling update" {
  run kubectl rollout restart sts/qdrant -n distributed-update-test
  [ $status -eq 0 ]
}

setup() {
  # Using a custom namespace to avoid side effects with other tests
  kubectl create ns distributed-update-test
  helm upgrade --install qdrant charts/qdrant -n distributed-update-test --wait --timeout 3m --set replicaCount=3 --set image.tag=v1.7.3
  kubectl rollout status statefulset qdrant -n distributed-update-test
}

teardown() {
  helm uninstall qdrant -n distributed-update-test
  kubectl delete ns distributed-update-test
}

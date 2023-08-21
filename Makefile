lint:
	helm lint charts/qdrant

test-unit:
	cd test && go test -v

test-integration:
	kind create cluster -n qdrant-helm-integration
	helm install qdrant ./charts/qdrant --wait
	kubectl get pods
	helm test qdrant
	kind delete cluster -n qdrant-helm-integration

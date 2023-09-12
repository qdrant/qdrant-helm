lint:
	helm lint charts/qdrant

test-unit:
	go test -v ./test

test-integration:
	bats test/integration --verbose-run --show-output-of-passing-tests

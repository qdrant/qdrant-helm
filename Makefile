lint:
	helm lint charts/qdrant
	gofmt -w -s ./test
	golangci-lint run ./test

test-unit:
	go test -v ./test

test-integration:
	bats test/integration --verbose-run --show-output-of-passing-tests

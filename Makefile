lint:
	helm lint charts/qdrant

test-unit:
	go test ./test

test-integration:
	bats test/integration

test-unit-lint:
	gofmt -w -s ./test
	golangci-lint run ./test

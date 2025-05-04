.PHONY: run
run: lint test
	go run cmd/main.go

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test: lint
	set -o pipefail; go test ./... | grep -v "no test files"

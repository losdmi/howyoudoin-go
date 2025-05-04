.PHONY: run
run: lint test
	go run cmd/main.go

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test: lint
	set -o pipefail; go test ./... | grep -v "no test files"

.PHONY: build_for_windows
build_for_windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build \
		-trimpath \
		-ldflags="-s -w -extldflags '-static'" \
		-o target/windows/howyoudoin.exe \
		cmd/main.go

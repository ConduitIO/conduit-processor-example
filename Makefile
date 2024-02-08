.PHONY: default
default: fmt lint test

.PHONY: test
test:
	go test $(GOTEST_FLAGS) -race ./...

.PHONY: build
build:
	GOOS=wasip1 GOARCH=wasm go build -ldflags "-X 'github.com/conduitio/conduit-processor-example.version=${VERSION}'" -o conduit-processor-example cmd/processor/main.go

.PHONY: lint
lint:
	golangci-lint run

.PHONY: fmt
fmt:
	gofumpt -l -w .

.PHONY: install-tools
install-tools:
	@echo Installing tools from tools.go
	@go list -e -f '{{ join .Imports "\n" }}' tools.go | xargs -tI % go install %
	@go mod tidy

.PHONY: generate
generate:
	go generate ./...

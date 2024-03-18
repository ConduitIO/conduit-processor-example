build:
	GOOS=wasip1 GOARCH=wasm go build -o complex.wasm cmd/processor/main.go

.PHONY: lint
lint:
	golangci-lint run
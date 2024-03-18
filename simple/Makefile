build:
	GOOS=wasip1 GOARCH=wasm go build -o simple.wasm cmd/processor/main.go


.PHONY: lint
lint:
	golangci-lint run
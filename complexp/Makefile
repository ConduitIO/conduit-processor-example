build:
	GOOS=wasip1 GOARCH=wasm go build -o complex.wasm cmd/processor/main.go
test:
	go test $(GOTEST_FLAGS) -race ./...
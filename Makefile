test:
	go test -v -cover -covermode=atomic ./internal/usecase/...

unittest:
	go test -short  ./internal/usecase/...

lint:
	golangci-lint run
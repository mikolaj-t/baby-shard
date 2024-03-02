install:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/tools/cmd/goimports@latest

fmt:
	goimports -w .

lint:
	golangci-lint run

run:
	go run ./cmd/main.go

run-test-servers:
	go run ./example/ping_server.go 1000
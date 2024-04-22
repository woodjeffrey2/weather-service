
.PHONY: test build run mocks
test:
	go test ./...

build:
	go build -o bin/server src/server/main.go

run:
	go run src/server/main.go

mocks:
	mockery

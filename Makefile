
.PHONY: test build run run-docker mocks
test:
	go test ./...

build:
	go build -o bin/server src/server/main.go

run:
	go run src/server/main.go

run-docker:
	docker compose up

mocks:
	mockery

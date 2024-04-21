
.PHONY: test build run
test:
	go test ./...

build:
	go build -o bin/server src/server/main.go

run:
	go run src/server/main.go


.PHONY: test build
test:
	go test ./...

build:
	go build -o bin/server src/server/main.go

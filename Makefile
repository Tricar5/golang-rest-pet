.PHONY: build

build:
	go build -v ./cmd/apiserver

run:
	./apiserver

format:
	 go fmt ./...

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build
ifneq (,$(wildcard .env))
    include .env
    export
endif

.PHONY: dev build test lint fmt tidy docs docker-build migrate

dev:
	go run ./cmd

build:
	CGO_ENABLED=0 go build -o home-env-server ./cmd

preview:
	./home-env-server

test:
	go test ./...

lint:
	go vet ./...

fmt:
	gofmt -w .
	go run github.com/swaggo/swag/cmd/swag@latest fmt

tidy:
	go mod tidy

docs:
	go run github.com/swaggo/swag/cmd/swag@latest init -g cmd/main.go -o docs --outputTypes json,yaml

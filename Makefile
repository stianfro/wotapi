#!make
include .env
export $(shell sed 's/=.*//' .env)

init:
	go install github.com/cosmtrek/air@latest
	go install github.com/securego/gosec/v2/cmd/gosec@latest
	go install gotest.tools/gotestsum@latest
	brew install golangci-lint

deps:
	go get
	go mod tidy
	go mod vendor

swagger:
	swag i
	git add docs
	git commit -m "docs(swagger): update api documentation"

build:
	go build -o bin/main main.go

test:
	gotestsum --format testname

image:
	docker build .

lint:
	golangci-lint run

gosec:
	gosec ./...

govulncheck:
	govulncheck ./...

security: gosec govulncheck

clean:
	go clean
	rm -f bin/*

check: build test lint security clean

run:
	air

run/docker:
	docker-compose up

get/health:
	http http://localhost:${PORT}/api/v1/healthz
get/manga:
	http http://localhost:${PORT}/api/v1/manga
get/volume:
	http http://localhost:${PORT}/api/v1/manga/volume

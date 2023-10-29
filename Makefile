.PHONY: build
build:
	go build -v ./cmd/url-shortener-service

.PHONY: test
test:
	go test -v -race -count=1 -timeout 5s ./...

.DEFAULT_GOAL := build
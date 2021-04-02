RELEASES:= \
	_bin/mdgo-linux-amd64 \
	_bin/mdgo-darwin-amd64

.PHONY: all lint test install serve build build-release
.FORCE:

all: test lint

lint:
	golangci-lint run ./...

test:
	go test -v -race ./...	

install:
	go install ./cmd/mdgo

run-example:
	DEBUG=1 go run ./cmd/mdgo-example

build-release: $(RELEASES)

_bin/mdgo-linux-amd64: .FORCE
_bin/mdgo-linux-amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
		go build -o $@ ./cmd/mdgo

_bin/mdgo-darwin-amd64: .FORCE
_bin/mdgo-darwin-amd64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 \
		go build -o $@ ./cmd/mdgo

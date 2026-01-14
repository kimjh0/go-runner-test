.PHONY: all build test clean fmt fmt-check

BINARY_NAME := go-runner-test

all: build

build:
	go build -o $(BINARY_NAME) .

test:
	go test -v ./...

clean:
	rm -f $(BINARY_NAME)

fmt:
	go fmt ./...

fmt-check:
	@test -z "$$(gofmt -l .)" || (echo "Run 'make fmt' to format code" && exit 1)

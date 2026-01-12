# Load environment variables from .env file if it exists
ifneq (,$(wildcard .env))
    include .env
    export
endif

# Default target
all: lint test
.PHONY: all

# Run golangci-lint
lint:
	golangci-lint run ./recombee/...
.PHONY: lint

# Run unit-tests
test:
	go test -short -v ./test
.PHONY: test

# Run integration tests
test-integration:
	DB_ID=$(DB_ID) PRIVATE_TOKEN=$(PRIVATE_TOKEN) go test -v ./test
.PHONY: test-integration
SHELL=/bin/bash
.DEFAULT_GOAL := help

# https://gist.github.com/tadashi-aikawa/da73d277a3c1ec6767ed48d1335900f3
.PHONY: $(shell grep --no-filename -E '^[a-zA-Z0-9_-]+:' $(MAKEFILE_LIST) | sed 's/://')

BINARY_PATH = ./zz_example/main

# Phony Targets

clean: ## Clean
	rm -f $(BINARY_PATH)

build: fmt clean ## Build
	go build -o $(BINARY_PATH) .

run: fmt ## Run
	go run .

clean-test: cache test vet ## Clean and Test

test: fmt ## Test
	go test ./parser

test-all: fmt ## Test all
	go test $$(go list ./... | grep -v zz_example)

vet: ## Vet
	go vet $$(go list ./... | grep -v zz_example)

fmt: ## Format
	go fmt ./...

cache: ## Cache clear
	go clean -cache -testcache

# https://postd.cc/auto-documented-makefile/
help: ## Show help
	@grep --no-filename -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-40s\033[0m %s\n", $$1, $$2}'

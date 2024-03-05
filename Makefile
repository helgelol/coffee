ifeq ($(OS),Windows_NT)
	# Do nothing, to not crash in Windows because 'uname' doesn't exist
else
	ifeq ($(shell uname -s),Linux)
		SHELL := /bin/bash
	endif
endif

.DEFAULT_GOAL := help

##@ General

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Backend

.PHONY: env
env: ## Create a new .env file from .env.example
	docker compose up

.PHONY: dev
dev: ## Start local development server
	docker compose up

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

.PHONY: test
test: fmt vet ## Run tests.
	go test ./...

##@ Build

.PHONY: build
build: test ## Build API
	go build api/main.go

.PHONY: docker-build
docker-build: ## Build docker images (WIP)
	docker compose build

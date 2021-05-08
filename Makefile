
# Versioning information
GIT_COMMIT := $(shell git rev-parse HEAD)
GIT_BRANCH := $(shell git name-rev --name-only HEAD | sed "s/~.*//")

## Gets the current tag name or commit SHA
VERSION ?= $(shell git describe --tags ${COMMIT} 2> /dev/null || echo "$(GIT_COMMIT)")

## Gets the -ldflags for the go build command, this lets us set the version number in the binary
ROOT := github.com/yukitsune/chameleon
LD_FLAGS := -X '$(ROOT).Version=$(VERSION)'

## Whether the repo has uncommitted changes
GIT_DIRTY := false
ifneq ($(shell git status -s),)
	GIT_DIRTY=true
endif

## Common docker build args
DOCKER_BUILD_ARGS := \
	--build-arg GIT_COMMIT="$(GIT_COMMIT)" \
	--build-arg GIT_BRANCH="$(GIT_BRANCH)" \
	--build-arg GIT_DIRTY="$(GIT_DIRTY)" \
	--build-arg VERSION="$(VERSION)" \

.DEFAULT_GOAL := help

.PHONY: help
help: ## Shows this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## Builds all programs and places their binaries in the bin/ directory
	mkdir -p bin
	go build -ldflags="$(LD_FLAGS)" -o ./bin/  ./cmd/...

.PHONY: clean
clean: ## Removes the bin/ directory
	rm -rf bin

.PHONY: build-containers
build-containers: build-mtd-container build-api-container ## Builds a docker container for all programs

.PHONY: build-mtd-container
build-mtd-container: ## Builds the docker container for the MTD
	docker build \
		-t chameleon-mtd \
		-f build/package/chameleon-mtd/Dockerfile \
		$(DOCKER_BUILD_ARGS) \
		.

.PHONY: build-api-container
build-api-container: ## Builds the docker container for the REST API server
	docker build \
		-t chameleon-api-server \
		-f build/package/chameleon-api-server/Dockerfile \
		$(DOCKER_BUILD_ARGS) \
		.

.PHONY: compose
compose: ## Runs docker compose
	docker compose -f ./deployments/docker-compose.yml up

.PHONY: compose-fresh
compose-fresh: ## Rebuilds the containers and forces a recreation
	docker compose -f ./deployments/docker-compose.yml up --build --force-recreate

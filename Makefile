
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

# Builds all programs and places their binaries in the bin directory
build-all:
	mkdir -p bin
	go build -ldflags="$(LD_FLAGS)" -o ./bin/  ./cmd/...

# Removes the bin directory
clean-all:
	rm -rf bin

# Builds the docker container for the SMTP server
build-smtp-container:
	docker build \
		-t chameleon-smtp-server \
		-f build/package/chameleon-smtp-server/Dockerfile \
		$(DOCKER_BUILD_ARGS) \
		.

# Builds the docker container for the REST API server
build-api-container:
	docker build \
		-t chameleon-api-server \
		-f build/package/chameleon-api-server/Dockerfile \
		$(DOCKER_BUILD_ARGS) \
		.

# Runs docker compose
compose:
	docker compose -f ./deployments/docker-compose.yml up

# Rebuilds the containers and forces a recreation
compose-fresh:
	docker compose -f ./deployments/docker-compose.yml up --build --force-recreate

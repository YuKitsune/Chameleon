
# Builds all programs and places their binaries in the bin directory
build-all:
	mkdir -p bin
	go build -o ./bin/ ./cmd/...

# Removes the bin directory
clean-all:
	rm -rf bin

# Builds the docker container for the SMTP server
build-smtp-container:
	docker build -t chameleon-smtp-server -f build/package/chameleon-smtp-server/Dockerfile .

# Builds the docker container for the REST API server
build-api-container:
	docker build -t chameleon-api-server -f build/package/chameleon-api-server/Dockerfile .

# Runs docker compose
compose:
	docker compose -f ./deployments/docker-compose.yml up

# Rebuilds the containers and forces a recreation
compose-fresh:
	docker compose -f ./deployments/docker-compose.yml up --build --force-recreate

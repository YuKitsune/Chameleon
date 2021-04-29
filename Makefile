build_smtp_server_container:
	docker build -t chameleon-smtp-server -f build/package/chameleon-smtp-server/Dockerfile .

build_api_container:
	docker build -t chameleon-api-server -f build/package/chameleon-api-server/Dockerfile .

compose:
	docker compose -f ./deployments/docker-compose.yml up

compose-fresh:
	docker compose -f ./deployments/docker-compose.yml up --build --force-recreate

##############
## BINARIES ##
##############
FRONT_END_BINARY=client
BROKER_BINARY=broker
AUTH_BINARY=auth

############
## DOCKER ##
############
BUILD=docker build -t

##############
## COMMANDS ##
##############
# build: stops docker-compose (if running), builds the images and starts docker compose
build:
	@echo "Stopping Docker images..."
	docker-compose down
	@echo "Building service binarys..."
	$(MAKE) build_broker
	$(MAKE) build_auth
	@echo "Binarys built!"

# up: up all containers in the background
up:
	@echo "Stopping Docker images..."
	docker-compose down
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

# down: stops all containers
down:
	@echo "Stopping Docker images..."
	docker-compose down
	@echo "Docker images stopped!"

# build_broker: builds the broker image
build_broker:
	@echo "Building broker binary..."
	cd ./broker && env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "Broker binary built!"

# build_auth: builds the auth image
build_auth:
	@echo "Building auth binary..."
	cd ./auth && env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${AUTH_BINARY} ./cmd/api
	@echo "Auth binary built!"

# build_client: builds the client image
build_client:
	@echo "Building client image..."
	cd ./client && env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${FRONT_END_BINARY} ./cmd/api
	@echo "Client image built!"

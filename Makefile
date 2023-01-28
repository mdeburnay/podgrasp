FRONT_END_BINARY=client
BROKER_BINARY=broker

# up: starts all containers in the background without forcing the build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

# up_build: stops docker-compose (if running), builds the images and starts docker compose
up_build: build_broker
	@echo "Stopping Docker images..."
	docker-compose down
	@echo "Building Docker images..."
	docker-compose up --build -d
	@echo "Docker images started!"

# down: stops all containers
down:
	@echo "Stopping Docker images..."
	docker-compose down
	@echo "Docker images stopped!"

## build_broker: builds the broker image
build_broker:
	@echo "Building broker binary..."
	cd ./broker && env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "Broker binary built!"

## build_client: builds the client image
build_client:
	@echo "Building client image..."
	docker build -t client -f Dockerfile .
	@echo "Client image built!"

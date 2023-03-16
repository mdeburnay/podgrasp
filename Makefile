FRONTEND_BINARY=FRONTEND_BINARY
BROKER_BINARY=BROKER_BINARY

## UP: Starts all containers in the background without forcing a build.
up: 
	@echo "Starting Docker image..."
	@docker-compose up -d
	@echo "Docker images started."

## UP_BUILD: Builds all containers and starts them in the background.
up_build:
	@echo "Stopping Docker images (if running...)"
	@docker-compose down
	@echo "Building Docker images."
	@docker-compose up --build -d
	@echo "Docker images built and started."

## DOWN: Stops all containers.
down: 
	@echo "Stopping Docker images..."
	@docker-compose down
	@echo "Docker images stopped."

## BUILD: builds the binaries as linux executables
build:
	@echo "Building broker binary..."
	cd ./broker && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	cd ./frontend && env GOOS=linux CGO_ENABLED=0 go build -o ${FRONTEND_BINARY} ./cmd/web
	@echo "Done!"

## BUILD_FRONT
build_front:
	@echo "Building front end binary..."
	cd ./frontend && env CGO_ENABLED=0 go build -o ${FRONTEND_BINARY} ./cmd/web
	@echo "Done!"

## START: starts the front end
start: build_front
	@echo "Starting front end"
	cd ./frontend && ./${FRONTEND_BINARY}

## STOP: stop the front end
stop:
	@echo "Stopping front end..."
	@-pkill -SIGTERM -f "./${FRONTEND_BINARY}"
	@echo "Stopped front end!"

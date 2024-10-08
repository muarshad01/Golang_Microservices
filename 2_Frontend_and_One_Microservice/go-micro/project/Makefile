FRONT_END_BINARY=frontApp
BROKER_BINARY=brokerApp

################################

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## down: stop docker-compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

########################################

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_broker
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"
	
## build_broker: builds the broker binary as a linux executable
build_broker:
	@echo "Building broker binary..."
	cd ../broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "Done!"

## build_front: builds the frone end binary
build_front:
	@echo "Building front end binary..."
	cd ../frontend && env CGO_ENABLED=0 go build -o ${FRONT_END_BINARY} ./cmd/web
	@echo "Done!"

##############################

## start: starts the frontend
start: build_front
	@echo "Starting front end"
	cd ../frontend && go build -o ${FRONT_END_BINARY} ./cmd/web
	cd ../frontend && ./${FRONT_END_BINARY} &

## stop: stop the frontend
stop:
	@echo "Stopping front end..."
	@-pkill -SIGTERM -f "./${FRONT_END_BINARY}"
	@echo "Stopped front end!"
CORE_BINARY=botcore

install_local:
	@echo "Installing dependecies..."
	go mod download
	@echo "Done!"

build_local: install
	@echo "Building binary..."
	go build  -o ${CORE_BINARY} main.go
	@echo "Done!"

run_local:
	@echo "Running bot..."
	./${CORE_BINARY}
	@echo "Done!"

up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

up_build:
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"
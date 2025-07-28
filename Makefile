.PHONY: build run test clean docker-build docker-run help

# Variables
APP_NAME := fizzbuzz-api
DOCKER_IMAGE := $(APP_NAME):latest
PORT := 8080

# Default target
help: ## Show this help message
	@echo "Available commands:"
	@awk 'BEGIN {FS = ":.*##"} /^[a-zA-Z_-]+:.*##/ {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the Go application
	@echo "Building $(APP_NAME)..."
	go build -o bin/$(APP_NAME) ./cmd/$(APP_NAME)

run: ## Run the application locally
	@echo "Starting $(APP_NAME) on port $(PORT)..."
	go run main.go

test: ## Run tests
	@echo "Running tests..."
	go test -v ./...

test-coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	go test -v -cover ./...

clean: ## Clean build artifacts
	@echo "Cleaning..."
	rm -rf bin/
	go clean

docker-build: ## Build Docker image
	@echo "Building Docker image $(DOCKER_IMAGE)..."
	docker build -t $(DOCKER_IMAGE) .

docker-run: docker-build ## Build and run Docker container
	@echo "Running Docker container on port $(PORT)..."
	docker run -p $(PORT):$(PORT) $(DOCKER_IMAGE)

docker-compose-up: ## Start with docker-compose
	@echo "Starting with docker-compose..."
	docker-compose up --build

docker-compose-down: ## Stop docker-compose services
	@echo "Stopping docker-compose services..."
	docker-compose down

deps: ## Download dependencies
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy

lint: ## Run golangci-lint (requires golangci-lint to be installed)
	@echo "Running linter..."
	golangci-lint run

format: ## Format Go code
	@echo "Formatting code..."
	go fmt ./...

dev: deps ## Setup development environment
	@echo "Setting up development environment..."
	go mod tidy

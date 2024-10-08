BINARY_NAME=citymanager
DOCKER_IMAGE_NAME=citymanager-image
DOCKER_CONTAINER_NAME=citymanager-container

GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin

.PHONY: all build clean run test docker-build docker-run docker-stop

all: build

build:
	echo "Building..."
	go build -o $(GOBIN)/$(BINARY_NAME) ./cmd/server

clean:
	echo "Cleaning..."
	go clean
	rm -rf $(GOBIN)

run: build
	echo "Running..."
	$(GOBIN)/$(BINARY_NAME)

test:
	echo "Testing..."
	go test ./... -v

docker-build:
	echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE_NAME) .

docker-run: docker-build
	echo "Running Docker container..."
	docker run -d -p 50051:50051 --name $(DOCKER_CONTAINER_NAME) $(DOCKER_IMAGE_NAME)

docker-stop:
	echo "Stopping Docker container..."
	docker stop $(DOCKER_CONTAINER_NAME)
	docker rm $(DOCKER_CONTAINER_NAME)


proto:
	echo "Generating protobuf files..."
	protoc - protoc -I api/proto api/proto/*.proto --go_out=./api/gen --go_opt=paths=source_relative --go-grpc_out=./api/gen


install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.41.1

lint: install-lint-deps ### check by golangci linter
	echo "Starting linters"
	golangci-lint run

#install-lint-deps:
#	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.50.1
#
#lint: install-lint-deps ### check by golangci linter
#	echo "Running linter..."
#	golangci-lint run

deps:
	echo "Updating dependencies..."
	go mod tidy

# Запуск в Docker Compose
up:
	echo "Starting services with Docker Compose..."
	docker-compose up -d

down:
	echo "Stopping services with Docker Compose..."
	docker-compose down

help:
	echo "Available commands:"
	echo "  make build          - Build the project"
	echo "  make clean          - Clean build files"
	echo "  make run            - Run the project"
	echo "  make test           - Run tests"
	echo "  make docker-build   - Build Docker image"
	echo "  make docker-run     - Run Docker container"
	echo "  make docker-stop    - Stop and remove Docker container"
	echo "  make proto          - Generate protobuf files"
	echo "  make lint           - Run linter"
	echo "  make deps           - Update dependencies"
	echo "  make up   			- Start services with Docker Compose"
	echo "  make down 			- Stop services with Docker Compose"
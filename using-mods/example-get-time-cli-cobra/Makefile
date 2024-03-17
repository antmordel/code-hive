# Name of the Docker image
IMAGE_NAME=get-time-cli

# Docker build command
build:
	@echo "Building Docker image..."
	docker build -t $(IMAGE_NAME) .

# Docker run command with 'time' argument
run:
	@echo "Running Docker container..."
	docker run -it --rm $(IMAGE_NAME) time

.PHONY: build run
# Define variables for repeated values
DOCKER_TAG := kerwenwwer/gossip-service:latest

# Phony targets for workflows
.PHONY: all bpf-objects build docker-build docker-push

# Default target to compile the application and build the Docker image
all: build docker-build

# Rule to generate BPF object files from C source
bpf-objects: 
	go generate ./bpf/

# Rule to build the main application
build: bpf-objects
	go build -o ./bin/xdp-gossip ./main.go

# Rule to build the Docker image
docker-build:
	docker build -t $(DOCKER_TAG) .

# Rule to push the Docker image to the repository
docker-push:
	docker push $(DOCKER_TAG)
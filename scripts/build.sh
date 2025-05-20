#!/bin/sh
# Build the Go microservice binary and Docker image
set -e

# Move to project root
cd "$(dirname "$0")/.."

echo "Building Docker image..."
docker build -t go-microservice-template:latest .

echo "Docker image build complete: go-microservice-template:latest"

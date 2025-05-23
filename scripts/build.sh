#!/bin/sh
set -e

# Check if version argument is provided
if [ -z "$1" ]; then
  echo "Usage: $0 <version>"
  exit 1
fi

VERSION=$1
IMAGE_NAME="<<microservice_name>>"

echo "Building Docker image: $IMAGE_NAME:$VERSION and $IMAGE_NAME:latest"
docker build -t $IMAGE_NAME:$VERSION -t $IMAGE_NAME:latest .

echo "Docker image built: $IMAGE_NAME:$VERSION and $IMAGE_NAME:latest"
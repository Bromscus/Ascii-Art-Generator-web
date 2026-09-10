#!/bin/bash

IMAGE_NAME="ascii-art-web"
CONTAINER_NAME="ascii-art-container"

echo "Building image..."
docker image build -t $IMAGE_NAME .

echo "Removing old container if it exists..."
docker container rm -f $CONTAINER_NAME 2>/dev/null

echo "Running container..."
docker container run -d -p 8080:8080 --name $CONTAINER_NAME $IMAGE_NAME

echo "Done: http://localhost:8080"
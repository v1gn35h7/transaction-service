#!/bin/bash

# Configuration
IMAGE_NAME=" ghcr.io/v1gn35h7/transaction-service/transaction-service"
TAG="latest"
CONTAINER_NAME="pismo_ts_container"
HOST_PORT=80
CONTAINER_PORT=80

echo "Step 1: Pulling image from GitHub Packages..."
docker pull ${IMAGE_NAME}:${TAG}

echo "Step 2: Stopping and removing old container if it exists..."
docker stop ${CONTAINER_NAME} 2>/dev/null || true
docker rm ${CONTAINER_NAME} 2>/dev/null || true

echo "Step 3: Running new container..."
docker run -d \
  --name ${CONTAINER_NAME} \
  -p ${HOST_PORT}:${CONTAINER_PORT} \
  ${IMAGE_NAME}:${TAG}

echo "Success! Container is running on port ${HOST_PORT}."
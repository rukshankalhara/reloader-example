#!/bin/bash

IMAGE_NAME="mathisve/reloader-example-1"
ARCHITECTURE="linux/amd64"  # x86 architecture
TAG="latest"  # You can modify this tag if needed

docker buildx build --platform $ARCHITECTURE -t $IMAGE_NAME:$TAG --push .


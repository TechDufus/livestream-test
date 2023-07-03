#!/usr/bin/env bash

# Exit on error
set -e

# Get the current directory
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

# Navigate to the root directory
pushd $DIR/../ > /dev/null

# Load the environment variables
source .env

# Login to Docker Hub
. $DIR/login.sh

# Push the image to Docker Hub
echo "Pushing the image to Docker Hub... $DOCKERHUB_USER/$DOCKER_IMAGE_NAME:$DOCKER_IMAGE_PROD_TAG"
docker push $DOCKERHUB_USER/$DOCKER_IMAGE_NAME:$DOCKER_IMAGE_PROD_TAG

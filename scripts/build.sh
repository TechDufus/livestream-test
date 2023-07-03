#!/usr/bin/env bash

# Exit on error
set -e

# Get the current directory
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Navigate to the root directory
pushd $DIR/../ > /dev/null

# Import .env vars
source .env

# Build the docker image
docker build . -t $DOCKERHUB_USER/$DOCKER_IMAGE_NAME:$DOCKER_IMAGE_PROD_TAG

# Navigate back to the original directory
popd > /dev/null

#!/usr/bin/env bash


# Exit on error
set -e

# Get the current directory
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

# Navigate to the root directory
pushd $DIR/../ > /dev/null

# Load the environment variables
source .env

is_logged_in() {
  cat ~/.docker/config.json | jq -r --arg url "${DOCKERHUB_REPO_URL}" '.auths | has($url)'
}

if is_logged_in > /dev/null; then
  echo "Already logged in to Docker Hub"
else
    # Login to Docker Hub
    echo "Logging in to Docker Hub... User: $DOCKERHUB_USER"
    docker login --username $DOCKERHUB_USER
fi

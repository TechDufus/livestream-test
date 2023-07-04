#!/usr/bin/env bash

# Get the current directory (so this script can be run from anywhere)
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

. $DIR/build.sh
. $DIR/push.sh

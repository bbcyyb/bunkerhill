#!/bin/bash

set -e

export CGO_ENABLED=0
export GOOS=linux

echo "** Running makefile to dev package"
make dev_docker \
    PARAMS_HOST=0.0.0.0 \
    PARAMS_PORT=3030

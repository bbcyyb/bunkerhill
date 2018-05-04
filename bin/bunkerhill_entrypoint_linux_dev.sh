#!/bin/bash

set -e

export CGO_ENABLED=0
export GOOS=linux

echo "** Running makefile to dev package"
make dev

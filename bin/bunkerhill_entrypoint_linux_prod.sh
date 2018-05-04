#!/bin/bash

set -e
echo $(pwd)
./bunkerhill-server --scheme=http --port=3030 --host=0.0.0.0

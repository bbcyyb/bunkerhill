#!/bin/bash

set -e
export PATH=${GOPATH}:${GOPATH}/bin:$PATH

# turns of c-code, which we aren't using and isn't allowed during cross-compile
export CGO_ENABLED=0 
export GOOS=linux


echo "** Fetching glide for docker environment"
go get github.com/Masterminds/glide
echo "** Fetching goimports for format code"
go get golang.org/x/tools/cmd/goimports
echo "** Fetching go-swagger from source code"
go get github.com/go-swagger/go-swagger/cmd/swagger

echo "** Running makefile to build package"
make all 

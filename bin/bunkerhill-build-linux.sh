#!/bin/bash

set -e
export PKG=github.com/bbcyyb/bunkerhill
export GOPATH=$(pwd)
export PATH=${GOPATH}:${GOPATH}/bin:$PATH

# turns of c-code, which we aren't using and isn't allowed during cross-compile
export CGO_ENABLED=0 
export GOOS=linux


mkdir -p src/${PKG}
mkdir -p bin

echo "** Fetching glide for docker environment"
go get github.com/Masterminds/glide
pushd src/github.com/Masterminds/glide
make build
cp ./glide ${GOPATH}/bin
popd

echo "** Fetching goimports for format code"
go get golang.org/x/tools/cmd/goimports
pushd src/golang.org/x/tools/cmd/goimports
go install
popd

echo "** Install go-swagger from source code"
# We use go-swagger to re generate swagger code but it's not dependency for bunkerhill source code
go get -u github.com/go-swagger/go-swagger/cmd/swagger

build_sub_copy=( "glide.yaml" "Makefile" "Makefile.variables" "cmd" "handlers" "models" "restapi" "storage" "swagger" "vendor")
echo "** Copying for each file and file which need to be built later"
for sub in "${build_sub_copy[@]}"; do
    cp -r $(pwd)/${sub} src/${PKG}/${sub}
done

pushd src/${PKG}
echo "** Running makefile to build package"
${CGO_ENABLED} ${GOOS} make all 
popd

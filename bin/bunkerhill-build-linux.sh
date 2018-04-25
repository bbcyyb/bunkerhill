#!/bin/bash

set -e
export PKG=github.com/bbcyyb/bunkerhill
export GOPATH=$(pwd)
export PATH=${GOPATH}:${GOPATH}/bin:$PATH


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

echo "** Install go-swagger as static binary"
# We use go-swagger to re generate swagger code but it's not dependency for bunkerhill source code
latestv=$(curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/latest | jq -r .tag_name)
curl -o /usr/local/bin/swagger -L'#' https://github.com/go-swagger/go-swagger/releases/download/$latestv/swagger_$(echo `uname`|tr '[:upper:]' '[:lower:]')_amd64
chmod +x /usr/local/bin/swagger

# build_sub_linked=( "cmd" "handlers" "models" "restapi" "storage" "swagger" "version" )
# echo "** Creating soft link for each folder and file which need to be built later"
# for sub in "${build_sub_linked[@]}"; do
#     if [ -L src/${PKG}/${sub} ]; then
#         echo "src/${PKG}/${sub} already linked to $(pwd)/${sub}, skipping"
#     else
#         ln -s $(pwd)/${sub} src/${PKG}/${sub}
#     fi
# done

build_sub_copy=( "glide.yaml" "Makefile" "Makefile.variables" "cmd" "handlers" "models" "restapi" "storage" "swagger")
echo "** Copying for each file and file which need to be built later"
for sub in "${build_sub_copy[@]}"; do
    cp -r $(pwd)/${sub} src/${PKG}/${sub}
done

pushd src/${PKG}
echo "** Running makefile to build package"
make all
ls -al restapi/*
echo "GOPATH is $GOPATH"
echo "pwd is $(pwd)"
make install
popd

# !/usr/bin/env bash

if [ ! -f install.sh ]; then
    echo 'install must be run within its containerfolder' 1>&2
    exit 1
fi

CURDIR=`pwd`
OLDGOPATH=${GOPATH}
MAIN_PATH=bunkerhill/cmd/bunkerhill-server
export GOPATH=$CURDIR

gofmt -w src

go install ${MAIN_PATH}

export GOPATH=${OLDGOPATH}

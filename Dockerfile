FROM golang:1.9.1 as builder

WORKDIR /golang/

COPY bin bin
COPY cmd cmd
COPY handlers handlers
COPY models models
COPY restapi restapi
COPY storage storage
COPY swagger swagger
COPY version version
COPY Makefile Makefile
COPY Makefile.variables Makefile.variables


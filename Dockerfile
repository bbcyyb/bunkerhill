FROM golang:1.10.1 as builder

WORKDIR /golang/

COPY bin bin
COPY cmd cmd
COPY handlers handlers
COPY models models
COPY restapi restapi
COPY storage storage
COPY swagger swagger
COPY version version
COPY vendor vendor
COPY glide.yaml glide.yaml
COPY Makefile Makefile
COPY Makefile.variables Makefile.variables

RUN chmod +x bin/*.sh

RUN apt-get update && apt-get install -y jq

RUN bin/bunkerhill-build-linux.sh

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /golang/bin/bunkerhill-server ./bukerhill-server

COPY version version

EXPOSE 3030

RUN ls -al

ENTRYPOINT ["./bunkerhill-server", "--scheme=http", "--port=3030", "--host=0.0.0.0"]

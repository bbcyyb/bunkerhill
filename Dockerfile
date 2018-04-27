FROM golang:1.10.1 as builder

WORKDIR /golang/

COPY bin script
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

RUN chmod +x script/*.sh

RUN apt-get update && apt-get install -y jq

RUN script/bunkerhill-build-linux.sh


# FROM alpine:latest
FROM golang:1.10.1-alpine3.7

RUN apk --no-cache add --update curl bash ca-certificates \
 && rm -rf /var/cache/apk/*

WORKDIR /root/

COPY --from=builder /golang/bin/bunkerhill-server ./bukerhill-server

RUN ls -al

EXPOSE 3030

ENTRYPOINT ["./bunkerhill-server", "--scheme=http", "--port=3030", "--host=0.0.0.0"]

FROM golang:1.10.1 as builder

WORKDIR /go/src/github.com/bbcyyb/bunkerhill

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

# use a minimal alpine image
FROM alpine:latest

RUN apk --no-cache add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /root/

COPY --from=builder /go/src/github.com/bbcyyb/bunkerhill/bin/bunkerhill-server .

COPY version version

EXPOSE 3030

ENTRYPOINT ["./bunkerhill-server", "--scheme=http", "--port=3030", "--host=0.0.0.0"]

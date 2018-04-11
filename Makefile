# Basic go commands
GOCMD=go
GOINSTALL=$(GOCMD) install
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Basic path
BIN=$(GOPATH)/bin
PKG=$(GOPATH)/pkg
SRC=$(GOPATH)/src
BINARY_NAME=bunkerhill
BINARY_UNIX=$(BINARY_NAME)_unix

all: test build
install:
	$(GOINSTALL) github.com/bbcyyb/bunkerhill/cmd/bunkerhill-server
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
docker-build:
	docker run --rm -it -v "$(GOPATH)":/go -w /go/src/bitbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY_UNIX)" -v

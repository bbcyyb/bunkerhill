SHELL := /bin/bash

include Makefile.variables

.PHONY: all
all: clean install

.PHONY: install
install:
	$(GOINSTALL) $(SERVER)

.PHONY: build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v $(SERVER)

.PHONY: test
test:
	$(GOTEST) -v ./...

.PHONY: clean
clean:
	$(GOCLEAN)
	rm -f $(BIN)/$(BINARY_NAME)

.PHONY: run
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

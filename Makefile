SHELL := /bin/bash

include Makefile.variables

.PHONY: all
all: clean install

.PHONY: install
install:
	$(GOINSTALL) $(SERVER_PATH)

.PHONY: build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v $(SERVER_PATH)

.PHONY: regen
regen:
	$(SWAGGERGEN) server --target $(SRC)/$(PROJECT_PATH) --name $(PROJECT_NAME) --spec $(SRC)/$(PROJECT_PATH)/swagger/swagger.yaml --exclude-main

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

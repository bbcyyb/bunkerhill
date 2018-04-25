SHELL := /bin/bash

include Makefile.variables

.PHONY: all
all: clean vendor_install regen install


.PHONY: dev
dev: clean install
	$(BIN)/$(BINARY_NAME) --host 10.62.59.210 --port 3000

.PHONY: install
install: fmt
	$(GOINSTALL) $(SERVER_PATH)

.PHONY: build
build: fmt
	$(GOBUILD) -o $(BINARY_NAME) -v $(SERVER_PATH)

.PHONY: test
test: fmt
	$(GOTEST) -v ./...

.PHONY: regen
regen:
	rm -rf $(SRC)/$(PROJECT_PATH)/models
	rm -rf $(SRC)/$(PROJECT_PATH)/restapi/operations
	$(SWAGGER_GEN) server --target $(SRC)/$(PROJECT_PATH) --name $(PROJECT_NAME) --spec $(SRC)/$(PROJECT_PATH)/swagger/swagger.yaml --exclude-main

.PHONY: validate
validate:
	$(SWAGGER_VALIDATE) swagger/swagger.yaml

.PHONY: vendor_init
vendor_init:
	rm -f $(SRC)/$(PROJECT_PATH)/glide.yaml
	glide init

.PHONY: vendor_update
vendor_update:
	rm -f $(SRC)/$(PROJECT_PATH)/glide.lock
	glide up

.PHONY: vendor_install
vendor_install:
	glide install

.PHONY: vendor
vendor: vendor_init vendor_update vendor_install

.PHONY: clean
clean:
	$(GOCLEAN)
	rm -f $(BIN)/$(BINARY_NAME)

.PHONY: fmt
fmt:
	$(GOFMT) -w $$(find . -type f -name "*.go" -not -path "./vendor/*")

SHELL := /bin/bash

include Makefile.variables

.PHONY: all
all: clean vendor_install regen build


.PHONY: dev
dev: clean install
	@echo "Makefile-------> $(BIN)/$(BINARY_NAME) --host 10.62.59.210 --port 3000"
	$(BIN)/$(BINARY_NAME) --host 10.62.59.210 --port 3000

.PHONY: install
install: fmt
	@echo "Makefile-------> CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) $(GOINSTALL) $(SERVER_PATH)"
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) $(GOINSTALL) $(SERVER_PATH)

.PHONY: build
build: fmt
	@echo "Makefile-------> CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) $(GOBUILD) -a -installsuffix cgo -o $(BINARY_NAME) $(SERVER_PATH)"
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) $(GOBUILD) -a -installsuffix cgo -o $(BINARY_NAME) ${SERVER_PATH}

.PHONY: test
test: fmt
	@echo "Makefile-------> $(GOTEST) -v ./..."
	$(GOTEST) -v ./...

.PHONY: regen
regen:
	rm -rf $(SRC)/$(PROJECT_PATH)/models
	rm -rf $(SRC)/$(PROJECT_PATH)/restapi/operations
	@echo "Makefile-------> $(SWAGGER_GEN) server --target $(SRC)/$(PROJECT_PATH) --name $(PROJECT_NAME) --spec $(SRC)/$(PROJECT_PATH)/swagger/swagger.yaml --exclude-main"
	$(SWAGGER_GEN) server --target $(SRC)/$(PROJECT_PATH) --name $(PROJECT_NAME) --spec $(SRC)/$(PROJECT_PATH)/swagger/swagger.yaml --exclude-main

.PHONY: validate
validate:
	@echo "Makefile-------> $(SWAGGER_VALIDATE) swagger/swagger.yaml"
	$(SWAGGER_VALIDATE) swagger/swagger.yaml

.PHONY: vendor_init
vendor_init:
	rm -f $(SRC)/$(PROJECT_PATH)/glide.yaml
	@echo "Makefile-------> glide init"
	glide init

.PHONY: vendor_update
vendor_update:
	rm -f $(SRC)/$(PROJECT_PATH)/glide.lock
	@echo "Makefile-------> glide up"
	glide up

.PHONY: vendor_install
vendor_install:
	@echo "Makefile-------> glide install"
	glide install

.PHONY: vendor
vendor: vendor_init vendor_update vendor_install

.PHONY: clean
clean:
	@echo "Makefile-------> $(GOCLEAN)"
	$(GOCLEAN)
	rm -f $(BIN)/$(BINARY_NAME)

.PHONY: fmt
fmt:
	@echo "Makefile-------> $(GOFMT) -w $$(find . -type f -name "*.go" -not -path "./vendor/*")"
	$(GOFMT) -w $$(find . -type f -name "*.go" -not -path "./vendor/*")

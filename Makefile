SHELL := /bin/bash

include Makefile.variables

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
	$(SWAGGERGEN) server --target $(SRC)/$(PROJECT_PATH) --name $(PROJECT_NAME) --spec $(SRC)/$(PROJECT_PATH)/swagger/swagger.yaml --exclude-main

.PHONY: vendor
vendor:
	rm -f $(SRC)/$(PROJECT_PATH)/glide.yaml
	rm -f $(SRC)/$(PROJECT_PATH)/glide.lock
	glide init
	glide install

.PHONY: clean
clean:
	$(GOCLEAN)
	rm -f $(BIN)/$(BINARY_NAME)

.PHONY: fmt
fmt:
	$(GOFMT) -w $$(find . -type f -name "*.go" -not -path "./vendor/*")

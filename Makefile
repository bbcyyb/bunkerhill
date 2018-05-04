SHELL := /bin/bash

include Makefile.variables

.PHONY: help
help:
	@echo "Usage:	make COMMAND"
	@echo ""
	@echo "Commands:"
	@echo "   all                 Run clean, vendor_install, regen, build in sequence"
	@echo "   build               Build compiles tha executable packages"
	@echo "   clean               Clean removes object files from package source directories"
	@echo " * compose_up          Use docker-compose to create and start services"
	@echo " * dev                 Dev mode, run compiles and runs the main package comprising the named Go source files"
	@echo " * docker_build        Build an image from a Dockerfile"
	@echo " * docker_build_dev    Build an image as development environment"
	@echo "   fmt                 Format Go code and update Go import lines, adding missling ones and removing unreferenced ones."
	@echo "   help                Get help on a command"
	@echo "   install             Install compiles and installs the packages"
	@echo " * k8s                 Use kubernetes to create and start services"
	@echo "   regen               Regenerate go-swagger code (main.go and configure_bunkerhill.go don't be rewritten)"
	@echo "   test                Automate testing the packages"
	@echo "   validate            Validate swagger file"
	@echo "   vendor_init         Use glide to initialize vendor package info, creating a glide.yaml file"
	@echo "   vendor_update       Use glide to update project's dependencies"
	@echo "   vendor_install      Use glide to install project's dependencies"
	@echo "   vendor              Run vendor_init, vendor_update, vendor_install in sequence"
	@echo " "

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
vendor-init:
	rm -f $(SRC)/$(PROJECT_PATH)/glide.yaml
	@echo "Makefile-------> glide init"
	glide init

.PHONY: vendor_update
vendor-update:
	rm -f $(SRC)/$(PROJECT_PATH)/glide.lock
	@echo "Makefile-------> glide up"
	glide up

.PHONY: vendor_install
vendor-install:
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

.PHONY: docker_build
docker_build:
	@echo "Makefile-------> $(DOCKER_BUILD) -t $(PROD_IMAGE_NAME) -f $(DOCKERFILE_PROD) ." 
	$(DOCKER_BUILD) -t $(PROD_IMAGE_NAME) -f $(DOCKERFILE_PROD) .
	
.PHONY: docker_build_dev
docker_build_dev:
	@echo "Makefile-------> $(DOCKER_BUILD) -t $(DEV_IMAGE_NAME) -f $(DOCKERFILE_DEV) ."	
	$(DOCKER_BUILD) -t $(DEV_IMAGE_NAME) -f $(DOCKERFILE_DEV) .	

.PHONY: compose_up
compose_up:
	$(COMPOSE_CMD) -f $(COMPOSEFILE) $(COMPOSE_UP)

SERVICE := noah

all: help

help: ## Show help messages
	@echo "Container - ${SERVICE} "
	@echo
	@echo "Usage:\tmake COMMAND"
	@echo
	@echo "Commands:"
	@sed -n '/##/s/\(.*\):.*##/  \1#/p' ${MAKEFILE_LIST} | grep -v "MAKEFILE_LIST" | column -t -c 2 -s '#'

APP = noah
BINARY := $(shell basename "$(PWD)")
VERSION := $(shell git describe --tags --dirty --always)
BUILD := $(shell git rev-parse HEAD)

LDFLAGS=-ldflags
LDFLAGS += "-X=github.com/airdb/adb/internal/adblib.Version=$(VERSION) \
            -X=github.com/airdb/adb/internal/adblib.Build=$(BUILD) \
            -X=github.com/airdb/adb/internal/adblib.BuildTime=$(shell date +%s)"

SYSTEM:=
#myos = $(word 1, $@)
#ifndef $myos
#	myos = "$(shell uname | tr A-Z a-z)"
#endif

.PHONY: test

#all: build

test:
	go test -v ./...

dev:
	#CGO_ENABLED=0 $(SYSTEM) GOARCH=amd64 go run $(LDFLAGS) main.go
	CGO_ENABLED=0 $(SYSTEM) go build $(LDFLAGS) main.go

build: ## Build binary
	bash ./build/util.sh until::build

lint:
	go fmt ./...
	golangci-lint run

install: ## Install adb
	cp output/adb $(shell which adb)

preinstall:
	go install -ldflags -X=github.com/airdb/adb/internal/adblib.BuildTime=$(date +%s) github.com/airdb/adb@dev

deploy:
	flyctl deploy -a ${APP}

conf secret:
	flyctl secrets import -a ${APP} < .env

bash:
	flyctl ssh console -a  ${APP} -C /bin/bash


PLATFORMS := windows linux darwin
os = $(word 1, $@)

.PHONY: $(PLATFORMS)
$(PLATFORMS):
	mkdir -p release
	CGO_ENABLED=0 GOOS=$(os) GOARCH=amd64 go build $(LDFLAGS) -o release/$(BINARY)-$(os)

#release: windows linux darwin
release: linux darwin

.PHONY: release build

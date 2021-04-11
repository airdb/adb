BINARY := $(shell basename "$(PWD)")
VERSION := $(shell git describe --dirty --always)
BUILD := $(shell git rev-parse HEAD)

LDFLAGS=-ldflags
LDFLAGS += "-X=airdb.io/airdb/adb/internal/adblib.Version=$(VERSION) \
            -X=airdb.io/airdb/adb/internal/adblib.Build=$(BUILD) \
            -X=airdb.io/airdb/adb/internal/adblib.BuildTime=$(shell date +%s)"

SYSTEM:=
#myos = $(word 1, $@)
#ifndef $myos
#	myos = "$(shell uname | tr A-Z a-z)"
#endif

.PHONY: test

all: build

test:
	go test -v ./...

build:
	CGO_ENABLED=0 $(SYSTEM) GOARCH=amd64 go build $(LDFLAGS) -o $(BINARY)

PLATFORMS := windows linux darwin
os = $(word 1, $@)

.PHONY: $(PLATFORMS)
$(PLATFORMS):
	mkdir -p release
	CGO_ENABLED=0 GOOS=$(os) GOARCH=amd64 go build $(LDFLAGS) -o release/$(BINARY)-$(os)

.PHONY: release
#release: windows linux darwin
release: linux darwin

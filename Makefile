BINARY := adb
VERSION = 0.0.0

LDFLAGS=-ldflags "-X=github.com/airdb/adb/cmd.Version=$(VERSION) -X=github.com/airdb/adb/cmd.BuildTime=$(shell date +%s)"

myos = $(word 1, $@)
ifndef $myos
	myos = "$(shell uname | tr A-Z a-z)"
endif

.PHONY: test

all: build

test:
	go test -v ./...

build:
	CGO_ENABLED=0 GOOS=$(myos) GOARCH=amd64 go build $(LDFLAGS) -o $(BINARY)

PLATFORMS := windows linux darwin
os = $(word 1, $@)

.PHONY: $(PLATFORMS)
$(PLATFORMS):
	mkdir -p release
	CGO_ENABLED=0 GOOS=$(os) GOARCH=amd64 go build $(LDFLAGS) -o release/$(BINARY)-$(os)

.PHONY: release
release: windows linux darwin

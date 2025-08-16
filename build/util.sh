#!/usr/bin/env bash

# BINARY := $(shell basename "$(PWD)")
VERSION=$(git describe --dirty --always)
BUILD=$(git rev-parse HEAD)
BUILD_TS=$(date +%s)

LDFLAGS="-s -w \
	-X=github.com/airdb/adb/internal/adblib.Version=$VERSION \
	-X=github.com/airdb/adb/internal/adblib.Build=$BUILD \
	-X=github.com/airdb/adb/internal/adblib.BuildTime=$BUILD_TS"

function until::build() {
	#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "$LDFLAGS" -o ./output/adb main.go
	CGO_ENABLED=0 GOOS=linux go build -ldflags "$LDFLAGS" -o ./output/adb main.go
}

$1

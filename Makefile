os = $(word 1, $@)

ifndef $os
	os = "$(shell uname | tr A-Z a-z)"
endif

.PHONY: test

all: build

test:
	go test -v ./...

build:
        timeNow=$(shell date +%s)
        timeflag = "github.com/airdb/adb/cmd.BuildTime=$timeNow"
        versionflag = "github.com/airdb/adb/cmd.Version=${{ steps.bump_version.outputs.tag }}"

        echo $timeflag, $versionflag
        CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o adb-linux -ldflags "-X $timeflag" -ldflags "-X $versionflag" main.go
        CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o adb-darwin -ldflags "-X $timeflag" -ldflags "-X $versionflag" main.go
        CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o adb-windows.exe -ldflags "-X $timeflag" -ldflags "-X $versionflag" main.go

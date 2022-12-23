package main

import (
	"github.com/airdb/adb/cmd"
	"github.com/airdb/adb/internal/adblib"
)

//go:generate go build -o adb main.go

func main() {
	adblib.InitDotEnv()

	cmd.Execute()
}

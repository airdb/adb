package main

import (
	"github.com/airdb/adb/cmd"
)

//go:generate go build -o adb main.go

func main() {
	// adblib.InitDotEnv()

	cmd.Execute()
}

package main

import (
	"fmt"
	"os"

	"github.com/airdb/adb/cmd"
	"github.com/airdb/adb/internal/adblib"
)

//go:generate go build -o adb main.go

func main() {
	if err := adblib.InitDotEnv(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	cmd.Execute()
}

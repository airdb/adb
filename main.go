package main

import (
	"log"

	"github.com/airdb/adb/cmd"
	"github.com/airdb/adb/internal/adblib"
)

//go:generate go build -o adb main.go

func main() {
	adblib.InitDotEnv()

	log.Println(adblib.AdbConfig)

	cmd.Execute()
}

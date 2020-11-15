package main

import (
	"airdb.io/airdb/adb/cmd"
)

//go:generate go build -o adb main.go

func main() {
	cmd.Execute()
}

package main

import (
	"log"

	"github.com/airdb/adb/cmd"
	"github.com/joho/godotenv"
)

//go:generate go build -o adb main.go

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cmd.Execute()
}

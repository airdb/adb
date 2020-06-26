package main

import (
	"github.com/airdb-template/gin-api/web"
)

//go:generate go build -o main main.go
func main() {
	web.Run()
}

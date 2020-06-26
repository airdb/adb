package main

import (
	"{{ .GoModulePath }}/web" // invalid
)

//go:generate go build -o main main.go
func main() {
	web.Run()
}

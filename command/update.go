package command

import (
	"log"
	"os/exec"
)

func update() error {
	cmd := exec.Command("go", "get", "github.com/airdb/adb")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("update successfully")
	}
	return err
}

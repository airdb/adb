package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os/exec"
)

var updateCommand = &cobra.Command{
	Use:                "update",
	Short:              "update self",
	Long:               "update self",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		update()
	},
}

func update() {
	repo := "github.com/airdb/adb"
	cmd := exec.Command("go", "get", repo)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("update successfully")
	}
}

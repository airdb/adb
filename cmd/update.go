package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/spf13/cobra"
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

	ldflag := fmt.Sprintf("-X github.com/airdb/adb/cmd.BuildTime=%d", time.Now().Unix())
	cmd := exec.Command("go", "get", "--ldflags", ldflag, repo)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("update successfully")
	}
}
package cmd

import (
	"fmt"
	"github.com/imroc/req"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"runtime"
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
	dl := "https://github.com/airdb/adb/releases/latest/download/adb"
	if runtime.GOOS == "darwin" {
		dl = dl + "-" + runtime.GOOS
		fmt.Println(dl)
	}

	resp, err := req.Get(dl)
	tmpPath := "/tmp/adb-latest"
	if err == nil {
		err = resp.ToFile(tmpPath)
	}

	if err != nil {
		log.Println("Error: download package failed! ", err)
		return
	}

	err = updateBinary(tmpPath)
	if err != nil {
		log.Println("update failed!")
	} else {
		log.Println("update successfully!")
	}
}

func updateBinary(tmpPath string) error {
	adbPath, err := exec.LookPath("adb")
	if err == nil {
		err = os.Rename(tmpPath, adbPath)
	}

	return err
}

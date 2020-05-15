package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/imroc/req"
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
	dl := "https://github.com/airdb/adb/releases/latest/download/adb"
	if runtime.GOOS == "darwin" {
		dl = dl + "-" + runtime.GOOS
	}

	fmt.Printf("It will take about 1 minute for downloading.\nDownload url: %s\n", dl)

	tmpPath := "/tmp/adb-latest"

	resp, err := req.Get(dl)
	if err == nil {
		err = resp.ToFile(tmpPath)
	}

	if err != nil {
		log.Println("Error: download package failed! ", err)
		return
	}

	err = os.Chmod(tmpPath, 0755)
	if err == nil {
		err = updateBinary(tmpPath)
	}

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

// completionCmd represents the completion command.
var completionBashCommand = &cobra.Command{
	Use:   "completion",
	Short: "Generates bash completion scripts",
	Long: `To load completion run

. <(bitbucket completion)

To configure your bash shell to load completions for each session add to your bashrc

# MacOS:
# adb completion >/usr/local/etc/bash_completion.d/adb
# ~/.bashrc or ~/.profile
. <(bitbucket completion)
`,
	Run: func(cmd *cobra.Command, args []string) {
		err := rootCmd.GenBashCompletion(os.Stdout)
		if err != nil {
			fmt.Println("Generates bash completion scripts failed!")
		}
	},
}

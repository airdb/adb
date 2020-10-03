package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/minio/selfupdate"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:                "update",
	Short:              "Self update adb",
	Long:               "Self update adb",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		doUpdate()
	},
}

func doUpdate() {
	dl := "https://github.com/airdb/adb/releases/latest/download/adb"
	if runtime.GOOS == "darwin" {
		dl = dl + "-" + runtime.GOOS
	}

	fmt.Printf("It will take about 1 minute for downloading.\nDownload url: %s\n", dl)

	client := &http.Client{}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, dl, nil)
	if err != nil {
		log.Println(err)

		return
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)

		return
	}

	defer resp.Body.Close()

	err = selfupdate.Apply(resp.Body, selfupdate.Options{})
	if err != nil {
		log.Println("update failed!")
	} else {
		log.Println("update successfully!")
	}
}

func updateCmdInit() {
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(completionBashCmd)

	completionBashCmd.PersistentFlags().BoolVarP(&writeCompletionFile, "write_file", "w", false, "write completion file")
}

var writeCompletionFile bool

var completionBashCmdLongDesc = `To load completion run

. <(bitbucket completion)

To configure your bash shell to load completions for each session add to your bashrc

# MacOS:
# adb completion >/usr/local/etc/bash_completion.d/adb
# ~/.bashrc or ~/.profile
. <(bitbucket completion)
`

// CompletionCmd represents the completion command.
var completionBashCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generates bash completion scripts",
	Long:  completionBashCmdLongDesc,
	Run: func(cmd *cobra.Command, args []string) {
		if writeCompletionFile {
			completionFile := "/usr/local/etc/bash_completion.d/adb"

			err := rootCmd.GenFishCompletionFile(completionFile, true)
			if err != nil {
				panic(err)
			}

			fmt.Println("Generates bash completion scripts successfully, file:", completionFile)

			return
		}

		err := rootCmd.GenBashCompletion(os.Stdout)
		if err != nil {
			fmt.Println("Generates bash completion scripts failed!")
		}
	},
}

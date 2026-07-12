package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/minio/selfupdate"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:                "update",
	Short:              "Self update adb",
	Long:               "Self update adb",
	DisableFlagParsing: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return doUpdate()
	},
}

const updateTimeout = 5 * time.Minute

func downloadURL() string {
	dl := "https://github.com/airdb/adb/releases/latest/download/adb"

	// Keep the legacy artifact names for amd64: "adb" (linux) and
	// "adb-darwin". Other combinations use "adb-<goos>-<goarch>".
	switch {
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		dl += "-darwin"
	default:
		dl += "-" + runtime.GOOS + "-" + runtime.GOARCH
	}

	return dl
}

func doUpdate() error {
	dl := downloadURL()

	fmt.Printf("It will take about 1 minute for downloading.\nDownload url: %s\n", dl)

	client := &http.Client{Timeout: updateTimeout}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, dl, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download %s: unexpected status %s", dl, resp.Status)
	}

	if err := selfupdate.Apply(resp.Body, selfupdate.Options{}); err != nil {
		return fmt.Errorf("update failed: %w", err)
	}

	fmt.Println("update successfully!")

	return nil
}

func updateCmdInit() {
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(completionBashCmd)

	completionBashCmd.PersistentFlags().BoolVarP(&writeCompletionFile, "write_file", "w", false, "write completion file")
}

var writeCompletionFile bool

var completionBashCmdLongDesc = `To load completion run

. <(adb completion)

To configure your bash shell to load completions for each session add to your bashrc

# MacOS:
# adb completion >/usr/local/etc/bash_completion.d/adb
# ~/.bashrc or ~/.profile
. <(adb completion)
`

// CompletionCmd represents the completion command.
var completionBashCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generates bash completion scripts",
	Long:  completionBashCmdLongDesc,
	RunE: func(cmd *cobra.Command, args []string) error {
		if writeCompletionFile {
			completionFile := "/usr/local/etc/bash_completion.d/adb"

			if err := rootCmd.GenBashCompletionFile(completionFile); err != nil {
				return err
			}

			fmt.Println("Generates bash completion scripts successfully, file:", completionFile)

			return nil
		}

		return rootCmd.GenBashCompletion(os.Stdout)
	},
}

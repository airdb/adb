package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

type VersionInfo struct {
	Version   string
	GitCommit string
}

var version *VersionInfo

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Version information",
	Long:  "Version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%v\n", rootCmd.Long)
		if version != nil {
			fmt.Printf("Version: %v\n", version.Version)
			fmt.Printf("Git Commit: %v\n", version.GitCommit)
		}
		fmt.Printf("Go Version: %v\n", runtime.Version())
	},
}

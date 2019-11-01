package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

// Build version info.
var BuildTime string
var BuildVersion string

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Version information",
	Long:  "Version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%v\n", rootCmd.Long)
		fmt.Printf("GoVersion: %v\n", runtime.Version())
		fmt.Printf("BuildTime: %v\n", BuildTime)
		fmt.Printf("BuildVersion: %v\n", BuildVersion)
	},
}

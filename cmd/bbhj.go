package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var bbhjCommand = &cobra.Command{
	Use:                "bbhj",
	Short:              "bbhj operation",
	Long:               "bbhj operation",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bbhj operation")
	},
}

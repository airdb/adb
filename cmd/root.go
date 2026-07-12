package cmd

import (
	"fmt"
	"os"

	"github.com/airdb/adb/internal/adblib"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "adb",
	Short:        "Airdb Development Builder",
	Long:         "Airdb Development Builder Command Line Interface",
	SilenceUsage: true,
}

func Execute() {
	if err := adblib.Init(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	rootCmd.Version = adblib.GetVersion()

	rootCmd.AddCommand(weatherCommand)
	rootCmd.AddCommand(wikiCommand)
	wikiCommand.AddCommand(interviewWikiCommand)

	mysqlCmdInit()
	serviceCmdInit()
	dnsTxtCmdInit()
	hostCmdInit()

	initSlack()
	initManCommand()
	initLogin()
	updateCmdInit()

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

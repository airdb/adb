package cmd

import (
	"fmt"
	"os"

	"github.com/airdb/adb/internal/adblib"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "adb",
	Short: "Airdb Development Builder",
	Long:  "Airdb Development Builder Command Line Interface",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	adblib.Init()
	fmt.Println(adblib.ConfigNew)

	rootCmd.Version = adblib.GetVersion()

	// rootCmd.PersistentFlags().StringVarP(&GlobalFlags.Type, "type", "t", "com", "Top level domain")
	rootCmd.AddCommand(envCommand)
	rootCmd.AddCommand(releaseCommand)

	rootCmd.AddCommand(weatherCommand)
	rootCmd.AddCommand(wikiCommand)
	wikiCommand.AddCommand(interviewWikiCommand)

	mysqlCmdInit()
	serviceCmdInit()
	hostCmdInit()

	// initConfigCmd()
	initCert()
	initSlack()
	initManCommand()

	initLogin()

	updateCmdInit()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

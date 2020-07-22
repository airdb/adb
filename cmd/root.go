package cmd

import (
	"fmt"
	"os"

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
	// rootCmd.PersistentFlags().StringVarP(&GlobalFlags.Type, "type", "t", "com", "Top level domain")
	rootCmd.AddCommand(versionCommand)
	rootCmd.AddCommand(envCommand)
	rootCmd.AddCommand(bbhjCommand)
	rootCmd.AddCommand(releaseCommand)
	rootCmd.AddCommand(manCommand)
	manCommand.AddCommand(gitInitCommand)
	manCommand.AddCommand(dockerInitCommand)
	manCommand.AddCommand(cloudInitCommand)
	manCommand.AddCommand(toolsInitCommand)
	manCommand.AddCommand(brewInitCommand)
	manCommand.AddCommand(githubInitCommand)
	manCommand.AddCommand(vimInitCommand)
	manCommand.AddCommand(osinitCommand)
	manCommand.AddCommand(kubeCommand)
	manCommand.AddCommand(helmCommand)
	manCommand.AddCommand(terraformCommand)
	manCommand.AddCommand(opensslCommand)

	rootCmd.AddCommand(loginCommand)
	rootCmd.AddCommand(weatherCommand)
	rootCmd.AddCommand(wikiCommand)
	wikiCommand.AddCommand(interviewWikiCommand)
	wikiCommand.AddCommand(listWikiCommand)

	genCmdInit()
	mysqlCmdInit()
	serviceCmdInit()
	hostCmdInit()

	updateCmdInit()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

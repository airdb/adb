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

func init() {
	rootCmd.AddCommand(versionCommand)
	rootCmd.AddCommand(sshCommand)
	rootCmd.AddCommand(sftpCommand)
	rootCmd.AddCommand(envCommand)
	rootCmd.AddCommand(bbhjCommand)
	rootCmd.AddCommand(hostCommand)
	rootCmd.AddCommand(releaseCommand)
	rootCmd.AddCommand(mysqlCommand)
	rootCmd.AddCommand(updateCommand)
	rootCmd.AddCommand(manCommand)
	manCommand.AddCommand(gitInitCommand)
	manCommand.AddCommand(dockerInitCommand)
	manCommand.AddCommand(cloudInitCommand)
	manCommand.AddCommand(toolsInitCommand)
	manCommand.AddCommand(brewInitCommand)
	manCommand.AddCommand(githubInitCommand)

	rootCmd.AddCommand(loginCommand)
	rootCmd.AddCommand(weatherCommand)
	rootCmd.AddCommand(wikiCommand)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

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
	rootCmd.AddCommand(osinitCommand)
	osinitCommand.AddCommand(gitInitCommand)
	osinitCommand.AddCommand(dockerInitCommand)
	osinitCommand.AddCommand(cloudInitCommand)
	osinitCommand.AddCommand(toolsInitCommand)

	rootCmd.AddCommand(loginCommand)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

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
	rootCmd.AddCommand(completionBashCommand)
	rootCmd.AddCommand(manCommand)
	manCommand.AddCommand(gitInitCommand)
	manCommand.AddCommand(dockerInitCommand)
	manCommand.AddCommand(cloudInitCommand)
	manCommand.AddCommand(toolsInitCommand)
	manCommand.AddCommand(brewInitCommand)
	manCommand.AddCommand(githubInitCommand)
	manCommand.AddCommand(vimInitCommand)

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

// completionCmd represents the completion command
var completionBashCommand = &cobra.Command{
	Use:   "completion",
	Short: "Generates bash completion scripts",
	Long: `To load completion run

. <(bitbucket completion)

To configure your bash shell to load completions for each session add to your bashrc

# MacOS:
# adb completion >/usr/local/etc/bash_completion.d/adb
# ~/.bashrc or ~/.profile
. <(bitbucket completion)
`,
	Run: func(cmd *cobra.Command, args []string) {
		err := rootCmd.GenBashCompletion(os.Stdout)
		if err != nil {
			fmt.Println("Generates bash completion scripts failed!")
		}
	},
}

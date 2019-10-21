package cmd

import (
	"fmt"
	"github.com/airdb/adb/airlib"
	"github.com/spf13/cobra"
)

var initAirdbCommand = &cobra.Command{
	Use:                "init",
	Short:              "init",
	Long:               "init",
}

var initSSHKeyCommand = &cobra.Command{
	Use:                "ssh",
	Short:              "add sshkey",
	Long:               "add sshkey",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		airlib.SetupSSHPublicKey()
	},
}

var initGitCommand = &cobra.Command{
	Use:                "git",
	Short:              "init git configuration",
	Long:               "init git configuration",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("docker git")
	},
}

var initDockerCommand = &cobra.Command{
	Use:                "docker",
	Short:              "init docker os",
	Long:               "init docker os",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("docker init")
	},
}

var initMacCommand = &cobra.Command{
	Use:                "mac",
	Short:              "init MacOS",
	Long:               "init MacOS",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mac init")
	},
}

func initAirdb() {
	initAirdbCommand.AddCommand(initSSHKeyCommand)
	initAirdbCommand.AddCommand(initGitCommand)
	initAirdbCommand.AddCommand(initDockerCommand)
	initAirdbCommand.AddCommand(initMacCommand)
}

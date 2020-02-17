package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var manCommand = &cobra.Command{
	Use:   "man",
	Short: "man command like linux",
	Long:  "display manual pages",
}

var gitInitCommand = &cobra.Command{
	Use:   "git",
	Short: "git operation",
	Long:  "git operation",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("git config --global core.hooksPath .github/hooks")
	},
}

var dockerInitCommand = &cobra.Command{
	Use:   "docker",
	Short: "docker operation",
	Long:  "docker operation",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("docker exec -it -e COLUMNS=`tput cols` -e LINES=`tput lines`  -it airdb/go bash")
		fmt.Println("docker exec -it -e COLUMNS=$(tput cols) -e LINES=$(tput lines) airdb/go bash")
	},
}

var cloudInitCommand = &cobra.Command{
	Use:   "cloud",
	Short: "cloud operation",
	Long:  "cloud operation",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("brew install aliyun-cli")
		fmt.Println("https://github.com/aliyun/aliyun-cli")
	},
}

var toolsInitCommand = &cobra.Command{
	Use:   "tools",
	Short: "install binary tools",
	Long:  "install binary tools",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("wget https://github.com/airdb/init/releases/latest/download/tools-linux-amd64.zip")
	},
}

var brewInitCommand = &cobra.Command{
	Use:   "brew",
	Short: "brew operation",
	Long:  "brew operation",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("brew outdated")
		fmt.Println("brew tap aidb/taps")
		fmt.Println("brew install adb")
	},
}

var githubInitCommand = &cobra.Command{
	Use:   "gh",
	Short: "github operation",
	Long:  "github operation",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("brew install github/gh/gh")
		fmt.Println("gh --repo bbhj/lbs issue status")
		fmt.Println("gh --repo bbhj/lbs issue view 1")
	},
}

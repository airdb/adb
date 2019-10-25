package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var osinitCommand = &cobra.Command{
	Use:   "init",
	Short: "init operation",
	Long:  "init operation",
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
		fmt.Println("docker exec  -e COLUMNS=`tput cols` -e LINES=`tput lines`  -it airdb/go bash")
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

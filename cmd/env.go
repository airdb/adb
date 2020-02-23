package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var envCommand = &cobra.Command{
	Use:                "env",
	Short:              "show env",
	Long:               "Show Environment",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("adb_config=~/.adb/config.json")
		fmt.Println("git_hook_path=./.github/hooks/")
		fmt.Println("aliyun_config=~/.aliyun/config.json")
	},
}

package cmd

import (
	"github.com/airdb/adb/internal/adblib"
	"github.com/spf13/cobra"
)

var loginCommand = &cobra.Command{
	Use:   "login",
	Short: "login airdb",
	Long:  "login airdb",
	RunE: func(cmd *cobra.Command, args []string) error {
		return adblib.Login()
	},
}

func initLogin() {
	rootCmd.AddCommand(loginCommand)
}

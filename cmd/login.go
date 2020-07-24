package cmd

import (
	"github.com/airdb/adb/internal/adblib"
	"github.com/spf13/cobra"
)

var loginCommand = &cobra.Command{
	Use:     "login",
	Aliases: []string{"hello"},
	Short:   "login airdb",
	Long:    "login airdb",
	Run: func(cmd *cobra.Command, args []string) {
		login()
	},
}

func login() {
	adblib.Login()
}

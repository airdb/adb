package cmd

import (
	"github.com/airdb/adb/internal/adblib"
	"github.com/spf13/cobra"
)

var loginCommand = &cobra.Command{
	Use:     "login",
	Aliases: []string{"hello", "logo"},
	Short:   "login airdb",
	Long:    "login airdb",
	Run: func(cmd *cobra.Command, args []string) {
		login()
	},
}

var loginWithToken bool

func initLogin() {
	rootCmd.AddCommand(loginCommand)

	// 	hostSSHCmd.PersistentFlags().StringVarP(&sshFlags.IdentityFile,
	//	"identity_file", "i", "~/.adb/id_rsa", "identity file")
	loginCommand.PersistentFlags().BoolVarP(&loginWithToken, "token", "t", false, "login with token")
}

func login() {
	adblib.Login()
}

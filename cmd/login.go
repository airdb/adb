package cmd

import (
	"airdb.io/airdb/adb/internal/adblib"
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

var loginWithToken bool

func initLogin() {
	rootCmd.AddCommand(loginCommand)

	// 	hostSSHCmd.PersistentFlags().StringVarP(&sshFlags.IdentityFile,
	//	"identity_file", "i", "~/.adb/id_rsa", "identity file")
	loginCommand.PersistentFlags().BoolVarP(&loginWithToken, "token", "t", false, "login with token")
}

func login() {
	if loginWithToken {
		adblib.LoginWithToken()

		return
	}

	adblib.Login()
}

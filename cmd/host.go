package cmd

import (
	"strings"

	"github.com/airdb/adb/internal/adblib"
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/cobra"
)

var hostCmd = &cobra.Command{
	Use:                "host",
	Short:              "Perform actions on hosts",
	Long:               "Perform actions on hosts",
	DisableFlagParsing: true,
	Aliases:            []string{"server", "servers", "hosts"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			host()
		}
	},
}

func hostCmdInit() {
	rootCmd.AddCommand(hostCmd)
	hostCmd.AddCommand(keyListCmd)
	/*
		hostCmd.AddCommand(hostListCmd)
		hostCmd.AddCommand(hostSSHCmd)
		hostCmd.AddCommand(hostSFTPCmd)

		sshOptions := []string{
			"StrictHostKeyChecking=no",
			"UserKnownHostsFile=/dev/null",
			"ConnectTimeout=3",
		}

		hostSSHCmd.PersistentFlags().StringVarP(&sshFlags.LoginName, "login_name", "l", DefaultSSHUser, "login name")
		hostSSHCmd.PersistentFlags().StringVarP(&sshFlags.IdentityFile, "identity_file", "i",
			"~/.config/ssh/id_rsa", "identity file")

		sshFlags.Options = hostSSHCmd.PersistentFlags().StringSliceP("option", "o", sshOptions, "ssh option")
		hostSSHCmd.PersistentFlags().StringVarP(&sshFlags.SFTPDestPath, "sftp_server_path", "d", "/tmp",
			"sftp server dest path")

		hostSFTPCmd.PersistentFlags().StringVarP(&sshFlags.LoginName, "login_name", "l", DefaultSSHUser, "login name")
		sshFlags.Options = hostSFTPCmd.PersistentFlags().StringSliceP("option", "o", sshOptions, "ssh option")

		hostSFTPCmd.PersistentFlags().StringVarP(&sshFlags.IdentityFile, "identity_file", "i", "~/.config/ssh/id_rsa",
			"identity file")

		hostSFTPCmd.PersistentFlags().StringVarP(&sshFlags.SFTPDestPath, "sftp_server_path", "d", "/tmp",
			"sftp server dest path")
	*/
}

var keyListCmd = &cobra.Command{
	Use:     "keys",
	Short:   "List ssh public keys",
	Long:    "List ssh public keys",
	Aliases: []string{"key"},
	// DisableFlagParsing: true,
	Example: "adb host keys >> ~/.ssh/authorized_keys",
	Run: func(cmd *cobra.Command, args []string) {
		listPubKeys()
	},
}

func host() {
	adblib.GetHosts()
}

func listPubKeys() {
	hostAdmins := strings.Split(adblib.ConfigNew.HostUsers, ",")
	adblib.GetGithubKeys(hostAdmins)
}

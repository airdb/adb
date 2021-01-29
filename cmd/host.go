package cmd

import (
	"fmt"
	"strings"

	"github.com/airdb/sailor"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/spf13/cobra"
)

var hostCmd = &cobra.Command{
	Use:                "host",
	Short:              "Perform actions on hosts",
	Long:               "Perform actions on hosts",
	DisableFlagParsing: true,
	Aliases:            []string{"server", "servers", "hosts"},
}

func hostCmdInit() {
	rootCmd.AddCommand(hostCmd)
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
}

var hostListCmd = &cobra.Command{
	Use:                "list",
	Short:              "List servers",
	Long:               "List servers",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		host()
	},
}

func host() {
	client, err := aliyunConfigInit()
	if err != nil {
		panic(err)
	}

	request := alidns.CreateDescribeDomainRecordsRequest()
	request.DomainName = HostDomain

	output, err := client.DescribeDomainRecords(request)
	if err != nil {
		fmt.Println(err)
	}

	for _, rr := range output.DomainRecords.Record {
		if rr.RR == sailor.DelimiterStar || rr.RR == sailor.DelimiterAt {
			continue
		}

		domain := rr.RR + "." + rr.DomainName
		fmt.Printf("%-8s\t%-32s\t%s\n", rr.RR, domain, rr.Value)
	}
}

type sshStruct struct {
	LoginName    string
	IdentityFile string
	Options      *[]string
	SFTPDestPath string
}

var sshFlags = sshStruct{}

var hostSSHCmd = &cobra.Command{
	Use:                "ssh [server]",
	Short:              "SSH servers",
	Long:               "SSH servers",
	DisableFlagParsing: true,
	Args:               cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hostSSH(args)
	},
}

func hostSSH(args []string) {
	host := args[0]
	if !strings.HasSuffix(host, "."+HostDomain) {
		host = host + "." + HostDomain
	}

	parms := []string{}

	parms = append(parms, host)
	parms = append(parms, "-l"+sshFlags.LoginName)
	parms = append(parms, "-i"+sshFlags.IdentityFile)

	for _, option := range *sshFlags.Options {
		parms = append(parms, "-o"+option)
	}

	sailor.Exec(CommandSSH, parms)
}

var hostSFTPCmd = &cobra.Command{
	Use:                "sftp [server]",
	Short:              "SFTP servers",
	Long:               "SFTP servers",
	Args:               cobra.ExactArgs(1),
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		hostSFTP(args)
	},
}

func hostSFTP(args []string) {
	host := args[0]
	if !strings.HasSuffix(host, "."+HostDomain) {
		host = host + "." + HostDomain
	}

	parms := []string{}

	parms = append(parms, "-i"+sshFlags.IdentityFile)
	for _, option := range *sshFlags.Options {
		parms = append(parms, "-o"+option)
	}

	sftpTarget := fmt.Sprintf("%s@%s:%s",
		sshFlags.LoginName,
		host,
		sshFlags.SFTPDestPath,
	)

	// Sftp target(user@host:/tmp) must at the end of params.
	parms = append(parms, sftpTarget)

	sailor.Exec(CommandSFTP, parms)
}

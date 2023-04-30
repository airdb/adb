package cmd

import (
	"fmt"

	"github.com/airdb/adb/internal/adblib"
	"github.com/spf13/cobra"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	lh "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"
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
	getLightHouse(regions.Singapore)
	getHostByRegion(regions.Shanghai)
	getHostByRegion(regions.Singapore)
	getHostByRegion(regions.Nanjing)
}

func getHostByRegion(region string) {
	credential := common.NewCredential(adblib.AdbConfig.TencentyunAccessKeyID, adblib.AdbConfig.TencentyunAccessKeySecret)
	client, _ := cvm.NewClient(credential, region, profile.NewClientProfile())

	request := cvm.NewDescribeInstancesRequest()
	output, err := client.DescribeInstances(request)

	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)

		return
	}

	for _, instance := range output.Response.InstanceSet {
		fmt.Printf("%s\t%s\t%s\t%s\t%s\t%s\n",
			*instance.InstanceId,
			*instance.ExpiredTime,
			region,
			*instance.InstanceName,
			*instance.PublicIpAddresses[0],
			*instance.PrivateIpAddresses[0],
		)
	}
}

func getLightHouse(region string) {
	credential := common.NewCredential(adblib.AdbConfig.TencentyunAccessKeyID, adblib.AdbConfig.TencentyunAccessKeySecret)
	client, _ := lh.NewClient(credential, region, profile.NewClientProfile())

	request := lh.NewDescribeInstancesRequest()
	output, err := client.DescribeInstances(request)

	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)

		return
	}

	for _, instance := range output.Response.InstanceSet {
		fmt.Printf("%s\t%s\t%s\t%s\t%s\t%s\n",
			*instance.InstanceId,
			*instance.ExpiredTime,
			region,
			*instance.InstanceName,
			*instance.PublicAddresses[0],
			*instance.PrivateAddresses[0],
		)
	}
}

var hostAdmins = []string{
	"deancn",
	"yino",
	"bumu",
	"hallelujah-shih",
	"xqbumu",
	"lovezsr",
	"phuslu",
	"wekeey",
	"hsluoyz",
	"carrot1234567",
	"servens",
	"cateleon",
	"cellan",
	"yuduxing",
}

func listPubKeys() {
	adblib.GetGithubKeys(hostAdmins)
}

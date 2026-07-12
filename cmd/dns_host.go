package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/airdb/adb/internal/adblib"
	"github.com/miekg/dns"
	"github.com/spf13/cobra"
)

var hostCmd = &cobra.Command{
	Use:     "host",
	Short:   "Perform actions on hosts",
	Long:    "Perform actions on hosts",
	Aliases: []string{"server", "servers", "hosts"},
	RunE: func(cmd *cobra.Command, args []string) error {
		return host()
	},
}

func hostCmdInit() {
	rootCmd.AddCommand(hostCmd)
	hostCmd.AddCommand(hostDeleteCmd)
	hostCmd.AddCommand(keyListCmd)
}

var keyListCmd = &cobra.Command{
	Use:     "keys",
	Short:   "List ssh public keys",
	Long:    "List ssh public keys",
	Aliases: []string{"key"},
	Example: "adb host keys >> ~/.ssh/authorized_keys",
	RunE: func(cmd *cobra.Command, args []string) error {
		return listPubKeys()
	},
}

var hostDeleteCmd = &cobra.Command{
	Use:   "delete [hostid]",
	Short: "Delete hostid",
	Long:  "Delete hostid",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return deleteRecord(args[0])
	},
}

func listPubKeys() error {
	if adblib.ConfigNew.HostUsers == "" {
		return errors.New("host_users is not configured, set it in ~/.config/adb/config.json")
	}

	hostAdmins := strings.Split(adblib.ConfigNew.HostUsers, ",")

	return adblib.GetGithubKeys(hostAdmins)
}

func host() error {
	records, err := describeRecords(HostDomain)
	if err != nil {
		return err
	}

	for _, rr := range records {
		if rr.Type == dns.TypeToString[dns.TypeA] {
			fmt.Printf("%-20s %-5s %-32s %-64s %s\n", rr.RecordId, rr.Type, rr.RR, rr.Value, rr.Remark)
		}
	}

	return nil
}

package cmd

import (
	"fmt"

	"github.com/airdb/toolbox/typeutil"
	"github.com/miekg/dns"
	"github.com/spf13/cobra"
)

// dnsTxtCmd represents the dns txt command.
var dnsTxtCmd = &cobra.Command{
	Use:     "txt",
	Short:   "Airdb dns txt client",
	Long:    `Airdb dns txt client`,
	Aliases: []string{"text"},
	RunE: func(cmd *cobra.Command, args []string) error {
		return dnsTxt()
	},
}

func dnsTxtCmdInit() {
	rootCmd.AddCommand(dnsTxtCmd)
}

func dnsTxt() error {
	records, err := describeRecords(ServiceDomain)
	if err != nil {
		return err
	}

	for _, rr := range records {
		if rr.Type != dns.TypeToString[dns.TypeTXT] {
			continue
		}

		if rr.RR == typeutil.DelimiterStar || rr.RR == typeutil.DelimiterAt {
			continue
		}

		fmt.Printf("%-20s %-32s %-64s %s\n", rr.RecordId, rr.RR, rr.Value, rr.Remark)
	}

	return nil
}

package cmd

import (
	"errors"
	"fmt"

	"github.com/airdb/adb/internal/adblib"
	"github.com/airdb/toolbox/typeutil"
	"github.com/miekg/dns"
	"github.com/spf13/cobra"
)

// serviceCmd represents the service command.
var serviceCmd = &cobra.Command{
	Use:     "srv",
	Short:   "Airdb service client",
	Long:    `Airdb service client`,
	Aliases: []string{"service", "services"},
	RunE: func(cmd *cobra.Command, args []string) error {
		return service()
	},
}

func serviceCmdInit() {
	rootCmd.AddCommand(serviceCmd)
	serviceCmd.AddCommand(servicesAddCmd)
	serviceCmd.AddCommand(servicesUpdateCmd)
	serviceCmd.AddCommand(servicesDeleteCmd)

	servicesUpdateCmd.PersistentFlags().StringVarP(&updateDNSFlag.RecordID,
		"id", "i", "", "srv record_id")
	servicesUpdateCmd.PersistentFlags().StringVarP(&updateDNSFlag.Remark,
		"remark", "m", "", "srv remark or comment")
}

type AliDNSStruct struct {
	RecordID string
	RR       string
	Value    string
	Remark   string
}

var updateDNSFlag AliDNSStruct

func service() error {
	records, err := describeRecords(ServiceDomain)
	if err != nil {
		return err
	}

	for _, rr := range records {
		if rr.Type != dns.TypeToString[dns.TypeSRV] {
			continue
		}

		if rr.RR == typeutil.DelimiterStar || rr.RR == typeutil.DelimiterAt {
			continue
		}

		fmt.Printf("%-20s %-32s %-64s %s\n", rr.RecordId, rr.RR, rr.Value, rr.Remark)
	}

	return nil
}

var servicesAddCmdMinArgs = 2

var servicesAddCmd = &cobra.Command{
	Use:     "add [service] [SRV record value]",
	Short:   "Add service",
	Long:    "Add service",
	Example: adblib.DNSSrvDoc,
	Args:    cobra.MinimumNArgs(servicesAddCmdMinArgs),
	RunE: func(cmd *cobra.Command, args []string) error {
		return addRecord(ServiceDomain, dns.TypeToString[dns.TypeSRV], args[0], args[1])
	},
}

var servicesUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update service remark",
	Long:  "Update service remark",
	RunE: func(cmd *cobra.Command, args []string) error {
		if updateDNSFlag.RecordID == "" {
			return errors.New("flag --id is required")
		}

		return updateRecordRemark(updateDNSFlag.RecordID, updateDNSFlag.Remark)
	},
}

var servicesDeleteCmd = &cobra.Command{
	Use:   "delete [record_id]",
	Short: "Delete service",
	Long:  "Delete service",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return deleteRecord(args[0])
	},
}

/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/airdb/adb/internal/adblib"
	"github.com/airdb/toolbox/typeutil"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/miekg/dns"
	"github.com/spf13/cobra"
)

// serviceCmd represents the service command.
var serviceCmd = &cobra.Command{
	Use:     "srv",
	Short:   "Airdb service client",
	Long:    `Airdb service client`,
	Aliases: []string{"service", "services"},
	Run: func(cmd *cobra.Command, args []string) {
		service()
	},
}

func serviceCmdInit() {
	rootCmd.AddCommand(serviceCmd)
	serviceCmd.AddCommand(servicesAddCmd)
	serviceCmd.AddCommand(servicesUpdateCmd)
	serviceCmd.AddCommand(servicesDeleteCmd)

	serviceCmd.PersistentFlags().BoolVarP(&serviceFlags.List, "list", "l", false, "list all services")
	servicesUpdateCmd.PersistentFlags().StringVarP(&updateDNSFlag.RecordID,
		"id", "i", "", "srv record_id")
	servicesUpdateCmd.PersistentFlags().StringVarP(&updateDNSFlag.Remark,
		"remark", "m", "", "srv remark or comment")
}

type serviceStruct struct {
	List bool
}

type AliDNSStruct struct {
	RecordID string
	RR       string
	Value    string
	Remark   string
}

var updateDNSFlag AliDNSStruct

var serviceFlags = serviceStruct{}

func service() {
	client, err := aliyunConfigInit()
	if err != nil {
		panic(err)
	}

	request := alidns.CreateDescribeDomainRecordsRequest()
	request.DomainName = ServiceDomain

	output, err := client.DescribeDomainRecords(request)
	if err != nil {
		fmt.Println(err)
	}

	for _, rr := range output.DomainRecords.Record {
		if rr.RR == typeutil.DelimiterStar || rr.RR == typeutil.DelimiterAt {
			continue
		}

		// fmt.Printf("%-20s\t%-32s\t%s\n", rr.RecordId, rr.RR, rr.Value)
		fmt.Printf("%-20s %-32s %-64s %s\n", rr.RecordId, rr.RR, rr.Value, rr.Remark)
	}
}

var servicesAddCmdMinArgs = 2

var servicesAddCmd = &cobra.Command{
	Use:     "add [service] [SRV record value]",
	Short:   "Add service",
	Long:    "Add service",
	Example: adblib.DNSSrvDoc,
	Args:    cobra.MinimumNArgs(servicesAddCmdMinArgs),
	Run: func(cmd *cobra.Command, args []string) {
		addService(args)
	},
}

func addService(args []string) {
	client, err := aliyunConfigInit()
	if err != nil {
		panic(err)
	}

	request := alidns.CreateAddDomainRecordRequest()
	request.DomainName = ServiceDomain
	request.Type = dns.TypeToString[dns.TypeSRV]
	request.RR = args[0]
	request.Value = args[1]

	output, err := client.AddDomainRecord(request)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(output)
}

var servicesUpdateCmd = &cobra.Command{
	Use:   "update [service]",
	Short: "Update service",
	Long:  "Update service",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		updateService()
	},
}

var servicesDeleteCmd = &cobra.Command{
	Use:   "delete [service]",
	Short: "Delete service",
	Long:  "Delete service",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		deleteService(args)
	},
}

func updateService() {
	client, err := aliyunConfigInit()
	if err != nil {
		panic(err)
	}

	request := alidns.CreateUpdateDomainRecordRemarkRequest()
	request.RecordId = updateDNSFlag.RecordID
	request.Remark = updateDNSFlag.Remark

	output, err := client.UpdateDomainRecordRemark(request)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(output)
}

func deleteService(args []string) {
	client, err := aliyunConfigInit()
	if err != nil {
		panic(err)
	}

	request := alidns.CreateDeleteDomainRecordRequest()
	request.RecordId = args[0]

	output, err := client.DeleteDomainRecord(request)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(output)
}

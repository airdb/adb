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
	"github.com/airdb/sailor"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
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
	serviceCmd.AddCommand(servicesDeleteCmd)

	serviceCmd.PersistentFlags().BoolVarP(&serviceFlags.List, "list", "l", false, "list all services")
}

type serviceStruct struct {
	List bool
}

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
		if rr.RR == sailor.DelimiterStar || rr.RR == sailor.DelimiterAt {
			continue
		}

		fmt.Printf("%-20s\t%-32s\t%s\n", rr.RecordId, rr.RR, rr.Value)
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
	request.Type = "SRV"
	request.RR = args[0]
	request.Value = args[1]

	output, err := client.AddDomainRecord(request)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(output)
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

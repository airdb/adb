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

	"github.com/airdb/toolbox/typeutil"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/spf13/cobra"
)

// dnsTxtCmd represents the service command.
var dnsTxtCmd = &cobra.Command{
	Use:     "txt",
	Short:   "Airdb dns txt client",
	Long:    `Airdb dns txt client`,
	Aliases: []string{"service", "services"},
	Run: func(cmd *cobra.Command, args []string) {
		dnsTxt()
	},
}

func dnsTxtCmdInit() {
	rootCmd.AddCommand(dnsTxtCmd)
	/*
		dnsTxtCmd.AddCommand(servicesAddCmd)
		dnsTxtCmd.AddCommand(servicesUpdateCmd)
		dnsTxtCmd.AddCommand(servicesDeleteCmd)

		dnsTxtCmd.PersistentFlags().BoolVarP(&serviceFlags.List, "list", "l", false, "list all services")
		servicesUpdateCmd.PersistentFlags().StringVarP(&updateDNSFlag.RecordID,
			"id", "i", "", "srv record_id")
		servicesUpdateCmd.PersistentFlags().StringVarP(&updateDNSFlag.Remark,
			"remark", "m", "", "srv remark or comment")
	*/
}

type dnsTxtStruct struct {
	List bool
}

func dnsTxt() {
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
		if rr.Type != "TXT" {
			continue
		}

		if rr.RR == typeutil.DelimiterStar || rr.RR == typeutil.DelimiterAt {
			continue
		}

		// fmt.Printf("%-20s\t%-32s\t%s\n", rr.RecordId, rr.RR, rr.Value)
		fmt.Printf("%-20s %-32s %-64s %s\n", rr.RecordId, rr.RR, rr.Value, rr.Remark)
	}
}

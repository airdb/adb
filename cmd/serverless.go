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
	"github.com/spf13/cobra"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"

	apigateway "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/apigateway/v20180808"
)

// serverlessCmd represents the serverless command.
var serverlessCmd = &cobra.Command{
	Use:     "serverless",
	Short:   "Airdb serverless client",
	Long:    `Airdb serverless client`,
	Aliases: []string{"sls"},
	Run: func(cmd *cobra.Command, args []string) {
		serverless()
	},
}

func serverlessCmdInit() {
	rootCmd.AddCommand(serverlessCmd)

	serverlessCmd.PersistentFlags().BoolVarP(&serverlessFlags.List, "list", "l", false, "list all serverlesss")
}

type serverlessStruct struct {
	List bool
}

var serverlessFlags = serverlessStruct{}

func serverless() {
	config := adblib.GetTencentYunConfig()

	credential := common.NewCredential(config.AccessKeyID, config.AccessKeySecret)
	client, _ := apigateway.NewClient(credential, regions.Shanghai, profile.NewClientProfile())

	request := apigateway.NewDescribeServicesStatusRequest()
	response, err := client.DescribeServicesStatus(request)

	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%s\n", response.ToJsonString())

	request1 := apigateway.NewDescribeServiceRequest()
	request1.ServiceId = response.Response.Result.ServiceSet[0].ServiceId
	response1, err := client.DescribeService(request1)

	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}

	for _, name := range response1.Response.ApiIdStatusSet {
		fmt.Printf("%v\t%v\t%v\n", *name.ModifiedTime, *name.Path, *name.ApiName)
	}
}

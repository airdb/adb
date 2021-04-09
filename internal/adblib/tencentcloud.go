package adblib

import (
	"errors"
	"fmt"
	"os"

	clb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/clb/v20180317"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	terrors "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const CLBEndpoint = "clb.tencentcloudapi.com"

type Client struct {
	*clb.Client
}

func NewCLBClient() (*Client, error) {
	secretID := os.Getenv("TENCENTCLOUD_SECRET_ID")
	secretKey := os.Getenv("TENCENTCLOUD_SECRET_KEY")
	region := os.Getenv("TENCENTCLOUD_REGION")

	credential := common.NewCredential(secretID, secretKey)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = CLBEndpoint

	client, err := clb.NewClient(credential, region, cpf)

	return &Client{
		client,
	}, err
}

func (client *Client) ListCLB() {
	request := clb.NewDescribeLoadBalancersRequest()

	params := "{}"

	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}

	response, err := client.DescribeLoadBalancers(request)

	var sdkError *terrors.TencentCloudSDKError
	if errors.As(err, &sdkError) {
		fmt.Printf("An API error has returned: %s", err)

		return
	}

	if err != nil {
		panic(err)
	}

	for _, lbSet := range response.Response.LoadBalancerSet {
		// fmt.Println(*lbSet.LoadBalancerId)
		client.ShowRS(*lbSet.LoadBalancerId)
	}
}

func (client *Client) ShowRS(lbID string) {
	request := clb.NewDescribeTargetsRequest()

	params := fmt.Sprintf("{\"LoadBalancerId\":\"%s\"}", lbID)

	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}

	response, err := client.DescribeTargets(request)

	var sdkError *terrors.TencentCloudSDKError
	if errors.As(err, &sdkError) {
		fmt.Printf("An API error has returned: %s", err)

		return
	}

	if err != nil {
		panic(err)
	}

	// fmt.Printf("%s", response.ToJsonString())
	for _, listener := range response.Response.Listeners {
		for _, rule := range listener.Rules {
			for _, target := range rule.Targets {
				fmt.Printf("%s%s\t%d\t%s\t%s\n",
					*rule.Domain,
					*rule.Url,
					*target.Weight,
					*target.InstanceId,
					*target.InstanceName,
				)
			}
		}
	}
}

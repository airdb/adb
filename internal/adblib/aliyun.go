package adblib

import "github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"

func NewAliyunClient() (*alidns.Client, error) {
	client, err := alidns.NewClientWithAccessKey(
		ConfigNew.AliyunRegionID,
		ConfigNew.AliyunAccessKeyID,
		ConfigNew.AliyunAccessKeySecret,
	)
	if err != nil {
		return nil, err
	}

	return client, nil
}

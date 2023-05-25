package adblib

import "github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"

const (
	CommandSSH     = "ssh"
	CommandSFTP    = "sftp"
	DefaultSSHUser = "ubuntu"
)

const CloudPlatformAliyun = "aliyun"

const (
	ServiceDomain = "airdb.space"
	HostDomain    = "airdb.host"
)

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

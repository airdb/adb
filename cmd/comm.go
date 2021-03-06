package cmd

import (
	"airdb.io/airdb/adb/internal/adblib"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

const (
	CommandSSH     = "ssh"
	CommandSFTP    = "sftp"
	DefaultSSHUser = "ubuntu"
)

const (
	ServiceDomain = "airdb.space"
	HostDomain    = "airdb.host"
)

const CloudPlatformAliyun = "aliyun"

// var aliyunConfig = map[string]string{}

func aliyunConfigInit() (*alidns.Client, error) {
	aliyunFlag := adblib.GetAliyunConfig()

	client, err := alidns.NewClientWithAccessKey(
		aliyunFlag.RegionID,
		aliyunFlag.AccessKeyID,
		aliyunFlag.AccessKeySecret,
	)
	if err != nil {
		return nil, err
	}

	return client, nil
}

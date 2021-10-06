package cmd

import (
	"github.com/airdb/adb/internal/adblib"
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
	client, err := alidns.NewClientWithAccessKey(
		adblib.AdbConfig.AliyunRegionID,
		adblib.AdbConfig.AliyunAccessKeyID,
		adblib.AdbConfig.AliyunAccessKeySecret,
	)
	if err != nil {
		return nil, err
	}

	return client, nil
}

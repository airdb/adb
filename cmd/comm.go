package cmd

import (
	"os"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/spf13/viper"
)

const (
	CommandSSH     = "ssh"
	CommandSFTP    = "sftp"
	DefaultSSHUser = "ubuntu"
)

const (
	ServiceDomain = "airdb.me"
	HostDomain    = "airdb.host"
)

const CloudPlatformAliyun = "aliyun"

var aliyunConfig = map[string]string{}

func aliyunConfigInit() (*alidns.Client, error) {
	configFile := os.Getenv("HOME") + "/.adb/config.json"

	viper.SetConfigFile(configFile)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.UnmarshalKey(CloudPlatformAliyun, &aliyunConfig); err != nil {
		panic(err)
	}

	regionID := aliyunConfig["region_id"]
	accessKeyID := aliyunConfig["access_key_id"]
	accessKeySecret := aliyunConfig["access_key_secret"]

	client, err := alidns.NewClientWithAccessKey(
		regionID,
		accessKeyID,
		accessKeySecret,
	)

	if err != nil {
		return nil, err
	}

	return client, nil
}

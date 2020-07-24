package config

import (
	"path"

	"github.com/mitchellh/go-homedir"
)

func ConfigDir() string {
	dir, _ := homedir.Expand("~/.config/adb")
	return dir
}

func ConfigFile() string {
	return path.Join(ConfigDir(), "config.json")
}

func IconFile() string {
	return path.Join(ConfigDir(), "icon")
}

func aliyunConfigFile() string {
	return path.Join(ConfigDir(), "aliyun.json")
}

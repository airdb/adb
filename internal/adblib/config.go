package adblib

import (
	"encoding/json"
	"path"

	"github.com/AlecAivazis/survey/v2"
	"github.com/airdb/sailor"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
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

// The questions to ask.
var qsAliyun = []*survey.Question{
	{
		Name: "access_key_id",
		// Name:     "access",
		Prompt:    &survey.Input{Message: "access_key_id"},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
	{
		Name:      "access_key_secret",
		Prompt:    &survey.Input{Message: "access_key_secret"},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
	{
		Name: "region_id",
		Prompt: &survey.Input{
			Message: "region_id",
			Default: "cn-hangzhou",
		},
	},
}

// The  flag will be written to this struct.
type AliyunFlag struct {
	AccessKeyID     string `json:"access_key_id" survey:"access_key_id" mapstructure:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret" survey:"access_key_secret" mapstructure:"access_key_secret"`
	RegionID        string `json:"region_id" survey:"region_id" mapstructure:"region_id"`
}

func GetAliyunConfig() *AliyunFlag {
	viper.SetConfigFile(aliyunConfigFile())

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var aliyunFlag AliyunFlag

	if err := viper.Unmarshal(&aliyunFlag); err != nil {
		panic(err)
	}

	return &aliyunFlag
}

func SetAliyunConfig() error {
	var aliyunFlag AliyunFlag

	// pPerform the questions.
	err := survey.Ask(qsAliyun, &aliyunFlag)
	if err != nil {
		return err
	}

	jsonByte, err := json.MarshalIndent(aliyunFlag, "", "\t")
	if err != nil {
		return err
	}

	err = sailor.WriteFile(aliyunConfigFile(), string(jsonByte))
	if err != nil {
		return err
	}

	return nil
}

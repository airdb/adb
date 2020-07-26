package adblib

import (
	"encoding/json"
	"path"

	"github.com/AlecAivazis/survey/v2"
	"github.com/airdb/sailor"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const (
	ServiceAliyun = "aliyun"
	ServiceSlack  = "slack"
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

func slackConfigFile() string {
	return path.Join(ConfigDir(), "slack.json")
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

// The questions to ask.
var qsSlack = []*survey.Question{
	{
		Name:     "token",
		Prompt:   &survey.Input{Message: "token"},
		Validate: survey.Required,
	},
	{
		Name: "channel",
		Prompt: &survey.Input{Message: "channel",
			Default: "#wiki",
		},
		Validate: survey.Required,
	},
}

// The  flag will be written to this struct.
type AliyunFlag struct {
	AccessKeyID     string `json:"access_key_id" survey:"access_key_id" mapstructure:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret" survey:"access_key_secret" mapstructure:"access_key_secret"`
	RegionID        string `json:"region_id" survey:"region_id" mapstructure:"region_id"`
}

// The  flag will be written to this struct.
type SlackFlag struct {
	Token   string `json:"token" survey:"token" mapstructure:"token"`
	Channel string `json:"channel" survey:"channel" mapstructure:"channel"`
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

func GetSlackConfig() *SlackFlag {
	viper.SetConfigFile(slackConfigFile())

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var config SlackFlag

	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	return &config
}

func SetSlackConfig() error {
	var config SlackFlag

	// pPerform the questions.
	err := survey.Ask(qsSlack, &config)
	if err != nil {
		return err
	}

	jsonByte, err := json.MarshalIndent(config, "", "\t")
	if err != nil {
		return err
	}

	err = sailor.WriteFile(slackConfigFile(), string(jsonByte))
	if err != nil {
		return err
	}

	return nil
}

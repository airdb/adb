package adblib

import (
	"encoding/json"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/airdb/sailor/fileutil"
	"github.com/joho/godotenv"
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

func TencentYunConfigFile() string {
	return path.Join(ConfigDir(), "tencentyun.json")
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
		Prompt: &survey.Input{
			Message: "channel",
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

type TencentYunFlag struct {
	AccessKeyID     string `json:"access_key_id" survey:"access_key_id" mapstructure:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret" survey:"access_key_secret" mapstructure:"access_key_secret"`
	RegionID        string `json:"region_id" survey:"region_id" mapstructure:"region_id"`
}

// The  flag will be written to this struct.
type SlackFlag struct {
	Token   string `json:"token" survey:"token" mapstructure:"token"`
	Channel string `json:"channel" survey:"channel" mapstructure:"channel"`
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

	err = fileutil.WriteFile(aliyunConfigFile(), string(jsonByte))
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

	err = fileutil.WriteFile(slackConfigFile(), string(jsonByte))
	if err != nil {
		return err
	}

	return nil
}

const EnvFile = ".config/adb/env"

type CFG struct {
	TencentyunAccessKeyID     string
	TencentyunAccessKeySecret string
	TencentyunRegionID        string

	AliyunAccessKeyID     string
	AliyunAccessKeySecret string
	AliyunRegionID        string

	SlackToken   string
	SlackChannel string
}

var AdbConfig CFG

func GetEnvFile() string {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	envfile := filepath.Join(homedir, EnvFile)

	return envfile
}

func InitDotEnv() {
	err := godotenv.Load(GetEnvFile())
	if err != nil {
		log.Fatal("Error loading .env file ", EnvFile, err)
	}

	AdbConfig = CFG{
		TencentyunAccessKeyID:     os.Getenv("TencentyunAccessKeyID"),
		TencentyunAccessKeySecret: os.Getenv("TencentyunAccessKeySecret"),
		TencentyunRegionID:        os.Getenv("TencentyunRegionID"),
		AliyunAccessKeyID:         os.Getenv("AliyunAccessKeyID"),
		AliyunAccessKeySecret:     os.Getenv("AliyunAccessKeySecret"),
		AliyunRegionID:            os.Getenv("AliyunRegionID"),
		SlackToken:                os.Getenv("SlackToken"),
		SlackChannel:              os.Getenv("SlackChannel"),
	}
}

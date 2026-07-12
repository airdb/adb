package adblib

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	AuthIssuer string `json:"auth_issuer" mapstructure:"auth_issuer"`
	ClientID   string `json:"client_id" mapstructure:"client_id"`

	AliyunAccessKeyID     string `json:"aliyun_access_key_id" mapstructure:"aliyun_access_key_id"`
	AliyunAccessKeySecret string `json:"aliyun_access_key_secret" mapstructure:"aliyun_access_key_secret"`

	HostUsers string `json:"host_users" mapstructure:"HostUsers"`
}

const EnvFile = ".config/adb/env"

type CFG struct {
	AuthIssuer string
	ClientID   string

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

func GetEnvFile() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homedir, EnvFile), nil
}

func InitDotEnv() error {
	envFile, err := GetEnvFile()
	if err != nil {
		return err
	}

	if _, err := os.Stat(envFile); os.IsNotExist(err) {
		return nil
	}

	if err := godotenv.Load(envFile); err != nil {
		return fmt.Errorf("load env file %s: %w", envFile, err)
	}

	AdbConfig = CFG{
		AuthIssuer: os.Getenv("AuthIssuer"),
		ClientID:   os.Getenv("CLIENT_ID"),

		TencentyunAccessKeyID:     os.Getenv("TencentyunAccessKeyID"),
		TencentyunAccessKeySecret: os.Getenv("TencentyunAccessKeySecret"),
		TencentyunRegionID:        os.Getenv("TencentyunRegionID"),
		AliyunAccessKeyID:         os.Getenv("AliyunAccessKeyID"),
		AliyunAccessKeySecret:     os.Getenv("AliyunAccessKeySecret"),
		AliyunRegionID:            os.Getenv("AliyunRegionID"),
		SlackToken:                os.Getenv("SlackToken"),
		SlackChannel:              os.Getenv("SlackChannel"),
	}

	return nil
}

var ConfigNew = &Config{}

// Init loads ~/.config/adb/config.json into ConfigNew. A missing config file
// is not an error so that commands which need no config keep working.
func Init() error {
	viper.AddConfigPath("$HOME/.config/adb/")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		var notFound viper.ConfigFileNotFoundError
		if !errors.As(err, &notFound) {
			return fmt.Errorf("read config file: %w", err)
		}

		return nil
	}

	if err := viper.Unmarshal(ConfigNew); err != nil {
		return fmt.Errorf("parse config file: %w", err)
	}

	return nil
}

package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"

	"github.com/AlecAivazis/survey/v2"
	"github.com/MakeNowJust/heredoc"
	"github.com/airdb/sailor"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func ConfigDir() string {
	dir, _ := homedir.Expand("~/.config/adb")
	return dir
}

func ConfigFile() string {
	return path.Join(ConfigDir(), "config.json")
}

func aliyunConfigFile() string {
	return path.Join(ConfigDir(), "aliyun.json")
}

var ReadConfigFile = func(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, pathError(err)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return data, nil
}

var WriteConfigFile = func(filename string, data []byte) error {
	err := os.MkdirAll(path.Dir(filename), 0771)
	if err != nil {
		return pathError(err)
	}

	cfgFile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600) // cargo coded from setup
	if err != nil {
		return err
	}
	defer cfgFile.Close()

	n, err := cfgFile.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}

	return err
}

var BackupConfigFile = func(filename string) error {
	return os.Rename(filename, filename+".bak")
}

func pathError(err error) error {
	return err
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration for adb",
	Long: `Display or change configuration settings for adb.
Current respected settings:
- git_protocol: "https" or "ssh". Default is "https".
- editor: if unset, defaults to environment variables.
`,
}

var configGetCmd = &cobra.Command{
	Use:   "get <key>",
	Short: "Print the value of a given configuration key",
	Example: heredoc.Doc(`
	$ adb config get aliyun 
	`),
	Args: cobra.ExactArgs(1),
	RunE: configGet,
}

var configSetCmd = &cobra.Command{
	Use:   "set <key> <value>",
	Short: "Update configuration with a value for the given key",
	Example: heredoc.Doc(`
	$ adb config set aliyun
    [aliyun]
    ? access_key_id: xxxxxxxxxxxx_id
    ? access_key_secret: xxxxxxxxxxxx_secret
    ? region_id: (cn-hangzhou) 
	`),
	Args: cobra.ExactArgs(1),
	RunE: configSet,
}

func initConfigCmd() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configGetCmd)
	configCmd.AddCommand(configSetCmd)
}

// The questions to ask.
var qs = []*survey.Question{
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

func configSet(cmd *cobra.Command, args []string) error {
	fmt.Println(args)

	var aliyunFlag AliyunFlag

	// pPerform the questions.
	err := survey.Ask(qs, &aliyunFlag)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	jsonByte, err := json.MarshalIndent(aliyunFlag, "", "\t")
	if err != nil {
		return err
	}

	err = sailor.WriteFile(aliyunConfigFile(), string(jsonByte))
	if err != nil {
		fmt.Println("configure failed, error: ", err)
		return err
	}

	fmt.Println("configure successfully.")

	return nil
}

func configGet(cmd *cobra.Command, args []string) error {
	viper.SetConfigFile(aliyunConfigFile())

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var aliyunFlag AliyunFlag

	// if err := viper.UnmarshalKey(CloudPlatformAliyun, &aliyunFlag); err != nil {
	// 	panic(err)
	// }
	if err := viper.Unmarshal(&aliyunFlag); err != nil {
		panic(err)
	}

	// fmt.Println(viper.AllSettings())

	fmt.Printf("access_key_id: %s\naccess_key_secret: %s\nregion_id: %s\n",
		aliyunFlag.AccessKeyID,
		aliyunFlag.AccessKeySecret,
		aliyunFlag.RegionID,
	)

	return nil
}

func getAliyunConfig() *AliyunFlag {
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

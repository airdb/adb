package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"

	"airdb.io/airdb/adb/internal/adblib"
	"github.com/MakeNowJust/heredoc"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

func ConfigDir() string {
	dir, _ := homedir.Expand("~/.config/adb")

	return dir
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
	err := os.MkdirAll(path.Dir(filename), 0o771)
	if err != nil {
		return pathError(err)
	}

	cfgFile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o600) // cargo coded from setup
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
    access_key_id: xxxxxxxxxxxx_id
    access_key_secret: xxxxxxxxxxxx_secret
    region_id: cn-hangzhou
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

func configSet(cmd *cobra.Command, args []string) error {
	service := args[0]

	switch service {
	case adblib.ServiceAliyun:
		err := adblib.SetAliyunConfig()
		if err != nil {
			fmt.Println("configure failed, error: ", err)

			return err
		}
	case adblib.ServiceSlack:
		err := adblib.SetSlackConfig()
		if err != nil {
			fmt.Println("configure failed, error: ", err)

			return err
		}
	}

	fmt.Println("configure successfully.")

	return nil
}

func configGet(cmd *cobra.Command, args []string) error {
	service := args[0]

	var config interface{}

	switch service {
	case "aliyun":
		aliyunFlag := adblib.GetAliyunConfig()
		fmt.Printf("access_key_id: %s\naccess_key_secret: %s\nregion_id: %s\n",
			aliyunFlag.AccessKeyID,
			aliyunFlag.AccessKeySecret,
			aliyunFlag.RegionID,
		)
	case "slack":
		config = adblib.GetSlackConfig()
	}

	fmt.Println(config)

	return nil
}

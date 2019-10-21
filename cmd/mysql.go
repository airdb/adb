package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var mysqlCommand = &cobra.Command{
	Use:                "mysql",
	Short:              "mysql client",
	Long:               "Airdb mysql client",
	DisableFlagParsing: true,
	Aliases:            []string{"sql"},
	Run: func(cmd *cobra.Command, args []string) {
		mysql(args)
	},
}

func mysql(args []string) {
	viper.SetConfigFile("conf/dev.json")
	if len(args) != 0 {
		viper.SetConfigFile(args[0])
	}

	databases := make(map[string]*DatabaseItem)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
		return
	}

	if err := viper.UnmarshalKey("databases", &databases); err != nil {
		log.Fatal(err)
		return
	}

	for name := range databases {
		if databases[name].DefaultTableName {
			host, port, _ := net.SplitHostPort(databases[name].Address)
			aa := fmt.Sprintf("-h%s -P%s -u%s -p%s %s", host, port, databases[name].User, databases[name].Password, name)
			mysqlcmd(strings.Split(aa, " "))
		}
	}
}

func mysqlcmd(args []string) {
	mysqlPath, err := exec.LookPath("mysql")
	if err != nil {
		return
	}

	cmd := exec.Command(mysqlPath, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		return
	}
	err = cmd.Wait()
	if err != nil {
		fmt.Printf("exec error %v\n", err)
		if exiterror, ok := err.(*exec.ExitError); ok {
			os.Exit(exiterror.ExitCode())
		}
	}
}

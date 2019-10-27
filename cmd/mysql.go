package cmd

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"

	"github.com/airdb/sailor/config"
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
	for dbname, item := range config.GetDatabases() {
		host, port, _ := net.SplitHostPort(item.Address)
		if item.Default {
			aa := fmt.Sprintf("-h%s -P%s -u%s -p%s %s", host, port, item.User, item.Password, dbname)
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
		log.Println("adb exec failed.")
		if exiterror, ok := err.(*exec.ExitError); ok {
			os.Exit(exiterror.ExitCode())
		}
	}
}

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"strings"
)

var sshCommand = &cobra.Command{
	Use:                "ssh",
	Short:              "ssh client",
	Long:               "Airdb ssh client",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		ssh(args)
	},
}

const domainZone = "airdb.host"

type DatabaseItem struct {
	User             string
	Password         string
	Address          string
	Name             string
	DefaultTableName bool `mapstructure:"default_table_name"`
}

func ssh(args []string) {
	sshPath, err := exec.LookPath("ssh")
	if err != nil {
		return
	}


	sshArgs := getArgs(args[0])
	cmd := exec.Command(sshPath, sshArgs...)
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

func getArgs(arg string) []string {
	// user = "root"
	user := "ubuntu"
	host := arg
	args := strings.Split(arg, "@")

	if len(args) == 2 {
		user = args[0]
		host = args[1]
	}

	if !strings.HasSuffix(host, "."+domainZone) {
			host = host + "." + domainZone
	}

	sshArgs := []string{
		"-i~/.adb/id_rsa",
		"-oStrictHostKeyChecking=no",
		"-oUserKnownHostsFile=/dev/null",
		"-oConnectTimeout=3",
		"-l" + user,
		host,
	}

	return sshArgs
}

package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
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

var sftpCommand = &cobra.Command{
	Use:                "sftp",
	Short:              "sftp client",
	Long:               "Airdb sftp client",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		sftp(args)
	},
}

const (
	CommandSSH     = "ssh"
	CommandSFTP    = "sftp"
	DefaultSSHUser = "ubuntu"
)

const domainZone = "airdb.host"

type DatabaseItem struct {
	User             string
	Password         string
	Address          string
	Name             string
	DefaultTableName bool `mapstructure:"default_table_name"`
}

func ssh(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage:")
		fmt.Println("  adb ssh [ip|host|container]")
		return
	}

	sshPath, err := exec.LookPath(CommandSSH)
	if err != nil {
		return
	}

	sshArgs := getArgs(CommandSSH, args[0])
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

func sftp(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage:")
		fmt.Println("  adb sftp [ip|host|container]")
		return
	}

	sshPath, err := exec.LookPath(CommandSFTP)
	if err != nil {
		return
	}

	sshArgs := getArgs(CommandSFTP, args[0])
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

func getArgs(typ, arg string) []string {
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
	}

	switch typ {
	case CommandSFTP:
		sshArgs = append(sshArgs, user+"@"+host+":/tmp")
	case CommandSSH:
		sshArgs = append(sshArgs, "-l"+user)
		sshArgs = append(sshArgs, host)
	}

	return sshArgs
}

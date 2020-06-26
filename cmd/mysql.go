package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/miekg/dns"
	"github.com/spf13/cobra"
)

var mysqlCmd = &cobra.Command{
	Use:                "mysql",
	Short:              "mysql client",
	Long:               "Airdb mysql client",
	DisableFlagParsing: false,
	Aliases:            []string{"sql"},
	Run: func(cmd *cobra.Command, args []string) {
		mysql(args)
	},
}

type mysqlStruct struct {
	Host     string
	Port     uint16
	User     string
	Password string
	DB       string
}

var mysqlFlags mysqlStruct

func mysqlCmdInit() {
	rootCmd.AddCommand(mysqlCmd)

	mysqlCmd.PersistentFlags().StringVarP(&mysqlFlags.User, "user", "u", "root", "database username")
	mysqlCmd.PersistentFlags().StringVarP(&mysqlFlags.Password, "password", "p", "airdb.me", "database password")
	mysqlCmd.PersistentFlags().StringVarP(&mysqlFlags.DB, "db", "", "test", "database name")
}

func mysql(args []string) {
	fmt.Println("args: ", args)

	c := dns.Client{Timeout: 1 * time.Second}
	m := dns.Msg{}

	m.SetQuestion("hello.airdb.me.", dns.TypeSRV)
	r, _, err := c.Exchange(&m, "8.8.8.8:53")
	fmt.Println(r.Answer, err)

	for _, ans := range r.Answer {
		record, isType := ans.(*dns.SRV)
		if isType {
			mysqlFlags.Host = record.Target
			mysqlFlags.Port = record.Port
			fmt.Println(record.Target, record.Port)
		}
	}

	flags := fmt.Sprintf("-h%s -P%d -u%s -p%s %s",
		mysqlFlags.Host,
		mysqlFlags.Port,
		mysqlFlags.User,
		mysqlFlags.Password,
		mysqlFlags.DB,
	)

	args = strings.Split(flags, " ")
	mysqlcmd(args)
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

package cmd

import (
	"fmt"
	"strings"

	"airdb.io/airdb/adb/internal/adblib"
	"github.com/airdb/sailor"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/spf13/cobra"
)

var mysqlCmd = &cobra.Command{
	Use:                "mysql [service]",
	Short:              "mysql client",
	Long:               "Airdb mysql client",
	DisableFlagParsing: false,
	Args:               cobra.MinimumNArgs(1),
	Example:            adblib.SQLDoc,
	Aliases:            []string{"sql"},
	Run: func(cmd *cobra.Command, args []string) {
		mysql(args)
	},
}

type mysqlStruct struct {
	Host     string
	Port     string
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
	client, err := aliyunConfigInit()
	if err != nil {
		panic(err)
	}

	request := alidns.CreateDescribeDomainRecordsRequest()
	request.DomainName = ServiceDomain
	request.RRKeyWord = args[0]

	fmt.Println(request.RRKeyWord)

	output, err := client.DescribeDomainRecords(request)
	if err != nil {
		fmt.Println(err)
	}

	rrs := output.DomainRecords.Record
	if len(rrs) > 1 {
		for _, rr := range output.DomainRecords.Record {
			fmt.Printf("%-32s\t%s\n", rr.RR, rr.Value)
		}
	}

	values := strings.Split(rrs[0].Value, " ")

	mysqlFlags.Host = values[3]
	mysqlFlags.Port = values[2]

	flags := fmt.Sprintf("-h%s -P%s -u%s -p%s %s",
		mysqlFlags.Host,
		mysqlFlags.Port,
		mysqlFlags.User,
		mysqlFlags.Password,
		mysqlFlags.DB,
	)

	args = strings.Split(flags, " ")
	sailor.Exec("mysql", args)
}

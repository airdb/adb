package cmd

import (
	"fmt"
	"net"
	"strings"

	"airdb.io/airdb/adb/internal/adblib"
	"airdb.io/airdb/sailor"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/go-sql-driver/mysql"
	"github.com/miekg/dns"
	"github.com/spf13/cobra"
)

var mysqlCmd = &cobra.Command{
	Use:                "mysql [service]",
	Short:              "mysql client",
	Long:               "Airdb mysql client",
	DisableFlagParsing: false,
	Args:               cobra.MinimumNArgs(0),
	Example:            adblib.SQLDoc,
	Aliases:            []string{"sql"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			listDatabase()

			return
		}

		mysqlExec(args)
	},
}

var dsnAddCmd = &cobra.Command{
	Use:     "add [name] [dsn tet record value]",
	Short:   "Add new dsn",
	Long:    "Add new dsn",
	Example: adblib.DNSSrvDoc,
	Args:    cobra.MinimumNArgs(servicesAddCmdMinArgs),
	Run: func(cmd *cobra.Command, args []string) {
		addDsn(args)
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
	mysqlCmd.AddCommand(dsnAddCmd)

	mysqlCmd.PersistentFlags().StringVarP(&mysqlFlags.User, "user", "u", "root", "database username")
	mysqlCmd.PersistentFlags().StringVarP(&mysqlFlags.Password, "password", "p", "airdb.me", "database password")
	mysqlCmd.PersistentFlags().StringVarP(&mysqlFlags.DB, "db", "", "test", "database name")
}

func mysqlExec(args []string) {
	client, err := aliyunConfigInit()
	if err != nil {
		panic(err)
	}

	request := alidns.CreateDescribeDomainRecordsRequest()
	request.DomainName = ServiceDomain
	request.RRKeyWord = args[0]

	output, err := client.DescribeDomainRecords(request)
	if err != nil {
		fmt.Println(err)
	}

	// rrs := output.DomainRecords.Record
	dsn := ""

	for _, rr := range output.DomainRecords.Record {
		if rr.Type == dns.TypeToString[dns.TypeTXT] && rr.RR == request.RRKeyWord {
			// fmt.Printf("%-32s\t%s\n", rr.RR, rr.Value)
			dsn = rr.Value
		}
	}

	config, err := mysql.ParseDSN(dsn)
	if err != nil {
		fmt.Println(err)

		return
	}

	host, port, err := net.SplitHostPort(config.Addr)
	if err != nil {
		return
	}

	flags := fmt.Sprintf("-A -h%s -P%s -u%s -p%s %s",
		host,
		port,
		config.User,
		config.Passwd,
		config.DBName,
	)

	args = strings.Split(flags, " ")
	sailor.Exec("mysql", args)
}

func listDatabase() {
	client, err := aliyunConfigInit()
	if err != nil {
		panic(err)
	}

	request := alidns.CreateDescribeDomainRecordsRequest()
	request.DomainName = ServiceDomain

	output, err := client.DescribeDomainRecords(request)
	if err != nil {
		fmt.Println(err)
	}

	for _, rr := range output.DomainRecords.Record {
		if rr.RR == sailor.DelimiterStar || rr.RR == sailor.DelimiterAt {
			continue
		}

		// fmt.Printf("%-20s\t%-32s\t%s\n", rr.RecordId, rr.RR, rr.Value)
		fmt.Printf("%-20s %-5s %-32s %-64s %s\n", rr.RecordId, rr.Type, rr.RR, rr.Value, rr.Remark)
	}
}

func addDsn(args []string) {
	client, err := aliyunConfigInit()
	if err != nil {
		panic(err)
	}

	request := alidns.CreateAddDomainRecordRequest()
	request.DomainName = ServiceDomain
	request.Type = dns.TypeToString[dns.TypeTXT]
	request.RR = args[0]
	request.Value = args[1]

	output, err := client.AddDomainRecord(request)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(output)
}

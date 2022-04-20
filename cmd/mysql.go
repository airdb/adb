package cmd

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/airdb/sailor/osutil"
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
	Example:            SQLDoc,
	Aliases:            []string{"sql"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			listDatabase()

			return
		}

		mysqlExec(args)
	},
}

var mysqlExprCmd = &cobra.Command{
	Use:                "expr [dsn]",
	Short:              "mysql expr client",
	Long:               "mysql expr client",
	DisableFlagParsing: false,
	Args:               cobra.MinimumNArgs(0),
	Example:            SQLDoc,
	Aliases:            []string{"expr", "expression", "exp"},
	Run: func(cmd *cobra.Command, args []string) {
		GenMyqlExpr()
	},
}

var dsnAddCmd = &cobra.Command{
	Use:   "add [name] [dsn tet record value]",
	Short: "Add new dsn",
	Long:  "Add new dsn",
	Args:  cobra.MinimumNArgs(servicesAddCmdMinArgs),
	Run: func(cmd *cobra.Command, args []string) {
		addDsn(args)
	},
}

var dsnUpdateCmd = &cobra.Command{
	Use:   "update [name] [dsn tet record value]",
	Short: "Update new dsn",
	Long:  "Update new dsn",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		updateDsn()
	},
}

var dsnDeleteCmd = &cobra.Command{
	Use:   "delete [name] [dsn tet record value]",
	Short: "Delete new dsn",
	Long:  "Delete new dsn",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		deleteDsn()
	},
}

func mysqlCmdInit() {
	rootCmd.AddCommand(mysqlCmd)
	mysqlCmd.AddCommand(dsnAddCmd)
	mysqlCmd.AddCommand(dsnUpdateCmd)
	mysqlCmd.AddCommand(dsnDeleteCmd)
	mysqlCmd.AddCommand(mysqlExprCmd)

	dsnUpdateCmd.PersistentFlags().StringVarP(&updateDNSFlag.RecordID,
		"id", "i", "", "domain record_id")
	dsnUpdateCmd.PersistentFlags().StringVarP(&updateDNSFlag.RR,
		"rr", "r", "", "domain name prefix")
	dsnUpdateCmd.PersistentFlags().StringVarP(&updateDNSFlag.Value,
		"value", "v", "", "domain name prefix")

	dsnDeleteCmd.PersistentFlags().StringVarP(&updateDNSFlag.RecordID,
		"id", "i", "", "domain record_id")
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

	flags := fmt.Sprintf("-A --auto-rehash -h%s -P%s -u%s -p%s",
		host,
		port,
		config.User,
		config.Passwd,
	)

	args = strings.Split(flags, " ")
	args = append(args,
		"--prompt",
		fmt.Sprintf("mysql [%s]> ", config.DBName),
		config.DBName,
	)

	osutil.Exec("mysql", args)
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
		if rr.Type == dns.TypeToString[dns.TypeTXT] {
			// fmt.Printf("%-20s\t%-32s\t%s\n", rr.RecordId, rr.RR, rr.Value)
			fmt.Printf("%-20s %-5s %-32s %-64s %s\n", rr.RecordId, rr.Type, rr.RR, rr.Value, rr.Remark)
		}
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

func updateDsn() {
	client, err := aliyunConfigInit()
	if err != nil {
		panic(err)
	}

	request := alidns.CreateUpdateDomainRecordRequest()
	request.RecordId = updateDNSFlag.RecordID
	request.Type = dns.TypeToString[dns.TypeTXT]
	request.RR = updateDNSFlag.RR
	request.Value = updateDNSFlag.Value

	output, err := client.UpdateDomainRecord(request)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(output)
}

func deleteDsn() {
	client, err := aliyunConfigInit()
	if err != nil {
		panic(err)
	}

	request := alidns.CreateDeleteDomainRecordRequest()
	request.RecordId = updateDNSFlag.RecordID

	output, err := client.DeleteDomainRecord(request)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(output)
}

func GenMyqlExpr() {
	var dsn string

	var err error

	// check if there is something to read on STDIN
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// var stdin []byte
		reader := bufio.NewReader(os.Stdin)
		dsn, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Enter database dsn:")

		fmt.Scanf("%s", &dsn)
	}

	dsn = strings.TrimSpace(dsn)
	// fmt.Println("dsn = ", dsn)
	config, _ := mysql.ParseDSN(dsn)

	ip, port, _ := net.SplitHostPort(config.Addr)

	fmt.Println("mysql expression:")
	fmt.Printf("mysql -h%s -P%s -u%s -p%s --prompt \"mysql [%s]> \" %s\n",
		ip,
		port,
		config.User,
		config.Passwd,
		config.DBName,
		config.DBName,
	)

	fmt.Printf("mysqldump --single-transaction -h%s -P%s -u%s -p%s %s > %s_%s.dump.sql \n",
		ip,
		port,
		config.User,
		config.Passwd,
		config.DBName,
		config.DBName,
		time.Now().Format("2006_0102"),
	)
}

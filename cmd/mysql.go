package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/airdb/toolbox/osutil"
	"github.com/go-sql-driver/mysql"
	"github.com/miekg/dns"
	"github.com/spf13/cobra"
)

var mysqlCmd = &cobra.Command{
	Use:     "mysql [service]",
	Short:   "mysql client",
	Long:    "Airdb mysql client",
	Example: SQLDoc,
	Aliases: []string{"sql"},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return listDatabase()
		}

		return mysqlExec(args)
	},
}

var mysqlExprCmd = &cobra.Command{
	Use:     "expr [dsn]",
	Short:   "mysql expr client",
	Long:    "mysql expr client",
	Example: SQLDoc,
	Aliases: []string{"expression", "exp"},
	RunE: func(cmd *cobra.Command, args []string) error {
		return GenMysqlExpr()
	},
}

var dsnAddCmd = &cobra.Command{
	Use:   "add [name] [dsn txt record value]",
	Short: "Add new dsn",
	Long:  "Add new dsn",
	Args:  cobra.MinimumNArgs(servicesAddCmdMinArgs),
	RunE: func(cmd *cobra.Command, args []string) error {
		return addRecord(ServiceDomain, dns.TypeToString[dns.TypeTXT], args[0], args[1])
	},
}

var dsnUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update dsn",
	Long:  "Update dsn",
	RunE: func(cmd *cobra.Command, args []string) error {
		if updateDNSFlag.RecordID == "" || updateDNSFlag.RR == "" || updateDNSFlag.Value == "" {
			return errors.New("flags --id, --rr and --value are required")
		}

		return updateRecord(updateDNSFlag.RecordID,
			dns.TypeToString[dns.TypeTXT], updateDNSFlag.RR, updateDNSFlag.Value)
	},
}

var dsnDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete dsn",
	Long:  "Delete dsn",
	RunE: func(cmd *cobra.Command, args []string) error {
		if updateDNSFlag.RecordID == "" {
			return errors.New("flag --id is required")
		}

		return deleteRecord(updateDNSFlag.RecordID)
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
		"value", "v", "", "domain record value")

	dsnDeleteCmd.PersistentFlags().StringVarP(&updateDNSFlag.RecordID,
		"id", "i", "", "domain record_id")
}

func lookupDsn(service string) (string, error) {
	records, err := describeRecords(ServiceDomain)
	if err != nil {
		return "", err
	}

	for _, rr := range records {
		if rr.Type == dns.TypeToString[dns.TypeTXT] && rr.RR == service {
			return rr.Value, nil
		}
	}

	return "", fmt.Errorf("no dsn record found for service %q under %s", service, ServiceDomain)
}

func mysqlExec(args []string) error {
	dsn, err := lookupDsn(args[0])
	if err != nil {
		return err
	}

	config, err := mysql.ParseDSN(dsn)
	if err != nil {
		return fmt.Errorf("parse dsn: %w", err)
	}

	host, port, err := net.SplitHostPort(config.Addr)
	if err != nil {
		return fmt.Errorf("parse dsn address %q: %w", config.Addr, err)
	}

	// Pass the password via MYSQL_PWD so it does not show up in `ps`.
	if err := os.Setenv("MYSQL_PWD", config.Passwd); err != nil {
		return err
	}

	mysqlArgs := []string{
		"-A", "--auto-rehash",
		"-h" + host,
		"-P" + port,
		"-u" + config.User,
		"--prompt", fmt.Sprintf("mysql [%s]> ", config.DBName),
		config.DBName,
	}

	osutil.Exec("mysql", mysqlArgs)

	return nil
}

func listDatabase() error {
	records, err := describeRecords(ServiceDomain)
	if err != nil {
		return err
	}

	for _, rr := range records {
		if rr.Type == dns.TypeToString[dns.TypeTXT] {
			fmt.Printf("%-20s %-5s %-32s %-64s %s\n", rr.RecordId, rr.Type, rr.RR, rr.Value, rr.Remark)
		}
	}

	return nil
}

func GenMysqlExpr() error {
	var dsn string

	// check if there is something to read on STDIN
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		line, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil && line == "" {
			return err
		}

		dsn = line
	} else {
		fmt.Println("Enter database dsn:")
		fmt.Scanf("%s", &dsn)
	}

	dsn = strings.TrimSpace(dsn)

	config, err := mysql.ParseDSN(dsn)
	if err != nil {
		return fmt.Errorf("parse dsn: %w", err)
	}

	ip, port, err := net.SplitHostPort(config.Addr)
	if err != nil {
		return fmt.Errorf("parse dsn address %q: %w", config.Addr, err)
	}

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

	return nil
}

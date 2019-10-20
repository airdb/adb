package command

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"

	"github.com/spf13/viper"
)

type DatabaseItem struct {
	User             string
	Password         string
	Address          string
	Name             string
	DefaultTableName bool `mapstructure:"default_table_name"`
}

func mysql(args cli.Args) error {
	viper.SetConfigFile("conf/dev.json")
	if args.First() != "" {
		viper.SetConfigFile(args.First())
	}

	databases := make(map[string]*DatabaseItem)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
		return err
	}

	if err := viper.UnmarshalKey("databases", &databases); err != nil {
		log.Fatal(err)
		return err
	}

	for name := range databases {
		if databases[name].DefaultTableName {
			host, port, _ := net.SplitHostPort(databases[name].Address)
			fmt.Println(databases[name].Password)
			fmt.Printf("mysql -h%s -P%s -u%s -p %s\n", host, port, databases[name].User, name)
		}
	}
	return nil
}

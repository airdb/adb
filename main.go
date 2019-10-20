package main

import (
	"log"
	"os"
	//	"sort"

	"github.com/airdb/adb/command"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "adb"
	app.Usage = "Airdb Development Builder Command Line Interface"
	app.Description = "Airdb Development Builder help you initialise server or development environment.\n" +
		"\t Release your project and deploy your projects."
	app.Version = "1.0.0"
	app.Authors = []cli.Author{
		{
			Name:  "Dean CN",
			Email: "dean@airdb.com",
		},
	}
	app.Copyright = "https://www.airdb.com"

	//   app.EnableBashCompletion = true

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from `FILE`",
		},
	}

	app.Commands = command.InitCommand()

	//	sort.Sort(cli.FlagsByName(app.Flags))
	//	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

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
		"\t Release your project and deploy your project."
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

	/*
		app.Commands = []cli.Command{
			{
				Name:  "init",
				Usage: "init server or tool",
				Action: func(c *cli.Context) error {
					log.Println("aa")
					return nil
				},
			},
			{
				Name:    "release",
				Aliases: []string{"r"},
				Usage:   "release a git branch with",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:    "deploy",
				Aliases: []string{"d"},
				Usage:   "deoply project to cloud server or docker container",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
		}
	*/

	//	sort.Sort(cli.FlagsByName(app.Flags))
	//	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

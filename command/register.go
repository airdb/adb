package command

import (
	"log"

	"github.com/urfave/cli"
)

func InitCommand() []cli.Command {
	return []cli.Command{
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
}

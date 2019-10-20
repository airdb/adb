package command

import (
	"fmt"
	"github.com/urfave/cli"
)

func InitCommand() []cli.Command {
	return []cli.Command{
		{
			Name:  "env",
			Usage: "check local env",
			Action: func(c *cli.Context) error {
				return env()
			},
		},
		{
			Name:  "init",
			Usage: "init server or tool",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:    "release",
			Aliases: []string{"r"},
			Usage:   "release a git branch with",
			Action: func(c *cli.Context) error {
				return release()
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
		{
			Name:    "update",
			Aliases: []string{"u"},
			Usage:   "update adb tool",
			Action: func(c *cli.Context) error {
				return update()
			},
		},
		{
			Name:  "bbhj",
			Usage: "query or set bbhj information",
			Action: func(c *cli.Context) error {
				return bbhj()
			},
			BashComplete: func(c *cli.Context) {
				// This will complete if no args are passed
				if c.NArg() > 0 {
					return
				}

				tasks := []string{"cook", "clean", "laundry", "eat", "sleep", "code"}
				for _, t := range tasks {
					fmt.Println(t)
				}
			},
		},
		{
			Name:  "host",
			Usage: "host operation",
			Action: func(c *cli.Context) error {
				return host()
			},
		},
		{
			Name:  "mysql",
			Usage: "mysql client",
			Action: func(c *cli.Context) error {
				return mysql(c.Args())
			},
		},
	}
}

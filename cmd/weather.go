package cmd

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var weatherCommand = &cobra.Command{
	Use:     "wttr",
	Aliases: []string{"hello"},
	Short:   "show weather",
	Long:    "The right way to check the weather",
	Run: func(cmd *cobra.Command, args []string) {
		weather(args)
	},
}

const weatherAPI = "https://wttr.in/"

func weather(args []string) {
	apiurl := weatherAPI
	if len(args) != 0 {
		apiurl = weatherAPI + args[0]
	}

	fmt.Println(apiurl)
	cmd := exec.Command("/usr/bin/curl", apiurl)

	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		fmt.Println("Thanks for using adb tool!")
	} else {
		fmt.Println(strings.Trim(out.String(), "\n"))
	}
}

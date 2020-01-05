package cmd

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var wikiCommand = &cobra.Command{
	Use:   "wiki",
	Short: "airdb wiki",
	Long:  "airdb wiki, https://airdb.wiki",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			Usage()
			return
		}
		wiki(args[0])
	},
}

func Usage() {
	fmt.Println("Usage:")
	fmt.Println("  adb wiki [project_name]")
	fmt.Println()
	fmt.Println("Projects:")
	fmt.Println("\tdev")
	fmt.Println("\twork")
	fmt.Println("\tgo")
	fmt.Println("\tinterview")
	fmt.Println("\tai")
	fmt.Println("\tkube")
	fmt.Println("\tbbhj")
	fmt.Println("\tlinux")
	fmt.Println()
	fmt.Println("Airdb Wiki: https://airdb.wiki")
}

func wiki(wikiName string) {
	wikiArgs := getWikiArgs(wikiName)
	cmd := exec.Command("open", wikiArgs...)

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Thanks for using adb tool!", err)
	} else {
		fmt.Println(strings.Trim(out.String(), "\n"))
	}
}

func getWikiArgs(name string) []string {
	args := []string{
		"-a",
		"Google Chrome",
	}
	if name == "" {
		args = append(args, "https://airdb.wiki/")
	} else {
		args = append(args, "https://airdb.wiki/"+name)
	}
	return args
}

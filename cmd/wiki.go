package cmd

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

const AirdbWiki = "https://airdb-wiki.github.io/"

var wikiCommand = &cobra.Command{
	Use:     "wiki",
	Short:   "airdb wiki",
	Long:    "airdb wiki, https://airdb-wiki.github.io",
	Example: "adb wiki [project_name]",
}

func Usage() {
	fmt.Println("Usage:")
	fmt.Println("  adb wiki [project_name]")
	fmt.Println()
	fmt.Println()
	fmt.Printf("Airdb Wiki: %s.\n", AirdbWiki)
}

type Repo struct {
	ID     uint   `json:"id"`
	NodeID string `json:"nodeId"`
	Name   string `json:"name"`
}

func wiki(wikiName string) {
	wikiArgs := getWikiArgs(wikiName)
	cmd := exec.Command("open", wikiArgs...)

	var out bytes.Buffer

	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
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
		args = append(args, AirdbWiki)
	} else {
		args = append(args, AirdbWiki+name)
	}

	return args
}

var interviewWikiCommand = &cobra.Command{
	Use:   "interview",
	Short: "interview wiki",
	Long:  "interview wiki",
	Run: func(cmd *cobra.Command, args []string) {
		name := "interview"
		wiki(name)
	},
}

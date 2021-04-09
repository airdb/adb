package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"sort"
	"strings"

	"github.com/imroc/req"
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
	listRepos()
	fmt.Println()
	fmt.Printf("Airdb Wiki: %s.\n", AirdbWiki)
}

type Repo struct {
	ID     uint   `json:"id"`
	NodeID string `json:"node_id"`
	Name   string `json:"name"`
}

func listRepos() {
	apiurl := "https://api.github.com/orgs/airdb-wiki/repos"

	resp, err := req.Get(apiurl)
	if err != nil {
		log.Println("Query Github failed. https://github.com/airdb/airdb-wiki")

		return
	}

	repos := make([]Repo, 0)

	err = json.Unmarshal(resp.Bytes(), &repos)
	if err != nil {
		log.Println("json unmarshall failed.")

		return
	}

	sort.Slice(repos, func(i, j int) bool { return len(repos[i].Name) < len(repos[j].Name) })
	fmt.Println("Projects:")

	for _, repo := range repos {
		switch repo.Name {
		case strings.TrimPrefix(AirdbWiki, "https"):
		default:
			fmt.Printf("\t%s\t%30s\n", repo.Name, AirdbWiki+repo.Name)
		}
	}
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

var listWikiCommand = &cobra.Command{
	Use:   "list",
	Short: "list wiki",
	Long:  "list wiki",
	Run: func(cmd *cobra.Command, args []string) {
		listRepos()
	},
}

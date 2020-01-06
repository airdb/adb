package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/imroc/req"
	"log"
	"os/exec"
	"sort"
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
	listRepos()
	fmt.Println()
	fmt.Println("Airdb Wiki: https://airdb.wiki.")
}

type Repo struct {
	ID uint `json:"id"`
	NodeID string `json:"node_id"`
	Name string `json:"name"`
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
		case "airdb-wiki.github.io":
		default:
			fmt.Printf("\t%s%30s\n", repo.Name, "thttps://airdb.wiki/" + repo.Name)
		}
	}
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

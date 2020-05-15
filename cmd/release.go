package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/imroc/req"
	"github.com/spf13/cobra"
)

var releaseCommand = &cobra.Command{
	Use:                "release",
	Short:              "release operation",
	Long:               "release operation",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		release()
	},
}

type Tag struct {
	NodeID string `json:"node_id"`
	Object struct {
		Sha  string `json:"sha"`
		Type string `json:"type"`
		URL  string `json:"url"`
	} `json:"object"`
	Ref string `json:"ref"`
	URL string `json:"url"`
}

func release() {
	getTags()
}

func getTags() {
	repo := getRepoName()
	apiurl := "https://api.github.com/repos/" + repo + "/git/refs/tags"

	r, err := req.Get(apiurl)
	if err != nil {
		return
	}

	tagsList := make([]Tag, 0)
	err = json.Unmarshal(r.Bytes(), &tagsList)

	if err != nil {
		return
	}

	fmt.Printf("%-16s\t%s\n", "Repo", "Version")

	if len(tagsList) == 0 {
		fmt.Printf("%-16s\t%s\n", repo, "no tags")
		return
	}

	for _, tag := range tagsList {
		version := strings.TrimPrefix(tag.Ref, "refs/tags/")
		fmt.Printf("%-16s\t%s\n", repo, version)
	}
}

func getRepoName() string {
	cmd := exec.Command("git", "config", "remote.origin.url")

	var out bytes.Buffer

	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	repo := strings.TrimPrefix(out.String(), "https://github.com/")

	return strings.TrimRight(repo, "\n")
}

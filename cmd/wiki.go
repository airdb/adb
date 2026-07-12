package cmd

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

const AirdbWiki = "https://airdb-wiki.github.io/"

var wikiCommand = &cobra.Command{
	Use:     "wiki [project_name]",
	Short:   "airdb wiki",
	Long:    "airdb wiki, https://airdb-wiki.github.io",
	Example: "adb wiki [project_name]",
	RunE: func(cmd *cobra.Command, args []string) error {
		name := ""
		if len(args) > 0 {
			name = args[0]
		}

		return wiki(name)
	},
}

func wiki(wikiName string) error {
	url := AirdbWiki + wikiName

	if err := openBrowser(url); err != nil {
		return fmt.Errorf("open %s: %w", url, err)
	}

	fmt.Println("Opened", url)

	return nil
}

func openBrowser(url string) error {
	switch runtime.GOOS {
	case "darwin":
		return exec.Command("open", url).Run()
	case "windows":
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Run()
	default:
		return exec.Command("xdg-open", url).Run()
	}
}

var interviewWikiCommand = &cobra.Command{
	Use:   "interview",
	Short: "interview wiki",
	Long:  "interview wiki",
	RunE: func(cmd *cobra.Command, args []string) error {
		return wiki("interview")
	},
}

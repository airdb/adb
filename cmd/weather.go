package cmd

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var weatherCommand = &cobra.Command{
	Use:     "wttr",
	Aliases: []string{"hello"},
	Short:   "show weather",
	Long:    "The right way to check the weather",
	RunE: func(cmd *cobra.Command, args []string) error {
		return weather(args)
	},
}

const weatherAPI = "https://wttr.in/"

func weather(args []string) error {
	apiurl := weatherAPI
	if len(args) != 0 {
		apiurl += args[0]
	}

	fmt.Println(apiurl)

	client := &http.Client{Timeout: 30 * time.Second}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, apiurl, nil)
	if err != nil {
		return err
	}

	// wttr.in returns plain text for curl-like user agents.
	req.Header.Set("User-Agent", "curl/8.0")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))

	return nil
}

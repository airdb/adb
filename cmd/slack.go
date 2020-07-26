/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/airdb/adb/internal/adblib"
	"strings"

	"github.com/slack-go/slack"
	"github.com/spf13/cobra"
)

// slackCmd represents the slack command
var slackCmd = &cobra.Command{
	Use:   "slack",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		opSlack(args)
	},
}

func initSlack() {
	rootCmd.AddCommand(slackCmd)
}

func opSlack(args []string) {
	config := adblib.GetSlackConfig()
	api := slack.New(config.Token)

	channelName := config.Channel
	msg := slack.MsgOptionText(strings.Join(args, " "), false)

	channelID, timestamp, err := api.PostMessage(channelName, msg)
	if err != nil {
		fmt.Printf("%s\n", err)
		fmt.Printf("Send message to %s failed, id: %s timestamp: %s\n", channelName, channelID, timestamp)
		return
	}

	fmt.Printf("Send message to %s successfully\n", channelName)
}

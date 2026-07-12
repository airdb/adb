package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/airdb/adb/internal/adblib"
	"github.com/slack-go/slack"
	"github.com/spf13/cobra"
)

// slackCmd represents the slack command.
var slackCmd = &cobra.Command{
	Use:   "slack [message]",
	Short: "Send a message to the configured Slack channel",
	Long:  "Send a message to the Slack channel configured via SlackToken/SlackChannel in ~/" + adblib.EnvFile,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return opSlack(args)
	},
}

func initSlack() {
	rootCmd.AddCommand(slackCmd)
}

func opSlack(args []string) error {
	if adblib.AdbConfig.SlackToken == "" || adblib.AdbConfig.SlackChannel == "" {
		return errors.New("SlackToken and SlackChannel must be set in ~/" + adblib.EnvFile)
	}

	api := slack.New(adblib.AdbConfig.SlackToken)

	channelName := adblib.AdbConfig.SlackChannel
	msg := slack.MsgOptionText(strings.Join(args, " "), false)

	if _, _, err := api.PostMessage(channelName, msg); err != nil {
		return fmt.Errorf("send message to %s: %w", channelName, err)
	}

	fmt.Printf("Send message to %s successfully\n", channelName)

	return nil
}

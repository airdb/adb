package cmd

import (
	"fmt"

	"airdb.io/airdb/adb/internal/adblib"
	"github.com/spf13/cobra"
)

var bbhjCommand = &cobra.Command{
	Use:                "bbhj",
	Short:              "bbhj operation",
	Long:               "bbhj operation",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		bbhj()
	},
}

func bbhj() {
	client, err := adblib.NewCLBClient()
	if err != nil {
		fmt.Println(err)

		return
	}

	client.ListCLB()
}

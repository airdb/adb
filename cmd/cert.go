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

	"github.com/spf13/cobra"
)

// slackCmd represents the slack command.
var certCmd = &cobra.Command{
	Use:   "cert",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		opCert(args)
	},
}

type CertFlag struct {
	PrivateKeyFile string
	PublicKeyFile  string
}

var certFlag CertFlag

func initCert() {
	rootCmd.AddCommand(certCmd)

	certCmd.PersistentFlags().StringVarP(&certFlag.PrivateKeyFile, "private key", "-k", "/tmp/",
		"server.key")
	certCmd.PersistentFlags().StringVarP(&certFlag.PublicKeyFile, "public key", "-c", "/tmp/",
		"ssl.chain.crt")
}

func opCert(args []string) {
	fmt.Println(args)
}

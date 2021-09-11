package cmd

import (
	"fmt"

	"github.com/airdb/adb/internal/adblib"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var envCommand = &cobra.Command{
	Use:                "env",
	Short:              "show env",
	Long:               "Show Environment",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		env()
	},
}

func env() {
	envs, err := godotenv.Read(adblib.GetEnvFile())
	if err != nil {
		fmt.Println("read envfile err ", err)
	}

	for k, v := range envs {
		fmt.Printf("export %s=%s\n", k, v)
	}
}

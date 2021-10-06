package cmd

import (
	"fmt"
	"os"

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
	envfile := ".env"

	if _, err := os.Stat(envfile); os.IsNotExist(err) {
		envfile = adblib.GetEnvFile()
	}

	envs, err := godotenv.Read(envfile)
	if err != nil {
		fmt.Println("read envfile err ", err)
	}

	for k, v := range envs {
		// fmt.Printf("export %s=%s\n", k, v)
		fmt.Printf("export %s=\"%s\"\n", k, v)
	}
}

package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var loginCommand = &cobra.Command{
	Use:   "login",
	Short: "login airdb",
	Long:  "login airdb",
	Run: func(cmd *cobra.Command, args []string) {
		login()
	},
}

func login() {
	icon := os.Getenv("HOME") + "/.adb/icon"
	cmd := exec.Command("cat", icon)

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Thanks for using adb tool!")
	} else {
		fmt.Println(strings.Trim(out.String(), "\n"))
	}
}

package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/imroc/req"
	"github.com/spf13/cobra"
)

var loginCommand = &cobra.Command{
	Use:     "login",
	Aliases: []string{"hello"},
	Short:   "login airdb",
	Long:    "login airdb",
	Run: func(cmd *cobra.Command, args []string) {
		login()
	},
}

func login() {
	iconFileName := path.Join(os.Getenv("HOME"), ".adb/icon")
	mtod := "https://init.airdb.host/mtod/icon"
	r, err := req.Get(mtod)
	if err == nil {
		msg, _ := r.ToString()
		fmt.Print(msg)
		if err = r.ToFile(iconFileName); err == nil {
			return
		}
	}

	cmd := exec.Command("cat", iconFileName)

	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		fmt.Println("Thanks for using adb tool!")
	} else {
		fmt.Println(strings.Trim(out.String(), "\n"))
	}
}

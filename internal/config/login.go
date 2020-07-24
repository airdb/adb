package config

import (
	"bytes"
	"fmt"
	"github.com/imroc/req"
	"os/exec"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

// The  flag will be written to this struct.
type User struct {
	Name string `json:"name" survey:"name" mapstructure:"name"`
	Password string `json:"password" survey:"Password" mapstructure:"password"`
}

// The questions to ask.
var userLogin = []*survey.Question{
	{
		Name: "name",
		Prompt:    &survey.Input{Message: "username:"},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
	{
		Name:      "password",
		Prompt:    &survey.Password{ Message: "Password:" },
		Validate:  survey.Required,
		Transform: survey.Title,
	},
}

func Login() {
	var user User

	// pPerform the questions.
	err := survey.Ask(userLogin, &user)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	cmd := exec.Command("cat", IconFile())

	var out bytes.Buffer

	cmd.Stdout = &out

	err = cmd.Run()
	if err != nil {
		downloadIcon()
		fmt.Println("Thanks for using adb tool!")
	} else {
		fmt.Println(strings.Trim(out.String(), "\n"))
	}

	return
}

func downloadIcon() {
	mtod := "https://init.airdb.host/mtod/icon"
	r, err := req.Get(mtod)

	if err == nil {
		msg, _ := r.ToString()
		fmt.Print(msg)

		if err = r.ToFile(IconFile()); err == nil {
			return
		}
	}
}

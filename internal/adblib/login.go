package adblib

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/airdb/sailor"
	"github.com/imroc/req"
)

// The  flag will be written to this struct.
type User struct {
	Name     string `json:"name" survey:"name" mapstructure:"name"`
	Password string `json:"password" survey:"Password" mapstructure:"password"`
}

// The questions to ask.
var userLogin = []*survey.Question{
	{
		Name:     "name",
		Prompt:   &survey.Input{Message: "Username:"},
		Validate: survey.Required,
	},
	{
		Name:     "password",
		Prompt:   &survey.Password{Message: "Password:"},
		Validate: survey.Required,
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

	args := []string{IconFile()}

	out, err := sailor.ExecCommand("cat", args)
	if err != nil {
		downloadIcon()
		return
	}

	fmt.Println(out)
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

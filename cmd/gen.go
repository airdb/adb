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
	"bytes"
	"go/build"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"text/template"

	_ "github.com/airdb/adb/statik" // statik zip data, `statik -include='*' -src gin-template/ -f`
	"github.com/rakyll/statik/fs"
	"github.com/spf13/cobra"
)

// GenCmd represents the gen command.
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate a new airdb project",
	Long:  `Generate a new airdb project`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			println("usage: adb gen [module]")
			return
		}

		generate(args)
	},
}

func genCmdInit() {
	rootCmd.AddCommand(genCmd)

	genCmd.PersistentFlags().StringVarP(&genFlags.Module, "module", "m", "github.com/airdb/demo", "project module name")
}

type genStruct struct {
	Module string
}

var genFlags = genStruct{}

func generate(args []string) {
	module := args[0]

	projectDir := path.Join(build.Default.GOPATH, "src", module)

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal("new statik fs failed, err: ", err)
		return
	}

	err = fs.Walk(statikFS, "/", func(originPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		projectDir := path.Join(projectDir, originPath)

		srcFile, err := statikFS.Open(originPath)
		if err != nil {
			return err
		}

		err = GenerateFileFromReader(srcFile, projectDir, nil)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return
	}

	log.Printf("Generate project successfully, project dir: %s",
		projectDir,
	)
}

func GenerateString(str string, data interface{}) (string, error) {
	tmpl, err := template.New("").Parse(str)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func GenerateFileFromReader(reader io.Reader, dstPath string, data interface{}) error {
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	content, err := GenerateString(string(b), data)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(dstPath, []byte(content), 0600)
	if err != nil {
		log.Println("write file failed, err: ", err)
		return err
	}

	return err
}

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
	"go/build"
	"log"
	"os"
	"path"

	_ "github.com/airdb/adb/statik" // statik zip data, `statik -include='*' -src gin-template/ -f`
	"github.com/airdb/sailor"
	"github.com/rakyll/statik/fs"
	"github.com/spf13/cobra"
)

// GenCmd represents the gen command.
var genCmd = &cobra.Command{
	Use:     "gen",
	Short:   "Generate a new airdb project",
	Long:    `Generate a new airdb project`,
	Aliases: []string{"generate", "gene"},
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

		err = sailor.TemplateGenerateFileFromReader(srcFile, projectDir, nil)
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

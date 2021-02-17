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
	"go/build"
	"log"
	"os"
	"path"

	// statik zip data, `statik -include='*' -src gin-template/ -f`.
	_ "airdb.io/airdb/adb/statik"
	"airdb.io/airdb/sailor/fileutil"
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
		generate(args)
	},
}

func genCmdInit() {
	rootCmd.AddCommand(genCmd)

	genCmd.PersistentFlags().StringVarP(&generateFlags.GitURL, "git", "g", "github.com", "git domain")
	genCmd.PersistentFlags().StringVarP(&generateFlags.Owner, "owner", "o", "airdb", "owner or organization")
	genCmd.PersistentFlags().StringVarP(&generateFlags.Repo, "repo", "r", "demo", "project repo")
}

type generateStruct struct {
	// github.com
	GitURL string

	// airdb
	Owner string

	// demo
	Repo string

	// github.com/airdb/demo
	GoModulePath string
}

var generateFlags = generateStruct{}

func generate(args []string) {
	if len(args) > 0 {
		generateFlags.Repo = args[0]
	}

	generateFlags.GoModulePath = fmt.Sprintf("d%s/%s/%s",
		generateFlags.GitURL,
		generateFlags.Owner,
		generateFlags.Repo,
	)

	projectDir := path.Join(build.Default.GOPATH, "src", generateFlags.GoModulePath)

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

		err = fileutil.TemplateGenerateFileFromReader(srcFile, projectDir, generateFlags)
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

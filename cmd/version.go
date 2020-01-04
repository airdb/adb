package cmd

import (
	"encoding/json"
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

// Build version info.
type BuildInfo struct {
	BuildTime string
	GoVersion string
	Version   string
	Commit    string
}

var (
	BuildTime string
	Commit    string
	Version   string
)

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Version information",
	Long:  "Version information",
	Run: func(cmd *cobra.Command, args []string) {
		info := BuildInfo{
			BuildTime: BuildTime,
			GoVersion: runtime.Version(),
			Version:   Version,
			Commit:    Commit,
		}

		out, err := json.Marshal(info)
		if err != nil {
			panic(err)
		}

		fmt.Printf("version.BuildInfo%s\n", string(out))
	},
}

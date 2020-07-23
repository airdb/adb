package cmd

import (
	"encoding/json"
	"runtime"
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

func getVersion() string {
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

	// fmt.Printf("version.BuildInfo%s\n", string(out))

	return string(out)
}

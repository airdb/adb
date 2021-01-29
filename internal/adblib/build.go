package adblib

import (
	"encoding/json"
	"runtime"
)

// Build version info.
type BuildInfo struct {
	GoVersion string
	Version   string
	Build     string
	BuildTime string
}

var (
	Version   string
	Build     string
	BuildTime string
)

func GetVersion() string {
	info := BuildInfo{
		GoVersion: runtime.Version(),
		Version:   Version,
		Build:     Build,
		BuildTime: BuildTime,
	}

	out, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}

	// fmt.Printf("version.BuildInfo%s\n", string(out))

	return string(out) + "\n" + "https://github.com/airdb/adb/releases/latest"
}

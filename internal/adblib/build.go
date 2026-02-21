package adblib

import (
	"encoding/json"
	"runtime"
)

// Build version info.
type BuildInfo struct {
	GoVersion string
	GOOS      string
	GOARCH    string
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
		GOOS:      runtime.GOOS,
		GOARCH:    runtime.GOARCH,
		Version:   Version,
		Build:     Build,
		BuildTime: BuildTime,
	}

	out, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}

	// fmt.Printf("version.BuildInfo%s\n", string(out))

	ret := string(out)
	ret += "\n" + "https://github.com/airdb/adb/releases/latest"
	ret += "\n"
	ret += "\n" + "go install -ldflags -X=github.com/airdb/adb/internal/adblib.BuildTime= github.com/airdb/adb@dev"
	return ret
}

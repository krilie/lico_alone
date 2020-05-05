package run_env

import "os"

var (
	VERSION    string
	BUILD_TIME string
	GO_VERSION string
	GIT_COMMIT string
)

// RunEnv is the app's run env
type RunEnv struct {
	AppName   string
	AppHost   string
	Version   string
	BuildTime string
	GoVersion string
	GitCommit string
}

var RunEnvLocal = &RunEnv{
	AppName:   "lico_alone",
	AppHost:   GetHostName(),
	Version:   VERSION,
	BuildTime: BUILD_TIME,
	GoVersion: GO_VERSION,
	GitCommit: GIT_COMMIT,
}

func GetHostName() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "host-name"
	}
	return hostname
}

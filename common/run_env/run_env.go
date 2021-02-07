package run_env

import (
	"github.com/krilie/lico_alone/common/utils/id_util"
	"os"
)

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
	AppName:   "myapp",
	AppHost:   GetHostName(),
	Version:   VERSION,
	BuildTime: BUILD_TIME,
	GoVersion: GO_VERSION,
	GitCommit: GIT_COMMIT,
}

var tempHostName = id_util.NextSnowflake()

func GetHostName() string {
	hostname, err := os.Hostname()
	if err != nil {
		return tempHostName
	}
	return hostname
}

func (env RunEnv) GetShortGitCommitSha() string {
	if len(env.GitCommit) >= 6 {
		return env.GitCommit[0:6]
	}
	return env.GitCommit
}

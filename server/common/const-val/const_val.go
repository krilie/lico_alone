package const_val

import (
	"github.com/krilie/lico_alone/common/com-model/run-env"
)

// RunEnv 在程序开始运行时设置的全局变量
var RunEnv = run_env.RunEnv{
	AppName:   "lico_alone",
	AppHost:   "local",
	ClientId:  "",
	UserId:    "",
	Version:   "",
	BuildTime: "",
	GoVersion: "",
	GitCommit: "",
}

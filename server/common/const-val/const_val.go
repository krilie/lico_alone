package const_val

import (
	context_enum "github.com/krilie/lico_alone/common/model/context-enum"
)

// RunEnv 在程序开始运行时设置的全局变量
var RunEnv = context_enum.RunEnv{
	AppName:   "lico_alone",
	AppHost:   "local",
	ClientId:  "",
	UserId:    "",
	Version:   "",
	BuildTime: "",
	GoVersion: "",
	GitCommit: "",
}

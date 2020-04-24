package dig

import (
	"github.com/krilie/lico_alone/common/config"
	context_enum "github.com/krilie/lico_alone/common/model/context-enum"
	"go.uber.org/dig"
)

var Container = dig.New()

func MustProvide(constructor interface{}, opts ...dig.ProvideOption) {
	CheckErr(Container.Provide(constructor, opts...))
}
func MustInvoke(function interface{}, opts ...dig.InvokeOption) {
	CheckErr(Container.Invoke(function, opts...))
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	// 配置文件
	err := Container.Provide(func() *config.Config {
		return &config.Cfg
	})
	CheckErr(err)
	// runEnv
	err = Container.Provide(func() *context_enum.RunEnv {
		return &context_enum.RunEnv{
			AppName:    "AppName",
			AppVersion: "AppVersion",
			AppHost:    "AppHost",
			ClientId:   "ClientId",
			UserId:     "UserId",
			Version:    "Version",
			BuildTime:  "BuildTime",
			GoVersion:  "GoVersion",
			GitCommit:  "GitCommit",
		}
	})
	CheckErr(err)
}

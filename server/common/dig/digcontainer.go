package dig

import (
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/common/run_env"
	"go.uber.org/dig"
)

type StructContainer struct {
	*dig.Container
}

var Container = &StructContainer{
	Container: dig.New(),
}

func (c *StructContainer) MustProvide(constructor interface{}, opts ...dig.ProvideOption) {
	CheckErr(c.Container.Provide(constructor, opts...))
}
func (c *StructContainer) MustInvoke(function interface{}, opts ...dig.InvokeOption) {
	CheckErr(c.Container.Invoke(function, opts...))
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
	err = Container.Provide(func() *run_env.RunEnv {
		return run_env.RunEnvLocal
	})
	CheckErr(err)
}

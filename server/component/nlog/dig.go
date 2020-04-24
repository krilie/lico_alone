package nlog

import (
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/common/dig"
	context_enum "github.com/krilie/lico_alone/common/model/context-enum"
)

func init() {
	err := dig.Container.Provide(func(runEnv *context_enum.RunEnv, cfg *config.Config) *NLog {
		return NewLogger(*runEnv, *cfg)
	})
	dig.CheckErr(err)
}

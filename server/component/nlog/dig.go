package nlog

import (
	context_enum "github.com/krilie/lico_alone/common/common-model/context-enum"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/common/dig"
)

func init() {
	dig.MustProvide(func(runEnv *context_enum.RunEnv, cfg *config.Config) *NLog {
		return NewLogger(*runEnv, *cfg)
	})
}

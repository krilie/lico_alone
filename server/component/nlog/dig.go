package nlog

import (
	context_enum "github.com/krilie/lico_alone/common/com-model/context-enum"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/common/dig"
)

var Log *NLog

func init() {
	dig.Container.MustProvide(func(runEnv *context_enum.RunEnv, cfg *config.Config) *NLog {
		return NewLogger(*runEnv, *cfg)
	})
	dig.Container.MustInvoke(func(log *NLog) {
		Log = log
	})
}

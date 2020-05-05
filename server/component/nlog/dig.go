package nlog

import (
	"github.com/krilie/lico_alone/common/com-model/run-env"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/common/dig"
)

var Log *NLog

func init() {
	dig.Container.MustProvide(func(runEnv *run_env.RunEnv, cfg *config.Config) *NLog {
		return NewLogger(*runEnv, *cfg)
	})
	dig.Container.MustInvoke(func(log *NLog) {
		Log = log
	})
}

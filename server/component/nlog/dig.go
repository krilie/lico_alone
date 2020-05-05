package nlog

import (
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/run-env"
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

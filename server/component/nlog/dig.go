package nlog

import (
	"github.com/krilie/lico_alone/common/dig"
)

var Log *NLog

func init() {
	dig.Container.MustProvide(NewLogger)
	dig.Container.MustInvoke(func(log *NLog) {
		Log = log
	})
}

// DigProvider provider
func DigProvider() {
	dig.Container.MustProvide(NewLogger)
	dig.Container.MustInvoke(func(log *NLog) {
		Log = log
	})
}

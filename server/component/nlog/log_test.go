package nlog

import (
	"github.com/krilie/lico_alone/common/dig"
	"testing"
)

func TestLog(t *testing.T) {
	dig.MustInvoke(func(log *NLog) {
		log.Error("hello dig here")
		log.Info("hello dig info")
	})
}

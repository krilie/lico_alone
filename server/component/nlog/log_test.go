package nlog

import (
	"github.com/krilie/lico_alone/common/dig"
	"testing"
)

func TestLog(t *testing.T) {
	dig.Container.Invoke(func(log *NLog) {
		log.Error("hello dig here")
	})
}

// +build !auto_test

package nlog

import (
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/component/ncfg"
	"testing"
)

func TestMain(m *testing.M) {
	ncfg.DigProviderByCfgStrFromEnv()
	DigProvider()
	m.Run()
}

func TestLog(t *testing.T) {
	dig.Container.MustInvoke(func(log *NLog) {
		log.Error("hello dig here")
		log.Info("hello dig info")
	})
}

package nlog

import (
	"github.com/krilie/lico_alone/common/appdig"
	"github.com/krilie/lico_alone/component/ncfg"
	"testing"
)

var container = appdig.NewAppDig()

func TestMain(m *testing.M) {
	container.MustProvide(ncfg.NewNConfig())
	container.MustProvide(NewLogger)
	m.Run()
}

func TestLog(t *testing.T) {
	container.MustInvoke(func(log *NLog) {
		log.Error("hello dig here")
		log.Info("hello dig info")
	})
}

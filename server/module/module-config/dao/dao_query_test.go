// +build !auto_test

package dao

import (
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/common/utils/str_util"
	"github.com/krilie/lico_alone/component"
	"testing"
)

func TestMain(m *testing.M) {
	component.DigProviderTest()
	DigProvider()
	m.Run()
}

func TestConfigDao_GetAllConfig(t *testing.T) {
	dig.Container.MustInvoke(func(svc *ConfigDao) {
		config, err := svc.GetAllConfig(context.NewContext(), "")
		svc.log.Info(str_util.ToJson(config))
		t.Log(err)
	})
}

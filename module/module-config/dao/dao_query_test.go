package dao

import (
	"github.com/krilie/lico_alone/common/appdig"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/utils/jsonutil"
	"github.com/krilie/lico_alone/component"
	"testing"
)

var container = func() *appdig.AppContainer {
	return appdig.
		NewAppDig().
		MustProvides(component.DigComponentProviderAll).
		MustProvide(NewConfigDao)
}()

func TestConfigDao_GetAllConfig(t *testing.T) {
	container.MustInvoke(func(svc *ConfigDao) {
		config, err := svc.GetAllConfig(context.EmptyAppCtx(), "")
		svc.log.Info(jsonutil.ToJson(config))
		t.Log(err)
	})
}

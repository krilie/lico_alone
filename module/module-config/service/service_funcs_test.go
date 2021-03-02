package service

import (
	"github.com/krilie/lico_alone/common/appdig"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/component"
	"github.com/krilie/lico_alone/module/module-config/model"
	"testing"
)

var container = appdig.
	NewAppDig().
	MustProvides(component.DigComponentProviderAll).
	MustProvides(DigModuleConfigProviderAll)

func TestAddItems(t *testing.T) {
	container.MustInvoke(func(svc *ConfigModule) error {
		return svc.SetJsonValue(context.EmptyAppCtx(), model.ConfigItemsIcpInfo.Val(), model.IcpInfo{
			Name:  "1",
			Link:  "2",
			Label: "3",
		})
	})
}

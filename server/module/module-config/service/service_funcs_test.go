package service

import (
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/module/module-config/model"
	"testing"
)

func TestAddItems(t *testing.T) {
	dig.Container.MustInvoke(func(svc *ConfigModule) error {
		return svc.SetJsonValue(context.NewContext(), model.ConfigItemsIcpInfo.Val(), model.IcpInfo{
			Name:  "1",
			Link:  "2",
			Label: "3",
		})
	})
}

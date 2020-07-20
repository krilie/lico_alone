// +build !auto_test

package service

import (
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/component"
	"github.com/krilie/lico_alone/module/module-config/model"
	"testing"
)

func TestMain(m *testing.M) {
	component.DigProviderTest()
	DigProviderWithDao()
	m.Run()
}

func TestAddItems(t *testing.T) {
	dig.Container.MustInvoke(func(svc *ConfigModule) error {
		return svc.SetJsonValue(context.NewContext(), model.ConfigItemsIcpInfo.Val(), model.IcpInfo{
			Name:  "1",
			Link:  "2",
			Label: "3",
		})
	})
}

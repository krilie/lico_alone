package service_user

import (
	"github.com/krilie/lico_alone/common/appdig"
	"github.com/krilie/lico_alone/component"
	"github.com/krilie/lico_alone/module/module"
	"testing"
)

var userService *UserService

func TestMain(m *testing.M) {
	component.DigComponentProviderTest()
	module.DigProviderModule()
	DigProvider()
	appdig.Container.MustInvoke(func(svc *UserService) {
		userService = svc
	})
	m.Run()
}

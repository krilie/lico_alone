// +build !auto_test

package service_user

import (
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/component"
	"github.com/krilie/lico_alone/module/module"
	"testing"
)

var userService *UserService

func TestMain(m *testing.M) {
	component.DigProviderTest()
	module.DigProviderModule()
	DigProvider()
	dig.Container.MustInvoke(func(svc *UserService) {
		userService = svc
	})
	m.Run()
}

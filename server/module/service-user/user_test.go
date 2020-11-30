package service_user

import (
	"github.com/krilie/lico_alone/common/appdig"
	"github.com/krilie/lico_alone/component"
	"github.com/krilie/lico_alone/module/module"
)

var container = appdig.NewAppDig().
	MustProvides(component.DigComponentProviderAllForTest).
	MustProvides(module.DigProviderModuleAll).
	MustProvide(NewUserService)

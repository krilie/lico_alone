package service

import (
	"github.com/krilie/lico_alone/common/appdig"
	"github.com/krilie/lico_alone/module/module-user/dao"
)

// DigProvider provider
func DigProvider() {
	appdig.Container.MustProvide(NewUserModule)
}

func DigProviderAll() {
	dao.DigProvider()
	DigProvider()
}

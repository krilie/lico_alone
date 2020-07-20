package service

import (
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/module/module-user/dao"
)

// DigProvider provider
func DigProvider() {
	dig.Container.MustProvide(NewUserModule)
}

func DigProviderAll() {
	dao.DigProvider()
	DigProvider()
}

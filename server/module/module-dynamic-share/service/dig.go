package service

import (
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/module/module-dynamic-share/dao"
)

// DigProvider provider
func DigProvider() {
	dig.Container.MustProvide(NewDynamicShareModule)
}

func DigProviderAll() {
	dao.DigProvider()
	DigProvider()
}

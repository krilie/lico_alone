package service

import (
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/module/module-config/dao"
)

// DigProvider provider
func DigProvider() {
	dig.Container.MustProvide(NewConfigModule)
}

// DigProvider provider
func DigProviderWithDao() {
	dao.DigProvider()
	dig.Container.MustProvide(NewConfigModule)
}

package service

import (
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/module/module-customer/dao"
)

// DigProvider provider
func DigProvider() {
	dig.Container.MustProvide(NewCustomerModule)
}

func DigProviderAll() {
	dao.DigProvider()
	DigProvider()
}

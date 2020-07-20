package service

import (
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/module/module-file/dao"
)

// DigProvider provider
func DigProvider() {
	dig.Container.MustProvide(NewFileModule)
}

func DigProviderWithDao() {
	dao.DigProvider()
	DigProvider()
}

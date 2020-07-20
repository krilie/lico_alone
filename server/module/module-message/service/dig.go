package service

import (
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/module/module-message/dao"
	"github.com/krilie/lico_alone/module/module-message/infra/email"
	"github.com/krilie/lico_alone/module/module-message/infra/sms"
)

// DigProvider provider
func DigProvider() {
	dig.Container.MustProvide(NewMessageModule)
}

func DigProviderAll() {
	dao.DigProvider()
	email.DigProvider()
	sms.DigProvider()
	DigProvider()
}

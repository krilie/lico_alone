package service

import (
	"github.com/krilie/lico_alone/module/module-message/dao"
	"github.com/krilie/lico_alone/module/module-message/infra/email"
	"github.com/krilie/lico_alone/module/module-message/infra/sms"
)

var DigModuleMessageProviderAll = []interface{}{
	dao.NewMessageDao,
	email.NewIEmail,
	sms.NewAliSms2,
	NewMessageModule,
}

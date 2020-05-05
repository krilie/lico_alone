package service

import (
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-message/dao"
	"github.com/krilie/lico_alone/module/module-message/infra/email"
	"github.com/krilie/lico_alone/module/module-message/infra/sms"
)

func init() {
	dig.Container.MustProvide(func(log *nlog.NLog, dao *dao.MessageDao, email email.IEmail, sms sms.IAliSms) *MessageService {
		return NewMessageService(log, dao, email, sms)
	})
}

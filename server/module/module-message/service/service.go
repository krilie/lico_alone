package service

import (
	context_enum "github.com/krilie/lico_alone/common/com-model/context-enum"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-message/dao"
	"github.com/krilie/lico_alone/module/module-message/infra/email"
	"github.com/krilie/lico_alone/module/module-message/infra/sms"
)

type MessageService struct {
	Dao   *dao.MessageDao
	log   *nlog.NLog
	email email.IEmail
	sms   sms.IAliSms
}

func NewMessageService(log *nlog.NLog, dao *dao.MessageDao, email email.IEmail, sms sms.IAliSms) *MessageService {
	log = log.WithField(context_enum.Module.Str(), "module message service")
	return &MessageService{
		Dao:   dao,
		log:   log,
		email: email,
		sms:   sms,
	}
}

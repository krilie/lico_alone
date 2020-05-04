package service

import (
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
	return &MessageService{
		Dao:   dao,
		log:   log,
		email: email,
		sms:   sms,
	}
}

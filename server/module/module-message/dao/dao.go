package dao

import (
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/component/nlog"
)

func NewMessageDao(db *ndb.NDb, log *nlog.NLog) *MessageDao {
	return &MessageDao{
		NDb:               db,
		log:               log,
		IMessageEmail:     &messageEmail{NDb: db, log: log},
		IMessageSms:       &messageSms{NDb: db, log: log},
		IMessageValidCode: &messageValidCode{NDb: db, log: log},
	}
}

type MessageDao struct {
	*ndb.NDb
	log *nlog.NLog
	IMessageEmail
	IMessageSms
	IMessageValidCode
}

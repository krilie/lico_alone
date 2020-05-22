package dao

import (
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-message/model"
)

func NewMessageDao(db *ndb.NDb, log *nlog.NLog) *MessageDao {
	err := db.GetDb(context.NewContext()).
		AutoMigrate(
			new(model.MessageEmail),
			new(model.MessageSms),
			new(model.MessageValidCode)).
		Error
	if err != nil {
		panic(err)
	}
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

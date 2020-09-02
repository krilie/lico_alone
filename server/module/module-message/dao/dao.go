package dao

import (
	context_enum "github.com/krilie/lico_alone/common/com-model/context-enum"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/component/nlog"
)

func NewMessageDao(db *ndb.NDb, log *nlog.NLog) *MessageDao {
	log = log.WithField(context_enum.Module.Str(), "module message dao")
	//err := db.GetDb(context.NewContext()).
	//	AutoMigrate(
	//		new(model.MessageEmail),
	//		new(model.MessageSms),
	//		new(model.MessageValidCode))
	//if err != nil {
	//	panic(err)
	//}
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

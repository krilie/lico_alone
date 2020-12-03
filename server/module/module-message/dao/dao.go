package dao

import (
	context2 "context"
	context_enum "github.com/krilie/lico_alone/common/com-model/context-enum"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/global"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-message/model"
)

func NewMessageDao(db *ndb.NDb, log *nlog.NLog) *MessageDao {
	log = log.WithField(context_enum.Module.Str(), "module message dao")
	if global.EnableAutoMigrate {
		err := db.GetDb(context.NewAppCtx(context2.Background())).
			AutoMigrate(
				new(model.MessageEmail),
				new(model.MessageSms),
				new(model.MessageValidCode))
		if err != nil {
			panic(err)
		}
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

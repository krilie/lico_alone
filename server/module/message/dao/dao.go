package dao

import (
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/component/nlog"
)

func NewMessageDao(db *ndb.NDb, log *nlog.NLog) *MessageDao {
	return &MessageDao{
		NDb: db,
		log: log,
	}
}

type MessageDao struct {
	*ndb.NDb
	log *nlog.NLog
}

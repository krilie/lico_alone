package dao

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/component/nlog"
)

type UserDao struct {
	log *nlog.NLog
	*ndb.NDb
}

func NewUserDao(db *ndb.NDb, log *nlog.NLog) *UserDao {
	return &UserDao{
		log: log,
		NDb: db,
	}
}

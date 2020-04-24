package dao

import (
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/component/nlog"
)

func init() {
	dig.MustProvide(func(db *ndb.NDb, log *nlog.NLog) *UserDao {
		return NewUserDao(db, log)
	})
}

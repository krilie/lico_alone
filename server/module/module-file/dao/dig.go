package dao

import (
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/component/nlog"
)

func init() {
	dig.Container.MustProvide(func(db *ndb.NDb, log *nlog.NLog) *FileDao {
		return NewFileDao(db, log)
	})
}

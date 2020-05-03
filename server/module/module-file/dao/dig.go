package dao

import (
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-file/model"
)

func init() {
	dig.Container.MustProvide(func(db *ndb.NDb, log *nlog.NLog) *FileDao {
		err := db.GetDb(context.NewContext()).
			AutoMigrate(&model.FileMaster{}).Error
		if err != nil {
			panic(err)
		}
		return NewFileDao(db, log)
	})
}

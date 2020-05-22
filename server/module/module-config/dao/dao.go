package dao

import (
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-config/model"
)

type ConfigDao struct {
	*ndb.NDb
	log *nlog.NLog
}

func NewConfigDao(ndb *ndb.NDb, log *nlog.NLog) *ConfigDao {
	err := ndb.GetDb(context.NewContext()).AutoMigrate(&model.Config{}).Error
	if err != nil {
		panic(err)
	}
	return &ConfigDao{
		NDb: ndb,
		log: log,
	}
}

package dao

import (
	context_enum "github.com/krilie/lico_alone/common/com-model/context-enum"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/component/nlog"
)

type ConfigDao struct {
	*ndb.NDb
	log *nlog.NLog
}

func NewConfigDao(ndb *ndb.NDb, log *nlog.NLog) *ConfigDao {
	log = log.WithField(context_enum.Module.Str(), "config dao")
	//err := ndb.GetDb(context.NewContext()).AutoMigrate(&model.Config{})
	//if err != nil {
	//	panic(err)
	//}
	return &ConfigDao{
		NDb: ndb,
		log: log,
	}
}

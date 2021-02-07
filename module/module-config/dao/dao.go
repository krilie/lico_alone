package dao

import (
	context2 "context"
	context_enum "github.com/krilie/lico_alone/common/com-model/context-enum"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/global"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-config/model"
)

type ConfigDao struct {
	*ndb.NDb
	log *nlog.NLog
}

func NewConfigDao(ndb *ndb.NDb, log *nlog.NLog) *ConfigDao {
	log = log.WithField(context_enum.Module.Str(), "config dao")
	if global.EnableAutoMigrate {
		err := ndb.GetDb(context.NewAppCtx(context2.Background())).AutoMigrate(&model.Config{})
		if err != nil {
			panic(err)
		}
	}
	return &ConfigDao{
		NDb: ndb,
		log: log,
	}
}

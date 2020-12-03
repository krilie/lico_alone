package dao

import (
	context2 "context"
	context_enum "github.com/krilie/lico_alone/common/com-model/context-enum"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/global"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-dynamic-share/model"
)

type DynamicShareDao struct {
	*ndb.NDb
	log *nlog.NLog
}

func NewDynamicShareDao(db *ndb.NDb, log *nlog.NLog) *DynamicShareDao {
	log = log.WithField(context_enum.Module.Str(), "module DynamicShareDao dao")
	if global.EnableAutoMigrate {
		err := db.GetDb(context.NewAppCtx(context2.Background())).AutoMigrate(new(model.DynamicShare), new(model.DynamicShareLabel))
		if err != nil {
			panic(err)
		}
	}
	return &DynamicShareDao{
		NDb: db,
		log: log,
	}
}

package dao

import (
	context2 "context"
	context_enum "github.com/krilie/lico_alone/common/com-model/context-enum"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/global"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-catchword/model"
)

type CatchwordDao struct {
	*ndb.NDb
	log *nlog.NLog
}

func NewCatchwordDao(ndb *ndb.NDb, log *nlog.NLog) *CatchwordDao {
	log = log.WithField(context_enum.Module.Str(), "Catchword dao")
	if global.EnableAutoMigrate {

		err0 := ndb.GetDb(context.NewAppCtx(context2.Background())).AutoMigrate(&model.Catchword{})
		if err0 != nil {
			panic(err0)
		}

	}
	return &CatchwordDao{
		NDb: ndb,
		log: log,
	}
}

package dao

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	context_enum "github.com/krilie/lico_alone/common/com-model/context-enum"
	context2 "github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/global"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/component/nlog"
)

type ZUnionDao struct {
	log *nlog.NLog
	*ndb.NDb
}

func NewZUnionDao(db *ndb.NDb, log *nlog.NLog) *ZUnionDao {
	log = log.WithField(context_enum.Module.Str(), "module user dao")
	if global.EnableAutoMigrate {
		err := db.GetDb(context2.NewAppCtx(context.Background())).AutoMigrate()
		if err != nil {
			panic(err)
		}
	}
	return &ZUnionDao{
		log: log,
		NDb: db,
	}
}

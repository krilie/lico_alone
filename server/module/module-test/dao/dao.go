package dao

import (
	context2 "context"
	context_enum "github.com/krilie/lico_alone/common/com-model/context-enum"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/global"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-test/model"
)

type TestArticleDao struct {
	*ndb.NDb
	log *nlog.NLog
}

func NewTestArticleDao(ndb *ndb.NDb, log *nlog.NLog) *TestArticleDao {
	log = log.WithField(context_enum.Module.Str(), "TestArticle dao")
	if global.EnableAutoMigrate {

		err0 := ndb.GetDb(context.NewAppCtx(context2.Background())).AutoMigrate(&model.BaseOne{})
		if err0 != nil {
			panic(err0)
		}

		err1 := ndb.GetDb(context.NewAppCtx(context2.Background())).AutoMigrate(&model.BaseTwo{})
		if err1 != nil {
			panic(err1)
		}

	}
	return &TestArticleDao{
		NDb: ndb,
		log: log,
	}
}

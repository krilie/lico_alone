package dao

import (
	"github.com/krilie/lico_alone/common/com-model/context-enum"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-statistic/model"
)

// StatisticDao 统计用户访问次数等内容
type StatisticDao struct {
	*ndb.NDb
	log *nlog.NLog
}

func NewStatisticDao(db *ndb.NDb, log *nlog.NLog) *StatisticDao {
	log = log.WithField(context_enum.Module.Str(), "StatisticDao")
	err := db.GetDb(context.NewContext()).AutoMigrate(new(model.StatVisitorLogs), new(model.StatArticleVisitorLogs))
	if err != nil {
		panic(err)
	}
	return &StatisticDao{
		NDb: db,
		log: log,
	}
}

func init() {
	dig.Container.MustProvide(NewStatisticDao)
}

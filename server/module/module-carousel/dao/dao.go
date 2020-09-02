package dao

import (
	"github.com/krilie/lico_alone/common/com-model/context-enum"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/component/nlog"
)

type CarouselDao struct {
	*ndb.NDb
	log *nlog.NLog
}

func NewCarouseDao(db *ndb.NDb, log *nlog.NLog) *CarouselDao {
	log = log.WithField(context_enum.Module.Str(), "CarouselDao")
	//err := db.GetDb(context.NewContext()).AutoMigrate(&model.Carousel{})
	//if err != nil {
	//	panic(err)
	//}
	return &CarouselDao{
		NDb: db,
		log: log,
	}
}

package dao

import (
	"github.com/krilie/lico_alone/common/com-model/context-enum"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-carousel/model"
)

type CarouselDao struct {
	*ndb.NDb
	log *nlog.NLog
}

func NewCarouseDao(db *ndb.NDb, log *nlog.NLog) *CarouselDao {
	log = log.WithField(context_enum.Module.Str(), "CarouselDao")
	err := db.GetDb(context.NewContext()).AutoMigrate(&model.Carousel{})
	if err != nil {
		panic(err)
	}
	return &CarouselDao{
		NDb: db,
		log: log,
	}
}

func init() {
	dig.Container.MustProvide(NewCarouseDao)
}

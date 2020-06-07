package service

import (
	"github.com/krilie/lico_alone/common/com-model/context-enum"
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-carousel/dao"
)

type CarouselModule struct {
	Dao *dao.CarouselDao
	log *nlog.NLog
}

func NewCarouselModule(dao *dao.CarouselDao, log *nlog.NLog) *CarouselModule {
	log = log.WithField(context_enum.Module.Str(), "CarouselModule")
	return &CarouselModule{
		Dao: dao,
		log: log,
	}
}

func init() {
	dig.Container.MustProvide(NewCarouselModule)
}

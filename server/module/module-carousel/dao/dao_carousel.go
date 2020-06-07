package dao

import (
	"context"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/common/utils/str_util"
	"github.com/krilie/lico_alone/module/module-carousel/model"
)

func (a *CarouselDao) QueryCarousel(ctx context.Context, isOnShow *bool) (list []*model.Carousel, err error) {
	log := a.log.Get(ctx).WithFuncName("QueryCarousel").WithField("is_on_show", isOnShow)
	log.Trace("on query carousel")
	db := a.GetDb(ctx).Model(new(model.Carousel))
	if isOnShow != nil {
		db = db.Where("is_on_show=?", isOnShow)
	}
	db.Order("created_at desc")
	list = make([]*model.Carousel, 0)
	err = db.Find(&list).Error
	if err != nil {
		log.WithField("err", err).Error("db operation err")
		return nil, errs.NewInternal().WithError(err)
	}
	return list, nil
}

func (a *CarouselDao) CreateCarousel(ctx context.Context, item *model.Carousel) error {
	log := a.log.Get(ctx).WithFuncName("CreateCarousel")
	log.WithField("param", str_util.ToJson(item)).Info("params data")

	if item.Id == "" {
		item.Id = id_util.GetUuid()
	}
	err := a.GetDb(ctx).Model(new(model.Carousel)).Create(item).Error
	if err != nil {
		log.WithField("err", err).Error("db err " + err.Error())
		return errs.NewInternal().WithError(err)
	}
	return nil
}

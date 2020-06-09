package dao

import (
	"context"
	"github.com/jinzhu/gorm"
	com_model "github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/common/utils/str_util"
	"github.com/krilie/lico_alone/module/module-carousel/model"
	"time"
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

func (a *CarouselDao) UpdateCarousel(ctx context.Context, item *model.UpdateCarouselModel) error {
	log := a.log.Get(ctx).WithFuncName("UpdateCarousel")
	var m = new(model.Carousel)
	err := a.GetDb(ctx).Model(new(model.Carousel)).Where("id=?", item.Id).Find(&m).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errs.NewNotExistsError().WithError(err).WithMsg("not found")
		}
		log.Error(err)
		return err
	}
	m.UpdatedAt = time.Now()
	m.Message = item.Message
	m.Url = item.Url
	m.IsOnShow = item.IsOnShow
	err = a.GetDb(ctx).Update(m).Error
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (a *CarouselDao) DeleteCarouselById(ctx context.Context, id string) error {
	err := a.GetDb(ctx).Delete(&model.Carousel{
		Model: com_model.Model{
			Id: id,
		},
	}).Error
	if err != nil {
		a.log.Error(err)
		return err
	}
	return nil
}

type ICarouselDao interface {
	QueryCarousel(ctx context.Context, isOnShow *bool) (list []*model.Carousel, err error)
	CreateCarousel(ctx context.Context, item *model.Carousel) error
	UpdateCarousel(ctx context.Context, item *model.UpdateCarouselModel) error
	DeleteCarouselById(ctx context.Context, id string) error
}

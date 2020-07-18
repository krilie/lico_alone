package service

import (
	"context"
	com_model "github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/module/module-carousel/model"
	"gorm.io/gorm"
	"time"
)

func (c *CarouselModule) QueryCarousel(ctx context.Context, isOnShow *bool) (list []*model.Carousel, err error) {
	return c.Dao.QueryCarousel(ctx, isOnShow)
}

func (c *CarouselModule) CreateCarousel(ctx context.Context, item *model.CreateCarouselModel) error {
	return c.Dao.CreateCarousel(ctx, &model.Carousel{
		Model: com_model.Model{
			Id:        id_util.GetUuid(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt{},
		},
		Message:  item.Message,
		Url:      item.Url,
		IsOnShow: item.IsOnShow,
	})
}

func (c *CarouselModule) UpdateCarousel(ctx context.Context, item *model.UpdateCarouselModel) error {
	return c.Dao.UpdateCarousel(ctx, item)
}

func (c *CarouselModule) DeleteCarouselById(ctx context.Context, id string) error {
	return c.Dao.DeleteCarouselById(ctx, id)
}

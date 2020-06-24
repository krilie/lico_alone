package dao

import (
	com_model "github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/common/utils/str_util"
	"github.com/krilie/lico_alone/module/module-carousel/model"
	"testing"
	"time"
)

func TestCarouselDao_QueryCarousel(t *testing.T) {
	dig.Container.MustInvoke(func(dao *CarouselDao) {
		ctx := context.NewContext()
		carousel := &model.Carousel{
			Model: com_model.Model{
				Id:        "",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
				DeletedAt: &time.Time{},
			},
			Message:  "",
			Url:      "",
			IsOnShow: false,
		}
		err2 := dao.CreateCarousel(ctx, carousel)
		if err2 != nil {
			t.Log(err2.Error())
			return
		}

		carousels, err := dao.QueryCarousel(ctx, nil)
		t.Log(str_util.ToJsonPretty(carousels))
		t.Log(err)
	})
}

func TestCarouselDao_UpdateCarousel(t *testing.T) {
	dig.Container.MustInvoke(func(dao *CarouselDao) {
		ctx := context.NewContext()
		err := dao.UpdateCarousel(ctx, &model.UpdateCarouselModel{
			Id:       "2e415049-feaa-45fe-8482-22712155253b",
			Message:  "333333333333344444444433333333333",
			Url:      "4444443244444444444444444444444444",
			IsOnShow: false,
		})
		t.Log(err)
	})
}

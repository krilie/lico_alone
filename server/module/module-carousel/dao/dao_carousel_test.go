package dao

import (
	"fmt"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/common/utils/str_util"
	"github.com/krilie/lico_alone/component"
	"github.com/krilie/lico_alone/module/module-carousel/model"
	"testing"
)

func TestMain(m *testing.M) {
	component.DigProviderTest()
	DigProvider()
	m.Run()
}

func TestCarouselDao_QueryCarousel(t *testing.T) {
	dig.Container.MustInvoke(func(dao *CarouselDao) {
		carousels, err := dao.QueryCarousel(context.NewContext(), nil)
		fmt.Println(str_util.ToJsonPretty(carousels))
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

package dao

import (
	"fmt"
	"github.com/krilie/lico_alone/common/appdig"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/utils/str_util"
	"github.com/krilie/lico_alone/component"
	"github.com/krilie/lico_alone/module/module-carousel/model"
	"testing"
)

var container = func() *appdig.AppContainer {
	container := appdig.NewAppDig()
	container.MustProvides(component.DigComponentProviderAllForTest)
	container.MustProvide(NewCarouseDao)
	return container
}()

func TestCarouselDao_QueryCarousel(t *testing.T) {
	container.MustInvoke(func(dao *CarouselDao) {
		carousels, err := dao.QueryCarousel(context.EmptyAppCtx(), nil)
		fmt.Println(str_util.ToJsonPretty(carousels))
		t.Log(err)
	})
}

func TestCarouselDao_UpdateCarousel(t *testing.T) {
	container.MustInvoke(func(dao *CarouselDao) {
		ctx := context.EmptyAppCtx()
		err := dao.UpdateCarousel(ctx, &model.UpdateCarouselModel{
			Id:       "2e415049-feaa-45fe-8482-22712155253b",
			Message:  "333333333333344444444433333333333",
			Url:      "4444443244444444444444444444444444",
			IsOnShow: false,
		})
		t.Log(err)
	})
}

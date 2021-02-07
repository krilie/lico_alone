package service_common

import (
	"context"
	"github.com/krilie/lico_alone/module/module-carousel/model"
)

func (a *CommonService) QueryCarousel(ctx context.Context) (list []*model.Carousel, err error) {
	isOnShow := true
	return a.ModuleCarousel.QueryCarousel(ctx, &isOnShow)
}

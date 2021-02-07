package service_user

import (
	"context"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/module/module-carousel/model"
)

func (a *UserService) QueryCarousel(ctx context.Context, isOnShow *bool) (list []*model.Carousel, err error) {
	return a.ModuleCarousel.QueryCarousel(ctx, isOnShow)
}

func (a *UserService) CreateCarousel(ctx context.Context, item *model.CreateCarouselModel) error {
	return a.ModuleCarousel.CreateCarousel(ctx, item)
}

func (a *UserService) UpdateCarousel(ctx context.Context, item *model.UpdateCarouselModel) error {
	return a.ModuleCarousel.UpdateCarousel(ctx, item)
}

func (a *UserService) DeleteCarouselById(ctx context.Context, id string) error {
	if id == "" {
		return errs.NewParamError().WithMsg("id is empty")
	}
	return a.ModuleCarousel.DeleteCarouselById(ctx, id)
}

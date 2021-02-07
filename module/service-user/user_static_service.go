package service_user

import (
	"context"
	"github.com/krilie/lico_alone/module/module-statistic/model"
)

func (a *UserService) GetAllVisitorPoint(ctx context.Context) ([]model.VisitorLonlatModel, error) {
	lat, err := a.ModuleStatistic.QueryAllVisitorLonLat(ctx)
	return lat, err
}

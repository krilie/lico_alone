package service

import (
	"context"
	"github.com/krilie/lico_alone/module/module-config/model"
)

func (a *ConfigModule) GetAllConfig(ctx context.Context, searchKey string) ([]*model.Config, error) {
	return a.Dao.GetAllConfig(ctx, searchKey)
}

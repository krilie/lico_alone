package dao

import (
	"context"
	"github.com/krilie/lico_alone/module/module-config/model"
)

func (a *ConfigDao) GetAllConfig(ctx context.Context, searchKey string) (configs []*model.Config, err error) {
	configs = []*model.Config{}
	err = a.GetDb(ctx).Model(&model.Config{}).Find(&configs).Error
	return
}

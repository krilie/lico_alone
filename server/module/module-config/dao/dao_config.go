package dao

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/module/module-config/model"
)

func (a *ConfigDao) GetConfigByName(ctx context.Context, name string) (*model.Config, error) {
	config := new(model.Config)
	err := a.GetDb(ctx).Where("name=?", name).Find(config).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, errs.NewInternal().WithError(err)
	}
	return config, nil
}

func (a *ConfigDao) CreateConfig(ctx context.Context, config *model.Config) error {
	err := a.GetDb(ctx).Create(config).Error
	if err != nil {
		return errs.NewInternal().WithError(err)
	}
	return nil
}

func (a *ConfigDao) DeleteConfig(ctx context.Context, name string) error {
	err := a.GetDb(ctx).Where("name=?", name).Delete(new(model.Config)).Error
	if err != nil {
		return errs.NewInternal().WithError(err)
	}
	return nil
}
func (a *ConfigDao) UpdateConfig(ctx context.Context, config *model.Config) error {
	err := a.GetDb(ctx).Model(config).Where("name=?", config.Name).Omit("create_time").Update(config).Error
	if err != nil {
		return errs.NewInternal().WithError(err)
	}
	return nil
}

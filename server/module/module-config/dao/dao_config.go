package dao

import (
	"context"
	"errors"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/module/module-config/model"
	"gorm.io/gorm"
)

func (a *ConfigDao) GetConfigByName(ctx context.Context, name string) (*model.Config, error) {
	config := new(model.Config)
	err := a.GetDb(ctx).Where("name=?", name).Find(config).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
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

func (a *ConfigDao) DeleteAllConfig(ctx context.Context) error {
	err := a.GetDb(ctx).Delete(new(model.Config)).Error
	if err != nil {
		return errs.NewInternal().WithError(err)
	}
	return nil
}

func (a *ConfigDao) UpdateConfig(ctx context.Context, config *model.Config) error {
	err := a.GetDb(ctx).Model(config).Where("name=?", config.Name).Omit("create_time").Updates(config).Error
	if err != nil {
		return errs.NewInternal().WithError(err)
	}
	return nil
}

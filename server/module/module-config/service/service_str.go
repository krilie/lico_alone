package service

import (
	"context"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/module/module-config/model"
	"time"
)

func (a *ConfigModule) GetValueStr(ctx context.Context, name string) (*string, error) {
	config, err := a.Dao.GetConfigByName(ctx, name)
	if err != nil {
		return nil, errs.NewInternal().WithError(err)
	}
	if config == nil {
		return nil, nil
	}
	return &config.Value, nil
}
func (a *ConfigModule) SetValueStr(ctx context.Context, name string, value string) error {
	config, err := a.Dao.GetConfigByName(ctx, name)
	if err != nil {
		return err
	}
	if config == nil {
		return a.Dao.CreateConfig(ctx, &model.Config{
			CreateTime: time.Now(),
			Name:       name,
			Value:      value,
		})
	} else {
		config.Value = value
		return a.Dao.UpdateConfig(ctx, config)
	}
}

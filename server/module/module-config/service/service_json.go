package service

import (
	"context"
	"encoding/json"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/utils/str_util"
	"github.com/krilie/lico_alone/module/module-config/model"
	"time"
)

func (a *ConfigService) GetJsonValue(ctx context.Context, name string, resOut interface{}) (content *model.Config, err error) {
	config, err := a.Dao.GetConfigByName(ctx, name)
	if err != nil {
		return nil, errs.NewInternal().WithError(err)
	}
	if config == nil {
		return nil, nil
	}
	err = json.Unmarshal([]byte(config.Value), resOut)
	if err != nil {
		return nil, errs.NewInternal().WithError(err)
	}
	return config, nil
}
func (a *ConfigService) SetJsonValue(ctx context.Context, name string, value interface{}) error {
	config, err := a.Dao.GetConfigByName(ctx, name)
	if err != nil {
		return err
	}
	if config == nil {
		return a.Dao.CreateConfig(ctx, &model.Config{
			CreateTime: time.Now(),
			Name:       name,
			Value:      str_util.ToJson(value),
		})
	} else {
		config.Value = str_util.ToJson(value)
		return a.Dao.UpdateConfig(ctx, config)
	}
}

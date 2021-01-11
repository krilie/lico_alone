package service_user

import (
	"context"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/module/module-config/model"
)

// 分页查询
// 设置修改字符串
// 不能删除 不能添加

func (a *UserService) GetAllConfig(ctx context.Context, searchKey string) ([]*model.Config, error) {
	a.log.Get(ctx).Info("on get all config func")
	return a.moduleConfig.GetAllConfig(ctx, searchKey)
}

func (a *UserService) UpdateConfig(ctx context.Context, name, value string) error {
	config, err := a.moduleConfig.Dao.GetConfigByName(ctx, name)
	if err != nil {
		return err
	}
	if config == nil {
		return errs.NewNotExistsError().WithMsg("配置未找到")
	}
	config.Value = value
	return a.moduleConfig.Dao.UpdateConfig(ctx, config)
}

func (a *UserService) GetAMapKey(ctx context.Context) (key string, err error) {
	config, err := a.moduleConfig.Dao.GetConfigByName(ctx, model.ConfigItemsAMapKey.Val())
	if err != nil {
		return "", err
	}
	if config == nil {
		return "", errs.NewNormal().WithMsg("未找到AMapKey")
	}
	return config.Value, nil
}

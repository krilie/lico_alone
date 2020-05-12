package user_service

import (
	"context"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/module/module-config/model"
)

// 分页查询
// 设置修改字符串
// 不能删除 不能添加

func (a *UserService) GetAllConfig(ctx context.Context, searchKey string) ([]*model.Config, error) {
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

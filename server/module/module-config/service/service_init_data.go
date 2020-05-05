package service

import (
	"context"
	"github.com/krilie/lico_alone/module/module-config/model"
)

// InitConfigData 初始化配置文件 不会删除原来的数据
func (a *ConfigService) InitConfigData(ctx context.Context) error {
	err := a.Dao.Transaction(ctx, func(ctx context.Context) error {
		err := a.Dao.DeleteAllConfig(ctx)
		if err != nil {
			return err
		}
		err = a.SetValueBool(ctx, model.ConfigItemsIsInitData.Value(), true)
		if err != nil {
			return err
		}
		err = a.SetValueStr(ctx, model.ConfigItemsNotificationEmail.Value(), "example@example.com")
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

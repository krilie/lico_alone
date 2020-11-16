package service

import (
	"context"
	"github.com/krilie/lico_alone/module/module-config/model"
)

// InitConfigData 初始化配置文件 不会删除原来的数据
func (a *ConfigModule) InitConfigData(ctx context.Context) error {
	err := a.Dao.Transaction(ctx, func(ctx context.Context) error {
		err := a.Dao.DeleteAllConfig(ctx)
		if err != nil {
			return err
		}
		err = a.SetValueBool(ctx, model.ConfigItemsIsInitData.Val(), true)
		if err != nil {
			return err
		}
		err = a.SetValueStr(ctx, model.ConfigItemsNotificationEmail.Val(), "example@example.com")
		if err != nil {
			return err
		}
		err = a.SetJsonValue(ctx, model.ConfigItemsIcpInfo.Val(), model.IcpInfo{Name: "1", Label: "2", Link: "3"})
		if err != nil {
			return err
		}
		err = a.SetValueStr(ctx, model.ConfigItemsAboutApp.Val(), "todo:关于本站")
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

package init_data_service

import (
	"context"
	"github.com/krilie/lico_alone/module/module-config/model"
)

// InitData 初始化需要初始化的数据
func (initData *InitDataService) InitData(ctx context.Context) {
	err := initData.GetNDb(ctx).Transaction(ctx, func(ctx context.Context) error {
		if !initData.IsInit(ctx) {
			// init functions
			err := initData.unionService.ModuleConfig.InitConfigData(ctx)
			if err != nil {
				return err
			}
			err = initData.unionService.ModuleUser.InitUserData(ctx)
			if err != nil {
				return err
			}
			return nil
		} else {
			return nil
		}
	})
	if err != nil {
		panic(err)
	}
}

// IsInit 是否有被初始化
func (initData *InitDataService) IsInit(ctx context.Context) bool {
	valueBool, err := initData.unionService.ModuleConfig.GetValueBool(ctx, model.ConfigItemsIsInitData.Val())
	if err != nil {
		panic(err)
	}
	if valueBool == nil {
		err := initData.unionService.ModuleConfig.SetValueBool(ctx, model.ConfigItemsIsInitData.Val(), false)
		if err != nil {
			panic(err)
		}
		return false
	}
	return *valueBool
}

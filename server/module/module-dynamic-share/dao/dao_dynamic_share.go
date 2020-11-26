package dao

import (
	"context"
	"github.com/krilie/lico_alone/module/module-dynamic-share/model"
)

func (a *DynamicShareDao) AddDynamicShare(ctx context.Context, m model.DynamicShare) error {
	err := a.GetDb(ctx).Create(m).Error
	return err
}

func (a *DynamicShareDao) DeleteDynamicShare(ctx context.Context, ids []string) error {
	err := a.GetDb(ctx).Where("id in ?", ids).Delete(&model.DynamicShare{}).Error
	return err
}

func (a *DynamicShareDao) UpdateDynamicShare(ctx context.Context, u model.UpdateDynamicShareModel) error {
	err := a.GetDb(ctx).Updates(u).Error
	return err
}

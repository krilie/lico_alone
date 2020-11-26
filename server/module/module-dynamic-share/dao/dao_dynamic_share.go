package dao

import (
	"context"
	"github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/module/module-dynamic-share/model"
)

func (a *DynamicShareDao) AddDynamicShare(ctx context.Context, m model.CreateDynamicShareModel) error {
	err := a.GetDb(ctx).Create(model.DynamicShare{
		Model:   com_model.NewModel(),
		Content: m.Content,
		Sort:    m.Sort,
	}).Error
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

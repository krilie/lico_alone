package dao

import (
	"context"
	"github.com/krilie/lico_alone/common/errs"
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
	affected, err := a.Exec(ctx,
		"update tb_dynamic_share set content=?,sort=? where id=? and deleted_at is null",
		u.Content, u.Sort, u.Id)
	if err != nil {
		return err
	}
	if affected == 0 {
		return errs.NewNormal().WithMsg("no updated")
	}
	return nil
}

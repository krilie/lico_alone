package dao

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/module/module-user/model"
)

type IPerm interface {
	GetPermByName(ctx context.Context, name string) (*model.Permission, error)
	DeletePermByName(ctx context.Context, name string) error
	CreatePerm(ctx context.Context, item *model.Permission) error
}

func (d *UserDao) GetPermByName(ctx context.Context, name string) (*model.Permission, error) {
	perm := new(model.Permission)
	err := d.GetDb(ctx).Model(&model.Permission{}).Where(&model.Permission{
		Name: name,
	}).Find(perm).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, errs.NewInternal().WithError(err)
	}
	return perm, nil
}

// DeletePermByName just delete it
func (d *UserDao) DeletePermByName(ctx context.Context, name string) error {
	err := d.GetDb(ctx).Model(&model.Permission{}).Where("name=?", name).Delete(&model.Permission{}).Error
	if err != nil {
		return errs.NewInternal().WithError(err)
	}
	return nil
}

func (d *UserDao) CreatePerm(ctx context.Context, item *model.Permission) error {
	err := d.GetDb(ctx).Model(&model.Permission{}).Create(item).Error
	if err != nil {
		return errs.NewInternal().WithError(err)
	}
	return nil
}

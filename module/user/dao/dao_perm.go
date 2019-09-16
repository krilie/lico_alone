package dao

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/module/user/model"
)

type IPerm interface {
	GetPermByName(ctx context.Context, name string) (*model.Permission, error)
	DeletePermByName(ctx context.Context, name string) error
	CreatePerm(ctx context.Context, item *model.Permission) error
}

func (d *Dao) GetPermByName(ctx context.Context, name string) (*model.Permission, error) {
	perm := new(model.Permission)
	err := d.Db.Model(&model.Permission{}).Where(&model.Permission{
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
func (d *Dao) DeletePermByName(ctx context.Context, name string) error {
	err := d.Db.Model(&model.Permission{}).Where("name=?", name).Delete(&model.Permission{}).Error
	if err != nil {
		return errs.NewInternal().WithError(err)
	}
	return nil
}

func (d *Dao) CreatePerm(ctx context.Context, item *model.Permission) error {
	err := d.Db.Model(&model.Permission{}).Create(item).Error
	if err != nil {
		return errs.NewInternal().WithError(err)
	}
	return nil
}

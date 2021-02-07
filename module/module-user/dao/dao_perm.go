package dao

import (
	"context"
	"errors"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/module/module-user/model"
	"gorm.io/gorm"
)

type IPerm interface {
	GetPermByName(ctx context.Context, name string) (*model.Permission, error)
	GetPermByMethodPath(ctx context.Context, method, path string) (*model.Permission, error)
	DeletePermByName(ctx context.Context, name string) error
	CreatePerm(ctx context.Context, item *model.Permission) error
	CreatePerms(ctx context.Context, items []model.Permission) error
}

func (d *UserDao) GetPermByName(ctx context.Context, name string) (*model.Permission, error) {
	perm := new(model.Permission)
	err := d.GetDb(ctx).Model(&model.Permission{}).Where(&model.Permission{
		Name: name,
	}).Find(perm).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
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

//CreatePerms
func (d *UserDao) CreatePerms(ctx context.Context, items []model.Permission) error {
	for i := range items {
		err := d.CreatePerm(ctx, &items[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *UserDao) GetPermByMethodPath(ctx context.Context, method, path string) (*model.Permission, error) {
	perm := new(model.Permission)
	err := d.GetDb(ctx).Model(&model.Permission{}).Where(&model.Permission{
		RefMethod: method,
		RefPath:   path,
	}).Find(perm).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errs.NewInternal().WithError(err)
	}
	return perm, nil
}

package dao

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/module/module-user/model"
)

type IRole interface {
	GetRoleByName(ctx context.Context, name string) (*model.Role, error)
	GetRolesByParentName(ctx context.Context, pName string) ([]*model.Role, error)
	CreateRole(ctx context.Context, item *model.Role) error
	GetAllRole(ctx context.Context, parents ...string) ([]*model.Role, error)
}

func (d *UserDao) GetRoleByName(ctx context.Context, name string) (*model.Role, error) {
	role := new(model.Role)
	err := d.GetDb(ctx).Model(role).Where("name=?", name).Find(role).Error
	if err != nil && gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	if err != nil {
		return nil, errs.NewInternal().WithError(err)
	}
	return role, nil
}

func (d *UserDao) GetRolesByParentName(ctx context.Context, pName string) ([]*model.Role, error) {
	role := new([]*model.Role)
	err := d.GetDb(ctx).Model(role).Where("parent_name=?", pName).Find(role).Error
	if err != nil && gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	if err != nil {
		return nil, errs.NewInternal().WithError(err)
	}
	return *role, nil
}

func (d *UserDao) CreateRole(ctx context.Context, item *model.Role) error {
	err := d.GetDb(ctx).Model(&model.Role{}).Create(item).Error
	if err != nil {
		return errs.NewInternal().WithError(err)
	}
	return nil
}

func (d *UserDao) GetAllRole(ctx context.Context, parents ...string) ([]*model.Role, error) {
	var list []*model.Role
	err := d.GetDb(ctx).Model(&model.Role{}).Where("parent_name in (?)", parents).Find(&list).Error
	if err != nil {
		return nil, errs.NewInternal().WithError(err)
	}
	return list, nil
}

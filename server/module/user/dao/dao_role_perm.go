package dao

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/model"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/module/user/model"
	"time"
)

type IRolePerm interface {
	CreateRolePerm(ctx context.Context, roleName, permName string) error
	DeleteRolePerm(ctx context.Context, roleName, permName string) error
	HasRolePerm(ctx context.Context, roleName, permName string) (bool, error)
	GetRolePerm(ctx context.Context, roleName, permName string) (*model.RolePermission, error)
	GetRolePermsByRoleName(ctx context.Context, roleName string) ([]*model.RolePermission, error)
}

func (d *Dao) CreateRolePerm(ctx context.Context, roleName, permName string) error {
	err := d.Db.Model(&model.RolePermission{}).Create(&model.RolePermission{
		Model: model.Model{
			Id:         id_util.NextSnowflake(),
			CreateTime: time.Now(),
		},
		RoleName:       roleName,
		PermissionName: permName,
	}).Error
	if err != nil {
		return errs.NewErrDbCreate().WithError(err)
	}
	return nil
}

func (d *Dao) DeleteRolePerm(ctx context.Context, roleName, permName string) error {
	err := d.Db.Model(&model.RolePermission{}).Where(&model.RolePermission{
		RoleName:       roleName,
		PermissionName: permName,
	}).Delete(&model.RolePermission{}).Error
	if err != nil {
		return errs.NewErrDbDelete().WithError(err)
	}
	return nil
}

func (d *Dao) HasRolePerm(ctx context.Context, roleName, permName string) (bool, error) {
	count := 0
	err := d.Db.Model(&model.RolePermission{}).Where(&model.RolePermission{RoleName: roleName, PermissionName: permName}).Count(&count).Error
	if err != nil {
		return false, errs.NewErrDbQuery().WithError(err)
	}
	return count != 0, nil
}

func (d *Dao) GetRolePerm(ctx context.Context, roleName, permName string) (*model.RolePermission, error) {
	item := new(model.RolePermission)
	err := d.Db.Model(&model.RolePermission{}).Where(&model.RolePermission{
		RoleName:       roleName,
		PermissionName: permName,
	}).Find(item).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, errs.NewErrDbQuery().WithError(err)
	}
	return item, nil
}

func (d *Dao) GetRolePermsByRoleName(ctx context.Context, roleName string) ([]*model.RolePermission, error) {
	var list []*model.RolePermission
	err := d.Db.Model(list).Where("role_name=?", roleName).Find(&list).Error
	if err != nil {
		return nil, errs.NewErrDbQuery().WithError(err)
	}
	return list, nil
}

func (d *Dao) GetRolePermsByRolesWithPermName(ctx context.Context, permName string, roleNames ...string) ([]*model.RolePermission, error) {
	var list []*model.RolePermission
	err := d.Db.Model(&model.RolePermission{}).Where(&model.RolePermission{PermissionName: permName}).Where("role_name in (?)", roleNames).Find(&list).Error
	if err != nil {
		return nil, errs.NewErrDbQuery().WithError(err)
	}
	return list, nil
}

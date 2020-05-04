package dao

import (
	"context"
	"github.com/jinzhu/gorm"
	common_model "github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/module/module-user/model"
	"time"
)

type IUserRole interface {
	HasUserRoleByName(ctx context.Context, userId string, roleName string) (bool, error)
	GetUserRoleByName(ctx context.Context, userId string, roleName string) (*model.UserRole, error)
	GetUserRolesByUserId(ctx context.Context, userId string) ([]*model.UserRole, error)
	CreateUserRole(ctx context.Context, userId string, roleId string) error
	DeleteUserRole(ctx context.Context, userId string, roleId string) error
	GetAllUserRole(ctx context.Context) ([]*model.UserRole, error)
	GetAllRolePermission(ctx context.Context) ([]*model.RolePermission, error)
}

func (d *UserDao) HasUserRoleByName(ctx context.Context, userId string, roleName string) (bool, error) {
	userRoleCount := 0
	err := d.GetDb(ctx).Model(&model.UserRole{}).Where("user_id=? and role_id=?",
		userId,
		d.GetDb(ctx).Model(&model.Role{}).Where("name=?", roleName).Select("id").SubQuery(),
	).Count(&userRoleCount).Error
	if err != nil {
		return false, errs.NewInternal().WithError(err)
	}
	if userRoleCount == 0 {
		return false, nil
	}
	return true, nil
}

func (d *UserDao) GetUserRoleByName(ctx context.Context, userId string, roleName string) (*model.UserRole, error) {
	userRole := new(model.UserRole)
	err := d.GetDb(ctx).Model(userRole).Where("user_id=? and role_id=?",
		userId,
		d.GetDb(ctx).Model(&model.Role{}).Where("name=?", roleName).Select("id").SubQuery(),
	).Find(userRole).Error
	if err != nil && gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	if err != nil {
		return nil, errs.NewInternal().WithError(err)
	}
	return userRole, nil
}

func (d *UserDao) CreateUserRole(ctx context.Context, userId string, roleName string) error {
	err := d.GetDb(ctx).Model(&model.UserRole{}).Create(&model.UserRole{
		Model: common_model.Model{
			Id:         id_util.NextSnowflake(),
			CreateTime: time.Now(),
		},
		RoleName: roleName,
		UserId:   userId,
	}).Error
	if err != nil {
		return errs.NewInternal().WithError(err)
	}
	return nil
}

func (d *UserDao) DeleteUserRole(ctx context.Context, userId string, roleName string) error {
	err := d.GetDb(ctx).Model(&model.UserRole{}).Where(&model.UserRole{
		RoleName: roleName,
		UserId:   userId,
	}).Delete(&model.UserRole{}).Error
	if err != nil {
		return errs.NewInternal().WithError(err)
	}
	return nil
}

func (d *UserDao) GetUserRolesByUserId(ctx context.Context, userId string) ([]*model.UserRole, error) {
	var list []*model.UserRole
	err := d.GetDb(ctx).Model(&model.UserRole{}).Where("user_id=?", userId).Find(&list).Error
	if err != nil {
		return nil, errs.NewInternal().WithError(err)
	}
	return list, nil
}

func (d *UserDao) GetAllUserRole(ctx context.Context) ([]*model.UserRole, error) {
	var list []*model.UserRole
	err := d.GetDb(ctx).Model(&model.UserRole{}).Select("role_name,user_id").Find(&list).Error
	if err != nil {
		return nil, errs.NewInternal().WithError(err)
	}
	return list, nil
}

func (d *UserDao) GetAllRolePermission(ctx context.Context) ([]*model.RolePermission, error) {
	var list []*model.RolePermission
	err := d.GetDb(ctx).Model(&model.RolePermission{}).Select("role_name,permission_name").Find(&list).Error
	if err != nil {
		return nil, errs.NewInternal().WithError(err)
	}
	return list, nil
}

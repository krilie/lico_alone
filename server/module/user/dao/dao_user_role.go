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

type IUserRole interface {
	HasUserRoleByName(userId string, roleName string) (bool, error)
	GetUserRoleByName(userId string, roleName string) (*model.UserRole, error)
	GetUserRolesByUserId(ctx context.Context, userId string) ([]*model.UserRole, error)
	CreateUserRole(userId string, roleId string) error
	DeleteUserRole(userId string, roleId string) error
	GetAllUserRole(ctx context.Context) ([]*model.UserRole, error)
	GetAllRolePermission(ctx context.Context) ([]*model.RolePermission, error)
}

func (d *Dao) HasUserRoleByName(userId string, roleName string) (bool, error) {
	userRoleCount := 0
	err := d.Db.Model(&model.UserRole{}).Where("user_id=? and role_id=?",
		userId,
		d.Db.Model(&model.Role{}).Where("name=?", roleName).Select("id").SubQuery(),
	).Count(&userRoleCount).Error
	if err != nil {
		return false, errs.NewInternal().WithError(err)
	}
	if userRoleCount == 0 {
		return false, nil
	}
	return true, nil
}

func (d *Dao) GetUserRoleByName(userId string, roleName string) (*model.UserRole, error) {
	userRole := new(model.UserRole)
	err := d.Db.Model(userRole).Where("user_id=? and role_id=?",
		userId,
		d.Db.Model(&model.Role{}).Where("name=?", roleName).Select("id").SubQuery(),
	).Find(userRole).Error
	if err != nil && gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	if err != nil {
		return nil, errs.NewInternal().WithError(err)
	}
	return userRole, nil
}

func (d *Dao) CreateUserRole(userId string, roleName string) error {
	err := d.Db.Model(&model.UserRole{}).Create(&model.UserRole{
		Model: model.Model{
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

func (d *Dao) DeleteUserRole(userId string, roleName string) error {
	err := d.Db.Model(&model.UserRole{}).Where(&model.UserRole{
		RoleName: roleName,
		UserId:   userId,
	}).Delete(&model.UserRole{}).Error
	if err != nil {
		return errs.NewInternal().WithError(err)
	}
	return nil
}

func (d *Dao) GetUserRolesByUserId(ctx context.Context, userId string) ([]*model.UserRole, error) {
	var list []*model.UserRole
	err := d.Db.Model(&model.UserRole{}).Where("user_id=?", userId).Find(&list).Error
	if err != nil {
		return nil, errs.NewErrDbQuery().WithError(err)
	}
	return list, nil
}

func (d *Dao) GetAllUserRole(ctx context.Context) ([]*model.UserRole, error) {
	var list []*model.UserRole
	err := d.Db.Model(&model.UserRole{}).Select("role_name,user_id").Find(&list).Error
	if err != nil {
		return nil, errs.NewErrDbQuery().WithError(err)
	}
	return list, nil
}

func (d *Dao) GetAllRolePermission(ctx context.Context) ([]*model.RolePermission, error) {
	var list []*model.RolePermission
	err := d.Db.Model(&model.RolePermission{}).Select("role_name,permission_name").Find(&list).Error
	if err != nil {
		return nil, errs.NewErrDbQuery().WithError(err)
	}
	return list, nil
}

package domain

import (
	"context"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/module/user/dao"
	"github.com/krilie/lico_alone/module/user/model"
)

type Role struct {
	dao  *dao.Dao
	role *model.Role
}

func NewRole(dao *dao.Dao, roleName string) (*Role, error) {
	role, err := dao.GetRoleByName(roleName)
	if err != nil {
		return nil, err
	}
	return &Role{
		dao:  dao,
		role: role,
	}, nil
}

func (a *Role) HasPermission(ctx context.Context, permName string) (bool, error) {
	has, err := a.dao.HasRolePerm(ctx, a.role.Name, permName)
	if err != nil {
		return false, err
	}
	if has {
		return false, nil
	}
	return true, nil
}

// AddPermission 调用时确保权限存在
func (a *Role) AddPermission(ctx context.Context, permName string) error {
	b, err := a.dao.HasRolePerm(ctx, a.role.Name, permName)
	if err != nil {
		return err
	}
	if b {
		return errs.NewBadRequest().WithMsg("已有此权限")
	}
	err = a.dao.CreateRolePerm(ctx, a.role.Name, permName)
	return err
}
func (a *Role) RemovePermission(ctx context.Context, permName string) error {
	b, err := a.dao.HasRolePerm(ctx, a.role.Name, permName)
	if err != nil {
		return err
	}
	if !b {
		return errs.NewBadRequest().WithMsg("没有此权限")
	}
	err = a.dao.DeleteRolePerm(ctx, a.role.Name, permName)
	return err
}

func (a *Role) Permissions(ctx context.Context) ([]*model.RolePermission, error) {
	return a.dao.GetRolePermsByRoleName(ctx, a.role.Name)
}

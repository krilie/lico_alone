package service

import (
	"context"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/mikespook/gorbac"
)

func (a *Service) AuthCacheLoadAll(ctx context.Context) {
	// 加载所有用户
	userIds, err := a.Dao.GetAllValidUserId(ctx)
	if err != nil {
		panic(err)
	}
	for _, v := range userIds {
		err := a.AuthRBAC.AddRole(v)
		if err != nil && err != gorbac.ErrRoleExist {
			panic(err)
		}
	}
	// 加载所有角色，层级结构
	parents := []string{""}
	for {
		if len(parents) == 0 {
			break
		}
		roles, err := a.Dao.GetAllRole(ctx, parents...)
		if err != nil {
			panic(err)
		}
		parents = parents[0:0]
		for _, v := range roles {
			parents = append(parents, v.Name)
			err := a.AuthRBAC.AddRole(v.Name)
			if err != nil && err != gorbac.ErrRoleExist {
				panic(err)
			}
			if v.ParentName != "" {
				err := a.AuthRBAC.SetParent(v.ParentName, v.Name)
				if err != nil && err != gorbac.ErrRoleNotExist {
					panic(err)
				}
			}
		}
	}
	// 加载所有用户角色
	userRoles, err := a.Dao.GetAllUserRole(ctx)
	if err != nil {
		panic(err)
	}
	for _, v := range userRoles {
		err = a.AuthRBAC.SetParent(v.UserId, v.RoleName)
		if err != nil && err != gorbac.ErrRoleNotExist {
			panic(err)
		}
	}
	// 加载所有角色权限
	permissions, err := a.Dao.GetAllRolePermission(ctx)
	if err != nil {
		panic(err)
	}
	for _, v := range permissions {
		role, err := a.AuthRBAC.GetRole(v.RoleName)
		if err != nil {
			panic(err)
		}
		if role == nil {
			panic(errs.NewBadRequest().WithMsg("no role find"))
			return
		}
		stdRole, ok := role.(*gorbac.StdRole)
		if ok {
			_ = stdRole.Assign(gorbac.NewStdPermission(v.PermissionName))
		}
	}
}

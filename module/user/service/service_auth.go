package service

import (
	"context"
	"github.com/mikespook/gorbac"
)

func (s *Service) MustAuthCacheLoadAll(ctx context.Context) {
	if err := s.AuthCacheLoadAll(ctx); err != nil {
		panic(err)
	}
}

func (s *Service) AuthCacheLoadAll(ctx context.Context) error {
	// 加载所有用户->加载所有角色->加载所有权限
	userIds, err := s.Dao.GetAllValidUserId(ctx)
	if err != nil {
		panic(err)
	}
	for _, v := range userIds {
		// 加载所有角色
		authRole := gorbac.NewStdRole(v)
		_ = s.AuthRBAC.Add(authRole)
		roles, err := s.Dao.GetUserRolesByUserId(ctx, v)
		if err != nil {
			return err
		}
		for _, userRole := range roles {
			role := gorbac.NewStdRole(userRole.RoleName)
			perms, err := s.Dao.GetRolePermsByRoleName(ctx, userRole.RoleName)
			if err != nil {
				return err
			}
			for _, v := range perms {
				if err := role.Assign(gorbac.NewStdPermission(v.PermissionName)); err != nil {
					return err
				}
			}
			_ = s.AuthRBAC.Add(role)
			if err := s.AuthRBAC.SetParent(role.ID(), authRole.ID()); err != nil {
				return err
			}
		}
	}
	return nil
}

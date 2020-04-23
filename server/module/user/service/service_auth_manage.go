package service

import (
	"context"
	"github.com/krilie/lico_alone/common/utils/str_util"
	"github.com/krilie/lico_alone/component/clog"
	"github.com/mikespook/gorbac"
)

func (s *Service) MustAuthCacheLoadAll(ctx context.Context) {
	if err := s.AuthCacheLoadAll(ctx); err != nil {
		panic(err)
	}
}

func (s *Service) AuthCacheLoadAll(ctx context.Context) error {
	log := clog.NewLog(ctx, "module/user/service/service_auth.go:15", "AuthCacheLoadAll")
	// 加载所有用户->加载所有角色->加载所有权限
	userIds, err := s.Dao.GetAllValidUserId(ctx)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	for _, v := range userIds {
		log.Infof("going add role to user: %v", str_util.ToJson(v))
		// 加载所有角色
		authRole := gorbac.NewStdRole(v)
		_ = s.AuthRBAC.Add(authRole)
		roles, err := s.Dao.GetUserRolesByUserId(ctx, v)
		if err != nil {
			log.Error(err)
			return err
		}
		for _, userRole := range roles {
			role := gorbac.NewStdRole(userRole.RoleName)
			// 加载所有角色 所有权限
			perms, err := s.Dao.GetRolePermsByRoleName(ctx, userRole.RoleName)
			if err != nil {
				log.Error(err)
				return err
			}
			for _, v := range perms {
				log.Infof("add perm:%v for role:%v", v.PermissionName, role.ID())
				if err := role.Assign(gorbac.NewStdPermission(v.PermissionName)); err != nil {
					log.Error(err)
					return err
				}
			}
			// 添加此角色
			_ = s.AuthRBAC.Add(role)
			log.Infof("add role:%v for user:%v", role.ID(), authRole.ID())
			if err := s.AuthRBAC.SetParent(role.ID(), authRole.ID()); err != nil {
				log.Error(err)
				return err
			}
		}
	}
	log.Infoln("auth role permission init done.")
	return nil
}

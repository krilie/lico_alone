package service

import (
	"context"
	"github.com/krilie/lico_alone/common/utils/jwt"
)

// 权限接口 动态加载

func (s *UserModule) HasUser(ctx context.Context, userId string) (bool, error) {
	user, err := s.Dao.GetUserMasterById(ctx, userId)
	if err != nil {
		return false, err
	}
	if user == nil {
		return false, nil
	}
	return true, nil
}

func (s *UserModule) HasPermission(ctx context.Context, userId, method, path string) (bool, error) {
	methodPath, err := s.Dao.GetPermByMethodPath(ctx, method, path)
	if err != nil {
		return false, err
	}
	if methodPath == nil {
		return false, nil
	}
	return true, nil
}

func (s *UserModule) HasRole(ctx context.Context, userId, roleName string) (bool, error) {
	userRole, err := s.Dao.GetUserRoleByName(ctx, userId, roleName)
	if err != nil {
		return false, err
	}
	if userRole == nil {
		return false, nil
	}
	return true, nil
}

func (s *UserModule) CheckJwtToken(tokenStr string) (userClaims jwt.UserClaims, err error) {
	return jwt.CheckJwtToken(s.jwtSecret, tokenStr)
}

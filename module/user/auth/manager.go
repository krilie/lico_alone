package auth

import (
	"github.com/krilie/lico_alone/common/clog"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/module/user/model"
	"time"
)

type UserManager interface {
	NewClientAccToken(ctx context.Context, userId, keyDescription string, Exp time.Time) (key *model.ClientUserAccessToken, err error)
	NewPermission(ctx context.Context, pName string, pDescription string) (p *model.Permission, err error)
	AddPermissionToRole(ctx context.Context, roleId string, permissionId string) error
	NewRole(ctx context.Context, roleName string, roleDescription string) (role *model.Role, err error)
	AddRoleToUser(ctx context.Context, roleId string, userId string) error
}

type UserManage struct{}

var User UserAuth

var log = clog.NewLog(context.NewContext(), "alone.module.userbase.auth", "init")

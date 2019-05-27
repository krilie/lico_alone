package auth

import (
	"github.com/krilie/lico_alone/common/context_util"
	"github.com/krilie/lico_alone/module/userbase/model"
	"time"
)

type UserManager interface {
	NewClientAccToken(ctx *context_util.Context, userId, keyDescription string, Exp time.Time) (key *model.ClientUserAccessToken, err error)
	NewPermission(ctx *context_util.Context, pName string, pDescription string) (p *model.Permission, err error)
	AddPermissionToRole(ctx *context_util.Context, roleId string, permissionId string) error
	NewRole(ctx *context_util.Context, roleName string, roleDescription string) (role *model.Role, err error)
	AddRoleToUser(ctx *context_util.Context, roleId string, userId string) error
}

type UserManage struct{}

var User UserAuth

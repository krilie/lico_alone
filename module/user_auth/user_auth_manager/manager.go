package user_auth_manager

import (
	"github.com/krilie/lico_alone/common/context_util"
	"github.com/krilie/lico_alone/module/user_auth/model"
	"time"
)

type Manager interface {
	ManagerClientUserNewAccToken(ctx *context_util.Context, userId, keyDescription string, Exp time.Time) (key *model.ClientUserAccessToken, err error)
	ManagerPermissionNewPermission(ctx *context_util.Context, pName string, pDescription string) (p *model.Permission, err error)
	ManagerRoleAddPermission(ctx *context_util.Context, roleId string, permissionId string) error
	ManagerRoleNewRole(ctx *context_util.Context, roleName string, roleDescription string) (role *model.Role, err error)
	ManagerUserAddRole(ctx *context_util.Context, roleId string, userId string) error
}

type Manage struct{}

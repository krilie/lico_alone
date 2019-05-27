package manager

import (
	"github.com/krilie/lico_alone/common/context_util"
	"github.com/krilie/lico_alone/module/userbase/model"
	"time"
)

func (AppManager) NewClientAccToken(ctx *context_util.Context, userId, keyDescription string, Exp time.Time) (key *model.ClientUserAccessToken, err error) {
	panic("implement me")
}

func (AppManager) NewPermission(ctx *context_util.Context, pName string, pDescription string) (p *model.Permission, err error) {
	panic("implement me")
}

func (AppManager) AddPermissionToRole(ctx *context_util.Context, roleId string, permissionId string) error {
	panic("implement me")
}

func (AppManager) NewRole(ctx *context_util.Context, roleName string, roleDescription string) (role *model.Role, err error) {
	panic("implement me")
}

func (AppManager) AddRoleToUser(ctx *context_util.Context, roleId string, userId string) error {
	panic("implement me")
}

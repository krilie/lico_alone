package auth

import (
	mapset "github.com/deckarep/golang-set"
	"github.com/krilie/lico_alone/common/context_util"
	"github.com/krilie/lico_alone/module/userbase/model"
)

type UserAuther interface {
	GetClientAccToken(ctx *context_util.Context, appUserId string) (list []model.ClientUserAccessToken, err error)
	HasClientAccToken(ctx *context_util.Context, userId, accTokenStr string) (token *model.ClientUserAccessToken, err error)
	HasPermission(ctx *context_util.Context, userId, permissionName string) (bool, error)
	HasRole(ctx *context_util.Context, userId, roleName string) (bool, error)
	GetPermissions(ctx *context_util.Context, userId string) (set mapset.Set, err error)
	GetRoles(ctx *context_util.Context, userId string) (roles mapset.Set, err error)
}
type UserAuth struct{}

var User UserAuth

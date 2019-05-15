package user_auth

import (
	mapset "github.com/deckarep/golang-set"
	"github.com/krilie/lico_alone/common/context_util"
	"github.com/krilie/lico_alone/module/user_auth/model"
)

type UserAuther interface {
	UserAuthClientAccToken(ctx *context_util.Context, appUserId string) (list []model.ClientUserAccessToken, err error)
	UserAuthClientHasAccToken(ctx *context_util.Context, userId, accTokenStr string) (token *model.ClientUserAccessToken, err error)
	UserAuthHasPermission(ctx *context_util.Context, userId, permissionName string) (bool, error)
	UserAuthHasRole(ctx *context_util.Context, userId, roleName string) (bool, error)
	UserAuthPermissions(ctx *context_util.Context, userId string) (set mapset.Set, err error)
	UserAuthRoles(ctx *context_util.Context, userId string) (roles mapset.Set, err error)
}
type UserAuth struct{}

var User UserAuth

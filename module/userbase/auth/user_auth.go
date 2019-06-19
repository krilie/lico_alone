package auth

import (
	mapset "github.com/deckarep/golang-set"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/module/userbase/model"
)

type UserAuther interface {
	GetClientAccToken(ctx context.Context, appUserId string) (list []model.ClientUserAccessToken, err error)
	HasClientAccToken(ctx context.Context, userId, accTokenStr string) (token *model.ClientUserAccessToken, err error)
	HasPermission(ctx context.Context, userId, permissionName string) (bool, error)
	HasRole(ctx context.Context, userId, roleName string) (bool, error)
	GetPermissions(ctx context.Context, userId string) (set mapset.Set, err error)
	GetRoles(ctx context.Context, userId string) (roles mapset.Set, err error)
}
type UserAuth struct{}

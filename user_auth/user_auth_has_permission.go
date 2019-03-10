package user_auth

import (
	"github.com/asaskevich/govalidator"
	"github.com/lico603/lico-my-site-user/common/context_util"
	"github.com/lico603/lico-my-site-user/common/errs"
	"github.com/lico603/lico-my-site-user/common/log"
	"github.com/lico603/lico-my-site-user/common/validator_util"
	"github.com/lico603/lico-my-site-user/model"
)

//判断用户是否有这个权限，联合查询
func UserAuthHasPermission(ctx *context_util.Context, userId, permissionName string) (bool, error) {
	if !(validator_util.IsIdStr(userId) && govalidator.IsAlpha(permissionName)) {
		log.Infoln("UserAuthHasPermission", "param error:", userId, permissionName)
		return false, errs.ErrParam
	}
	return model.IsPermissionExistsWithUser(model.Db, userId, permissionName)
}

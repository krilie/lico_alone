package user_auth

import (
	"github.com/asaskevich/govalidator"
	"github.com/krilie/lico_alone/common/common_struct/errs"
	"github.com/krilie/lico_alone/common/context_util"
	"github.com/krilie/lico_alone/common/log"
	"github.com/krilie/lico_alone/common/validator_util"
	"github.com/krilie/lico_alone/module/user_auth/model"
)

//判断用户是否有这个权限，联合查询
func (UserAuth) UserAuthHasPermission(ctx *context_util.Context, userId, permissionName string) (bool, error) {
	if !(validator_util.IsIdStr(userId) && govalidator.IsAlpha(permissionName)) {
		log.Infoln("UserAuthHasPermission", "param error:", userId, permissionName)
		return false, errs.ErrParam
	}
	return model.IsPermissionExistsWithUser(model.Db, userId, permissionName)
}

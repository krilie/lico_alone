package auth

import (
	"github.com/asaskevich/govalidator"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/model/errs"
	"github.com/krilie/lico_alone/common/utils/validator"
	"github.com/krilie/lico_alone/module/user/dao"
	"github.com/krilie/lico_alone/module/user/model"
)

//判断用户是否有这个权限，联合查询
func (UserAuth) HasPermission(ctx context.Context, userId, permissionName string) (bool, error) {
	if !(validator.IsIdStr(userId) && govalidator.IsAlpha(permissionName)) {
		log.Infoln("HasPermission", "param error:", userId, permissionName)
		return false, errs.ErrParam
	}
	return model.IsPermissionExistsWithUser(dao.Db, userId, permissionName)
}

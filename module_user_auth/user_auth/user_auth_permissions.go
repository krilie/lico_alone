package user_auth

import (
	"github.com/deckarep/golang-set"
	"github.com/krilie/lico_alone/common/common_struct/errs"
	"github.com/krilie/lico_alone/common/context_util"
	"github.com/krilie/lico_alone/common/log"
	"github.com/krilie/lico_alone/common/validator_util"
	"github.com/krilie/lico_alone/module_user_auth/model"
)

//获取用户的 permission 列表,连表查询
func UserAuthPermissions(ctx *context_util.Context, userId string) (set mapset.Set, err error) {
	if !validator_util.IsIdStr(userId) {
		log.Infoln("UserAuthPermissions", "user id format error:", userId)
		return nil, errs.ErrParam
	}
	return model.GetAllPermissionByUserId(model.Db, userId)
}

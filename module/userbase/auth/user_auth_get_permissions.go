package auth

import (
	"github.com/deckarep/golang-set"
	"github.com/krilie/lico_alone/common/common_struct/errs"
	"github.com/krilie/lico_alone/common/context_util"
	"github.com/krilie/lico_alone/common/log"
	"github.com/krilie/lico_alone/common/validator_util"
	"github.com/krilie/lico_alone/module/userbase/model"
)

//获取用户的 permission 列表,连表查询
func (UserAuth) GetPermissions(ctx *context_util.Context, userId string) (set mapset.Set, err error) {
	if !validator_util.IsIdStr(userId) {
		log.Infoln("GetPermissions", "user id format error:", userId)
		return nil, errs.ErrParam
	}
	return model.GetAllPermissionByUserId(model.Db, userId)
}

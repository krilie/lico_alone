package auth

import (
	"github.com/deckarep/golang-set"
	"github.com/krilie/lico_alone/common/comstruct/errs"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/utils/validator"
	"github.com/krilie/lico_alone/module/userbase/model"
)

//获取用户的 permission 列表,连表查询
func (UserAuth) GetPermissions(ctx *context.Context, userId string) (set mapset.Set, err error) {
	if !validator.IsIdStr(userId) {
		log.Infoln("GetPermissions", "user id format error:", userId)
		return nil, errs.ErrParam
	}
	return model.GetAllPermissionByUserId(model.Db, userId)
}

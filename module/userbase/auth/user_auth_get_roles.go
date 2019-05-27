package auth

import (
	"github.com/deckarep/golang-set"
	"github.com/krilie/lico_alone/common/comstruct/errs"
	"github.com/krilie/lico_alone/common/context_util"
	"github.com/krilie/lico_alone/common/validator_util"
	"github.com/krilie/lico_alone/module/userbase/model"
)

//获取用户的所有角色
func (UserAuth) GetRoles(ctx *context_util.Context, userId string) (roles mapset.Set, err error) {
	if !validator_util.IsIdStr(userId) {
		return nil, errs.ErrParam
	}
	return model.GetAllRolesByUserId(model.Db, userId)
}

package user_auth

import (
	"github.com/deckarep/golang-set"
	"github.com/lico603/lico_user/common/common_struct/errs"
	"github.com/lico603/lico_user/common/context_util"
	"github.com/lico603/lico_user/common/validator_util"
	"github.com/lico603/lico_user/model"
)

//获取用户的所有角色
func UserAuthRoles(ctx *context_util.Context, userId string) (roles mapset.Set, err error) {
	if !validator_util.IsIdStr(userId) {
		return nil, errs.ErrParam
	}
	return model.GetAllRolesByUserId(model.Db, userId)
}

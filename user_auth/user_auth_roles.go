package user_auth

import (
	"github.com/deckarep/golang-set"
	"github.com/lico603/lico-my-site-user/common/errs"
	"github.com/lico603/lico-my-site-user/common/validator_util"
	"github.com/lico603/lico-my-site-user/model"
)

//获取用户的所有角色
func UserAuthRoles(userId string) (roles mapset.Set, err error) {
	if !validator_util.IsIdString(userId) {
		return nil, errs.ErrParam
	}
	return model.GetAllRolesByUserId(model.Db, userId)
}

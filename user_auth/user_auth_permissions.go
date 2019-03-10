package user_auth

import (
	"github.com/deckarep/golang-set"
	"github.com/lico603/lico-my-site-user/common/context_util"
	"github.com/lico603/lico-my-site-user/common/errs"
	"github.com/lico603/lico-my-site-user/common/log"
	"github.com/lico603/lico-my-site-user/common/validator_util"
	"github.com/lico603/lico-my-site-user/model"
)

//获取用户的 permission 列表,连表查询
func UserAuthPermissions(ctx *context_util.Context, userId string) (set mapset.Set, err error) {
	if !validator_util.IsIdString(userId) {
		log.Infoln("UserAuthPermissions", "user id format error:", userId)
		return nil, errs.ErrParam
	}
	return model.GetAllPermissionByUserId(model.Db, userId)
}

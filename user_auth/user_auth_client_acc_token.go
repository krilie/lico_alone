package user_auth

import (
	"github.com/lico603/lico-my-site-user/common/common_struct/errs"
	"github.com/lico603/lico-my-site-user/common/context_util"
	"github.com/lico603/lico-my-site-user/common/log"
	"github.com/lico603/lico-my-site-user/common/validator_util"
	"github.com/lico603/lico-my-site-user/model"
)

//取到app角色用户的所有keys
func UserAuthClientAccToken(ctx *context_util.Context, appUserId string) (list []model.ClientUserAccessToken, err error) {
	//校验参数
	if !validator_util.IsIdStr(appUserId) {
		log.Infoln("UserAuthAppKeys", "err param:", appUserId)
		return nil, errs.ErrParam
	}
	//根据用户id 查询到 该用户的下的所有key 这个用户可能是android app//ios app//doc service 等
	list = make([]model.ClientUserAccessToken, 4)
	err = model.Db.Where("user_id=?", appUserId).Find(list).Error
	if err != nil {
		return nil, err
	} else {
		return list, nil
	}
}

package user_auth

import (
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
	"github.com/lico603/lico-my-site-user/common/common_struct/errs"
	"github.com/lico603/lico-my-site-user/common/context_util"
	"github.com/lico603/lico-my-site-user/common/log"
	"github.com/lico603/lico-my-site-user/common/validator_util"
	"github.com/lico603/lico-my-site-user/model"
)

//查看client是否有acc token
func UserAuthClientHasAccToken(ctx *context_util.Context, userId, accTokenStr string) (token *model.ClientUserAccessToken, err error) {
	if len(accTokenStr) == 0 || !govalidator.IsASCII(accTokenStr) || !validator_util.IsIdStr(userId) {
		log.Infoln("userAuthClientHasAccToken", "acc token error:", accTokenStr, userId)
		return nil, errs.ErrParam
	}
	token = new(model.ClientUserAccessToken)
	err = model.Db.Where(&model.ClientUserAccessToken{UserId: userId, Token: accTokenStr}).Find(token).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	} else {
		return token, nil
	}
}

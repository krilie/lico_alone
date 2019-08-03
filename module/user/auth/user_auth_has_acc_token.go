package auth

import (
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/model/errs"
	"github.com/krilie/lico_alone/common/utils/validator"
	"github.com/krilie/lico_alone/module/user/dao"
	"github.com/krilie/lico_alone/module/user/model"
)

//查看client是否有acc token
func (UserAuth) HasClientAccToken(ctx context.Context, userId, accTokenStr string) (token *model.ClientUserAccessToken, err error) {
	if len(accTokenStr) == 0 || !govalidator.IsASCII(accTokenStr) || !validator.IsIdStr(userId) {
		log.Infoln("userAuthClientHasAccToken", "acc token error:", accTokenStr, userId)
		return nil, errs.ErrParam
	}
	token = new(model.ClientUserAccessToken)
	err = dao.Db.Where(&model.ClientUserAccessToken{UserId: userId, Token: accTokenStr}).Find(token).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	} else {
		return token, nil
	}
}

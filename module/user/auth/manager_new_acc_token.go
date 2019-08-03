package auth

import (
	"context"
	"github.com/krilie/lico_alone/common/comstruct/errs"
	"github.com/krilie/lico_alone/common/utils/random"
	"github.com/krilie/lico_alone/common/utils/validator"
	"github.com/krilie/lico_alone/module/user/model"
	"time"
)

// 给app用户添加新的key 访问key
// admin 和cleint，如果没有admin而只有cleint权限，则检查登录者与userId是否一致
func (UserManage) NewClientAccToken(ctx context.Context, loginUserId, userId, keyDescription string, Exp time.Time) (key *model.ClientUserAccessToken, err error) {
	//参数检查
	if !validator.IsIdStr(userId) || len(keyDescription) == 0 {
		log.Infoln("", "param error:", userId, keyDescription)
		return nil, errs.ErrParam
	}
	//添加一个key
	key = new(model.ClientUserAccessToken)
	key.CreateTime = time.Now()
	key.UserId = userId
	key.CreateBy = loginUserId
	key.Description = keyDescription
	key.IsValid = true
	key.ExpirationTime = Exp
	key.Token = random.GetAToken()
	err = model.Db.Create(key).Error
	if err != nil {
		return nil, err
	} else {
		return key, nil
	}
}

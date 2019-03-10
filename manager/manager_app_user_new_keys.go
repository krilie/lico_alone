package manager

import (
	"github.com/lico603/lico-my-site-user/common/context_util"
	"github.com/lico603/lico-my-site-user/common/errs"
	"github.com/lico603/lico-my-site-user/common/log"
	"github.com/lico603/lico-my-site-user/common/random_token"
	"github.com/lico603/lico-my-site-user/common/validator_util"
	"github.com/lico603/lico-my-site-user/model"
	"time"
)

// 给app用户添加新的key 访问key
func ManagerAppUserNewKeys(ctx *context_util.Context, userId, keyDescription string, Exp time.Time) (key *model.AppUserAccessToken, err error) {
	//参数检查
	if !validator_util.IsIdString(userId) || len(keyDescription) == 0 {
		log.Infoln("", "param error:", userId, keyDescription)
		return nil, errs.ErrParam
	}
	//添加一个key
	key = new(model.AppUserAccessToken)
	key.CreateTime = time.Now()
	key.UserId = userId
	key.CreateBy = ctx.GetUserIdOrDefault(userId)
	key.Description = keyDescription
	key.IsValid = true
	key.Token = random_token.GetAToken()
	err = model.Db.Create(key).Error
	if err != nil {
		return nil, err
	} else {
		return key, nil
	}
}

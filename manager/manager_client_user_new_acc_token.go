package manager

import (
	"github.com/krilie/lico_alone/common/common_struct/errs"
	"github.com/krilie/lico_alone/common/context_util"
	"github.com/krilie/lico_alone/common/log"
	"github.com/krilie/lico_alone/common/random_token"
	"github.com/krilie/lico_alone/common/validator_util"
	"github.com/krilie/lico_alone/model"
	"github.com/krilie/lico_alone/user_auth"
	"time"
)

// 给app用户添加新的key 访问key
// admin 和cleint，如果没有admin而只有cleint权限，则检查登录者与userId是否一致
func ManagerClientUserNewAccToken(ctx *context_util.Context, userId, keyDescription string, Exp time.Time) (key *model.ClientUserAccessToken, err error) {
	//参数检查
	if !validator_util.IsIdStr(userId) || len(keyDescription) == 0 {
		log.Infoln("", "param error:", userId, keyDescription)
		return nil, errs.ErrParam
	}
	//判断是不是有client权限
	loginUserId := ctx.GetUserIdOrEmpty()
	if loginUserId == "" {
		return nil, errs.UnAuthorized
	}
	//判断是否没有admin权限而有client权限
	hasRole, err := user_auth.UserAuthHasRole(ctx, loginUserId, model.RoleAdmin)
	if err != nil {
		return nil, err
	}
	if !hasRole {
		if hasRoleClient, err := user_auth.UserAuthHasRole(ctx, loginUserId, model.RoleClient); err != nil {
			return nil, err
		} else if hasRoleClient {
			//没有admin只有client权限,检查登录者是否与target一致
			if ctx.GetUserIdOrEmpty() != loginUserId {
				return nil, errs.ErrNoPermission.NewWithMsg("只能给自已添加cleint acc key")
			}
		} else {
			return nil, errs.UnAuthorized
		}
	}
	//添加一个key
	key = new(model.ClientUserAccessToken)
	key.CreateTime = time.Now()
	key.UserId = userId
	key.CreateBy = ctx.GetUserIdOrDefault(userId)
	key.Description = keyDescription
	key.IsValid = true
	key.ExpirationTime = Exp
	key.Token = random_token.GetAToken()
	err = model.Db.Create(key).Error
	if err != nil {
		return nil, err
	} else {
		return key, nil
	}
}

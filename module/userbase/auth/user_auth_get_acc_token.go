package auth

import (
	"github.com/krilie/lico_alone/common/comstruct/errs"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/log"
	"github.com/krilie/lico_alone/common/validator"
	"github.com/krilie/lico_alone/module/userbase/model"
)

//取到app角色用户的所有keys
func (ua UserAuth) GetClientAccToken(ctx *context.Context, appUserId string) (list []model.ClientUserAccessToken, err error) {
	//校验参数
	if !validator.IsIdStr(appUserId) {
		log.Infoln("UserAuthAppKeys", "err param:", appUserId)
		return nil, errs.ErrParam
	}
	//判断是不是有client权限
	loginUserId := ctx.GetUserIdOrEmpty()
	if loginUserId == "" {
		return nil, errs.UnAuthorized
	}
	//判断是否没有admin权限而有client权限
	hasRole, err := ua.HasRole(ctx, loginUserId, model.RoleAdmin)
	if err != nil {
		return nil, err
	}
	if !hasRole {
		if hasRoleClient, err := ua.HasRole(ctx, loginUserId, model.RoleClient); err != nil {
			return nil, err
		} else if hasRoleClient {
			//没有admin只有client权限,检查登录者是否与target一致
			if ctx.GetUserIdOrEmpty() != loginUserId {
				return nil, errs.ErrNoPermission.NewWithMsg("只能查询自已的acc_token")
			}
		} else {
			return nil, errs.UnAuthorized
		}
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

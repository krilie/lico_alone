package user_base

import (
	"github.com/lico603/lico-my-site-user/common/context_util"
	"github.com/lico603/lico-my-site-user/common/string_util"
	"github.com/lico603/lico-my-site-user/model"
)

//由token取得用户基本信息

func UserGetInfo(ctx *context_util.Context, userId string) map[string]string {
	//已经登录了
	var user model.User
	err := model.Db.First(&user, "id = ?", userId).Error
	if err != nil {
		return nil
	} else {
		mUser := make(map[string]string, 4)
		mUser["id"] = user.ID
		mUser["login_name"] = user.LoginName
		mUser["nick_name"] = user.NickName
		mUser["phone"] = string_util.SqlStringOrEmpty(user.Phone)
		mUser["email"] = string_util.SqlStringOrEmpty(user.Email)
		return mUser
	}
}

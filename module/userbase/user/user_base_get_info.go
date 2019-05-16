package user

import (
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/common_struct/errs"
	"github.com/krilie/lico_alone/common/context_util"
	"github.com/krilie/lico_alone/common/string_util"
	"github.com/krilie/lico_alone/module/userbase/model"
)

//由token取得用户基本信息

func (User) UserBaseGetInfo(ctx *context_util.Context, userId string) (map[string]string, error) {
	//已经登录了
	var user model.User
	err := model.Db.First(&user, "id = ?", userId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.ErrNotFound.NewWithMsg("no this user:" + userId)
		} else {
			return nil, err
		}
	} else {
		mUser := make(map[string]string, 4)
		mUser["id"] = user.ID
		mUser["login_name"] = user.LoginName
		mUser["nick_name"] = user.NickName
		mUser["phone"] = string_util.SqlStringOrEmpty(user.Phone)
		mUser["email"] = string_util.SqlStringOrEmpty(user.Email)
		mUser["picture"] = string_util.SqlStringOrEmpty(user.Picture)
		return mUser, nil
	}
}

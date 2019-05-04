package user_base

import (
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/common_struct/errs"
	"github.com/krilie/lico_alone/common/context_util"
	"github.com/krilie/lico_alone/common/jwt"
	"github.com/krilie/lico_alone/common/log"
	"github.com/krilie/lico_alone/common/pswd_md5"
	"github.com/krilie/lico_alone/common/string_util"
	"github.com/krilie/lico_alone/common/uuid_util"
	"github.com/krilie/lico_alone/common/validator_util"
	"github.com/krilie/lico_alone/module/user_auth/model"
	"time"
)

// 用户登录,到这里说明参数有可能还是不正确的。检查参数,放到上层
func UserLogin(ctx *context_util.Context, loginName, password string) (jwtString string, err error) {
	//检查密码与用户名
	if !(validator_util.IsLoginName(loginName) && validator_util.IsPassword(password)) {
		log.Infoln("user loginName or password format error.")
		return "", errs.ErrParam
	}
	//取到用户的值
	user := new(model.User)
	err = model.Db.First(user, "login_name = ?", loginName).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		log.Errorln(err)
		return "", errs.ErrNoSuchUser
	} else if err != nil {
		log.Errorln("internal error,", err)
		return "", err
	}

	if pswd_md5.IsPasswordOk(password, user.Password, user.Salt) {
		var userClaims jwt.UserClaims
		userClaims.NickName = user.NickName
		userClaims.LoginName = user.LoginName
		// 把所有角色查出放在这里
		roles, err := model.GetAllRolesByUserId(model.Db, user.ID)
		if err != nil {
			log.Error("get roles err:", err)
			return "", err
		}
		userClaims.UserRoles = string_util.JoinWith(roles, ",")
		userClaims.AppId = ctx.GetAppIdOrEmpty()
		userClaims.Iss = "sys-user-module"
		userClaims.UserId = user.ID
		userClaims.Jti = uuid_util.GetUuid()
		userClaims.Iat = time.Now().Unix()
		userClaims.Picture = string_util.SqlStringOrEmpty(user.Picture)
		//userClaims.Exp = time.Now().Add(time.Hour).Unix()
		userClaims.Exp = time.Now().Add(time.Duration(jwtExpDuration) * time.Second).Unix()
		jwtToken, err := jwt.GetNewJwtToken(&userClaims)
		if err != nil {
			log.Error("jwt err:", err)
			return "", err //未知错误
		}
		return jwtToken, nil
	} else {
		return "", errs.ErrNameOrPassword
	}
}

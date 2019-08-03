package info

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/comstruct/errs"
	"github.com/krilie/lico_alone/common/jwt"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/common/utils/pswd_util"
	"github.com/krilie/lico_alone/common/utils/str_util"
	"github.com/krilie/lico_alone/common/utils/validator"
	"github.com/krilie/lico_alone/module/user/model"
	"time"
)

// 用户登录,到这里说明参数有可能还是不正确的。检查参数,放到上层
func (User) Login(ctx context.Context, clientId, loginName, password string) (jwtString string, err error) {
	//检查密码与用户名
	if !(validator.IsLoginName(loginName) && validator.IsPassword(password)) {
		log.Infoln("user loginName or password format error.")
		return "", errs.ErrParam.NewWithMsg("login name or password format error")
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

	if pswd_util.IsPasswordOk(password, user.Password, user.Salt) {
		var userClaims jwt.UserClaims
		userClaims.NickName = user.NickName
		userClaims.LoginName = user.LoginName
		// 把所有角色查出放在这里
		roles, err := model.GetAllRolesByUserId(model.Db, user.ID)
		if err != nil {
			log.Error("get roles err:", err)
			return "", err
		}
		userClaims.UserRoles = str_util.JoinWith(roles, ",")
		userClaims.ClientId = clientId
		userClaims.Iss = "sys-user-module"
		userClaims.UserId = user.ID
		userClaims.Jti = id_util.GetUuid()
		userClaims.Iat = time.Now().Unix()
		userClaims.Picture = str_util.SqlStringOrEmpty(user.Picture)
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

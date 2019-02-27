package user

import (
	"github.com/lico603/lico-my-site-user/common/pswd_md5"
	"github.com/lico603/lico-my-site-user/model"
)

// 用户登录,到这里说明参数有可能还是不正确的。检查参数,放到上层
func UserLogin(name, pswd string) (jwtString string, err error) {
	//取到用户的值
	user := new(model.User)
	model.Db.First(user, "name = ?", name)
	if user == nil {
		return "", ErrNoSuchUser
	}
	if user.Password == pswd_md5.GetMd5Password(pswd, user.Salt) {
		//生成jwt 并返回
		//var userClaims jwt.UserClaims
		// 角色  app admin normal
		// TODO:userClaims.UserType =

		return
	} else {
		return "", ErrNameOrPassword
	}
}

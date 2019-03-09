package user_base

import (
	"database/sql"
	"github.com/lico603/lico-my-site-user/common/context_util"
	"github.com/lico603/lico-my-site-user/common/log"
	"github.com/lico603/lico-my-site-user/common/pswd_md5"
	"github.com/lico603/lico-my-site-user/common/uuid_util"
	"github.com/lico603/lico-my-site-user/common/validator_util"
	"github.com/lico603/lico-my-site-user/model"
	"time"
)

//用户注册，注册，normal用户注册
func UserRegister(ctx *context_util.Context, loginName string, password string) error {
	//检查密码与用户名
	if !(validator_util.IsLoginName(loginName) && validator_util.IsPassword(password)) {
		log.Infoln("user loginName or password format error.")
		return ErrParam
	}
	//插入用户数据
	var user model.User
	user.ID = uuid_util.GetUuid()
	user.LoginName = loginName
	user.NickName = loginName
	user.Salt = pswd_md5.GetSalt(5)
	user.Password = pswd_md5.GetMd5Password(password, user.Salt)
	user.Version = 0
	user.CreateTime = time.Now()
	user.Phone = sql.NullString{Valid: false}
	user.Email = sql.NullString{Valid: false}

	err := model.Db.Create(&user).Error
	if err != nil {
		log.Infoln("database error:", err)
		return err
	} else {
		return nil
	}
}

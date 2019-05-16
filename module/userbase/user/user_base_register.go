package user

import (
	"database/sql"
	"github.com/krilie/lico_alone/common/common_struct/errs"
	"github.com/krilie/lico_alone/common/context_util"
	"github.com/krilie/lico_alone/common/log"
	"github.com/krilie/lico_alone/common/pswd_md5"
	"github.com/krilie/lico_alone/common/uuid_util"
	"github.com/krilie/lico_alone/common/validator_util"
	"github.com/krilie/lico_alone/module/userbase/model"
	"time"
)

//用户注册，注册，normal用户注册
func (User) UserBaseRegister(ctx *context_util.Context, loginName string, password string) error {
	//检查密码与用户名
	if !(validator_util.IsLoginName(loginName) && validator_util.IsPassword(password)) {
		log.Infoln("user loginName or password format error.")
		return errs.ErrParam
	}
	//插入用户数据
	var user model.User
	user.ID = uuid_util.GetUuid()
	user.LoginName = loginName
	user.NickName = loginName
	user.Salt = pswd_md5.GetSalt(5)
	user.Password = pswd_md5.GetMd5Password(password, user.Salt)
	user.CreateTime = time.Now()
	user.Phone = sql.NullString{Valid: false}
	user.Email = sql.NullString{Valid: false}
	user.Picture = sql.NullString{Valid: false}

	err := model.Db.Create(&user).Error
	if err != nil {
		log.Infoln("database error:", err)
		return err
	} else {
		return nil
	}
}

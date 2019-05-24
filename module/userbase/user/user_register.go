package user

import (
	"database/sql"
	"github.com/jinzhu/gorm"
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
func (User) Register(ctx *context_util.Context, loginName string, password string) error {
	// 检查密码与用户名格式
	if !(validator_util.IsLoginName(loginName) && validator_util.IsPassword(password)) {
		log.Infoln("user loginName or password format error.")
		return errs.ErrParam.NewWithMsg("用户名不能以数字开头&密码至少8位")
	}
	// 检查密码与用户名的存在
	var user model.User
	if e := model.Db.Find(&user, "login_name = ?", loginName).Error; e == nil {
		return errs.ErrParam.NewWithMsg("此名称已被使用")
	} else if e != gorm.ErrRecordNotFound {
		return e
	}
	// 插入用户数据
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

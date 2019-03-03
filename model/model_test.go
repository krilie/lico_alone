package model

import (
	"github.com/lico603/lico-my-site-user/common/pswd_md5"
	"github.com/lico603/lico-my-site-user/common/uuid_util"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	var user User
	user.ID = uuid_util.GetUuid()
	user.CreateTime = time.Now()
	user.Version = 0
	user.Salt = "123"
	user.Password = pswd_md5.GetMd5Password("12345678", user.Salt)
	user.Email = "me@example.com"
	user.NickName = "ii"
	user.Phone = "12323232323"
	if e := Db.Create(&user).Error; e != nil {
		t.Error(e)
	}
	Db.Begin()
}

func TestFunc(t *testing.T) {

}

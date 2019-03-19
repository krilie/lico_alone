package model

import (
	"database/sql"
	"fmt"
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
	user.Email = sql.NullString{String: "me@example.com", Valid: true}
	user.NickName = "ii"
	user.Phone = sql.NullString{String: "12323232323", Valid: true}
	if e := Db.Create(&user).Error; e != nil {
		t.Error(e)
	}
	Db.Begin()
}

func TestFunc(t *testing.T) {

}

func TestFuncGetRoleIdByName(t *testing.T) {
	id, err := GetRoleIdByName(Db, "123")
	fmt.Println(id, err)
}

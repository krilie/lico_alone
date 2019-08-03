package model

import (
	"database/sql"
	"fmt"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/common/utils/pswd_util"
	"github.com/krilie/lico_alone/module/user/dao"
	"net/url"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	var user User
	user.ID = id_util.GetUuid()
	user.CreateTime = time.Now()
	user.Salt = "123"
	user.Password = pswd_util.GetMd5Password("12345678", user.Salt)
	user.Email = sql.NullString{String: "me@example.com", Valid: true}
	user.NickName = "ii"
	user.Phone = sql.NullString{String: "12323232323", Valid: true}
	if e := dao.Db.Create(&user).Error; e != nil {
		t.Error(e)
	}
	dao.Db.Begin()
}

func TestFunc(t *testing.T) {
	fmt.Println(url.QueryEscape("Asia/Shanghai"))
}

func TestFuncGetRoleIdByName(t *testing.T) {
	id, err := GetRoleIdByName(dao.Db, "123")
	fmt.Println(id, err)
}

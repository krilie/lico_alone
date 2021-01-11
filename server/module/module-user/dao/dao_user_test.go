package dao

import (
	"github.com/krilie/lico_alone/common/appdig"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/component"
	"github.com/krilie/lico_alone/module/module-user/model"
	"testing"
	"time"
)

var container = appdig.
	NewAppDig().
	MustProvides(component.DigComponentProviderAllForTest).
	MustProvide(NewUserDao)

func TestDao_GetAllValidUserId(t *testing.T) {
	container.MustInvoke(func(dao *UserDao) {
		err := dao.UpdateUserPassword(context.EmptyAppCtx(), "12", "dd", "")
		t.Log(err)
	})
}

func TestUserDao_DeleteUserByPhone(t *testing.T) {
	ctx := context.EmptyAppCtx()
	container.MustInvoke(func(dao *UserDao) {
		phone := id_util.NextSnowflake()
		err := dao.CreateUserMaster(ctx, model.NewUserMaster(id_util.GetUuid(), time.Now(), time.Now(), nil, time.Now(), phone, phone, "11", "", "", ""))
		CheckErr(t, err)
		err = dao.DeleteUserByPhone(ctx, phone)
		CheckErr(t, err)
	})
}

func CheckErr(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

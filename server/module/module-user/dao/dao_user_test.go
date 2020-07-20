// +build !auto_test

package dao

import (
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/module/module-user/model"
	"testing"
	"time"
)

func TestDao_GetAllValidUserId(t *testing.T) {
	dig.Container.MustInvoke(func(dao *UserDao) {
		err := dao.UpdateUserPassword(context.NewContext(), "12", "dd", "")
		t.Log(err)
	})
}

func TestUserDao_DeleteUserByPhone(t *testing.T) {
	ctx := context.NewContext()
	dig.Container.MustInvoke(func(dao *UserDao) {
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

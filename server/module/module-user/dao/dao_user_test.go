package dao

import (
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/dig"
	"testing"
)

func TestDao_GetAllValidUserId(t *testing.T) {
	dig.Container.MustInvoke(func(dao *UserDao) {
		err := dao.UpdateUserPassword(context.NewContext(), "12", "dd", "")
		t.Log(err)
	})
}

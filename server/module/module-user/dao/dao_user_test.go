package dao

import (
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/module/module-user/model"
	"testing"
	"time"
)

func TestDao_GetAllValidUserId(t *testing.T) {
	dig.MustInvoke(func(dao *UserDao) {
		err := dao.CreatePerm(context.NewContext(), &model.Permission{
			Name:        "11111",
			CreateTime:  time.Now(),
			Description: "22222",
			RefMethod:   "33333",
			RefPath:     "4444",
			Sort:        123,
		})
		t.Log(err)
	})
}

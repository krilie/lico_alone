package dao

import (
	"context"
	"github.com/krilie/lico_alone/common/appdig"
	com_model "github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/utils/jsonutil"
	"github.com/krilie/lico_alone/common/utils/random"
	"github.com/krilie/lico_alone/component"
	"github.com/krilie/lico_alone/module/module-catchword/model"
	"github.com/prometheus/common/log"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"testing"
	"time"
)

var container = func() *appdig.AppContainer {
	container := appdig.NewAppDig()
	container.MustProvides(component.DigComponentProviderAllForTest)
	container.MustProvide(NewCatchwordDao)
	return container
}()

func TestAutoCatchwordDao_QueryList(t *testing.T) {
	container.MustInvoke(func(dao *CatchwordDao) {
		var catchword = randomCatchword()
		err := dao.CreateCatchword(context.Background(), catchword)
		require.Nil(t, err)

		pageInfo, list, err := dao.QueryList(context.Background(), catchword.Title, com_model.PageParams{
			PageNum:  1,
			PageSize: 10,
		})
		println(jsonutil.ToJson(pageInfo), err)
		require.Nil(t, err)
		require.Equal(t, 1, len(list))
		catchword.CreatedAt = list[0].CreatedAt
		catchword.UpdatedAt = list[0].UpdatedAt
		require.Equal(t, catchword, list[0])
		println(jsonutil.ToJson(list), err)

		err = dao.DeleteCatchwordById(context.Background(), catchword.Id)
		require.Nil(t, err)
	})
}

func TestAutoBase(t *testing.T) {
	container.MustInvoke(func(dao *CatchwordDao) {
		var catchword = randomCatchword()
		err := dao.CreateCatchword(context.Background(), catchword)
		// require 失败 不再向下执行
		require.Nil(t, err)

		log.Info(jsonutil.ToJson(catchword))
		catchword.Title = random.GetRandomNum(100)
		catchword.Content = random.GetRandomStr(400)
		require.Nil(t, dao.UpdateCatchwordById(context.Background(), catchword))
		formDb, err := dao.GetCatchwordById(context.Background(), catchword.Id)
		require.Nil(t, err)
		require.Equal(t, catchword.Title, formDb.Title)
		require.Equal(t, catchword.Content, formDb.Content)

		err = dao.DeleteCatchwordById(context.Background(), catchword.Id)
		require.Nil(t, err)
	})
}

func randomCatchword() *model.Catchword {
	return &model.Catchword{
		Model: com_model.Model{
			Id:        "",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt{Valid: false},
		},
		Title:   random.GetRandomStr(7),
		Content: random.GetRandomStr(7),
	}
}

func TestCatchwordDao_QueryListForWebShow(t *testing.T) {
	container.MustInvoke(func(dao *CatchwordDao) {
		show, err := dao.QueryListForWebShow(context.Background(), "", time.Now(), 4)
		require.Nil(t, err)
		println(jsonutil.ToJson(show))
	})
}

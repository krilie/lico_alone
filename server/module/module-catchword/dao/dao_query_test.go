package dao

import (
	"context"
	"github.com/krilie/lico_alone/common/appdig"
	com_model "github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/utils/jsonutil"
	"github.com/krilie/lico_alone/common/utils/random"
	"github.com/krilie/lico_alone/component"
	"github.com/krilie/lico_alone/module/module-catchword/model"
	"github.com/stretchr/testify/assert"
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

func TestAutoCatchwordDao_QueryList(t1 *testing.T) {
	container.MustInvoke(func(dao *CatchwordDao) {
		list, err := dao.QueryList(context.Background(), "a", com_model.PageParams{
			PageNum:  1,
			PageSize: 10,
		})
		println(jsonutil.ToJson(list), err)
	})
}

func TestAutoBase(t *testing.T) {
	container.MustInvoke(func(dao *CatchwordDao) {
		var catchword = randomCatchword()
		err := dao.CreateCatchword(context.Background(), catchword)
		assert.Nil(t, err)

		t.Log(jsonutil.ToJson(catchword))
		catchword.Title = random.GetRandomNum(100)
		catchword.Content = random.GetRandomStr(400)
		assert.Nil(t, dao.UpdateCatchwordById(context.Background(), catchword))
		formDb, err := dao.GetCatchwordById(context.Background(), catchword.Id)
		assert.Nil(t, err)
		assert.Equal(t, catchword.Title, formDb.Title)
		assert.Equal(t, catchword.Content, formDb.Content)

		err = dao.DeleteCatchwordById(context.Background(), catchword.Id)
		assert.Nil(t, err)
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

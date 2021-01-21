package dao

import (
	"context"
	"github.com/krilie/lico_alone/common/appdig"
	com_model "github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/utils/strutil"
	"github.com/krilie/lico_alone/component"
	"testing"
)

var container = func() *appdig.AppContainer {
	container := appdig.NewAppDig()
	container.MustProvides(component.DigComponentProviderAllForTest)
	container.MustProvide(NewCatchwordDao)
	return container
}()

func TestCatchwordDao_QueryList(t1 *testing.T) {
	container.MustInvoke(func(dao *CatchwordDao) {
		list, err := dao.QueryList(context.Background(), "a", com_model.PageParams{
			PageNum:  1,
			PageSize: 10,
		})
		println(strutil.ToJson(list), err)
	})
}

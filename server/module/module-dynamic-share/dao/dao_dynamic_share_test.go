package dao

import (
	"context"
	"github.com/issue9/assert"
	"github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/component"
	"github.com/krilie/lico_alone/module/module-dynamic-share/model"
	"testing"
)

func TestMain(m *testing.M) {
	component.DigProviderTest()
	DigProvider()
	m.Run()
}

var testData = []model.DynamicShare{
	{Model: com_model.NewModel(), Content: "1", Sort: 1},
	{Model: com_model.NewModel(), Content: "2", Sort: 2},
	{Model: com_model.NewModel(), Content: "3", Sort: 3},
}

func TestAutoDynamicShareDao_AddDynamicShare(t *testing.T) {
	dig.Container.MustInvoke(func(dao *DynamicShareDao) {
		err := dao.AddDynamicShare(context.Background(), testData[0])
		assert.Nil(t, err)
		err = dao.UpdateDynamicShare(context.Background(), model.UpdateDynamicShareModel{
			Id: testData[0].Id,
			CreateDynamicShareModel: model.CreateDynamicShareModel{
				Content: "12",
				Sort:    12,
			},
		})
		assert.Nil(t, err)
		var data = model.DynamicShare{}
		err = dao.GetDb(context.Background()).Where("id=?", testData[0].Id).Find(&data).Error
		assert.Nil(t, err)
		assert.Equal(t, data.Content, "12")
		assert.Equal(t, data.Sort, 12)
		err = dao.DeleteDynamicShare(context.Background(), []string{testData[0].Id})
		assert.Nil(t, err)
	})
}

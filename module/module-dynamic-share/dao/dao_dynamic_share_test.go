package dao

import (
	"context"
	"github.com/issue9/assert"
	"github.com/krilie/lico_alone/common/appdig"
	"github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/component"
	"github.com/krilie/lico_alone/module/module-dynamic-share/model"
	"testing"
)

var container = appdig.
	NewAppDig().
	MustProvides(component.DigComponentProviderAll).
	MustProvide(NewDynamicShareDao)

var testData = []model.DynamicShare{
	{Model: com_model.NewModel(), Content: "1", Sort: 1},
	{Model: com_model.NewModel(), Content: "2", Sort: 2},
	{Model: com_model.NewModel(), Content: "3", Sort: 3},
}

func TestAutoDynamicShareDao_AddDynamicShare(t *testing.T) {

	container.MustInvoke(func(dao *DynamicShareDao) {
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
	container.MustInvoke(func(dao *DynamicShareDao) {
		var testData = []model.DynamicShare{
			{Model: com_model.NewModel(), Content: "1", Sort: 1},
			{Model: com_model.NewModel(), Content: "2", Sort: 2},
			{Model: com_model.NewModel(), Content: "3", Sort: 3},
		}
		for _, item := range testData {
			err := dao.AddDynamicShare(context.Background(), item)
			assert.Nil(t, err)
		}
		share, err := dao.QueryDynamicShare(context.Background(), model.QueryDynamicShareModel{
			PageParams: com_model.PageParams{
				PageNum:  1,
				PageSize: 10,
			},
			ContentLike: "",
		})
		assert.Nil(t, err)
		assert.Equal(t, share.TotalCount, 3, "应该是三个")
		assert.Equal(t, len(share.Data), 3, "是三个才对")
		assert.Equal(t, share.Data[0].Sort, 3, "是三才对")
		err = dao.DeleteDynamicShare(context.Background(), []string{testData[0].Id, testData[1].Id, testData[2].Id})
		assert.Nil(t, err)
	})
}

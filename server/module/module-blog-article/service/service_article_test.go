package service

import (
	com_model "github.com/krilie/lico_alone/common/com-model"
	context2 "github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/common/utils/str_util"
	"testing"
)

func TestBlogArticleService_QueryArticleSamplePage(t *testing.T) {
	dig.Container.MustInvoke(func(svc *BlogArticleService) {
		ctx := context2.NewContext()
		page, count, data, err := svc.QueryArticleSamplePage(ctx, com_model.PageParams{
			PageIndex: 1,
			PageSize:  10,
		}, "")
		t.Log(err)
		t.Log(page)
		t.Log(count)
		t.Log(str_util.ToJsonPretty(data))
	})
}

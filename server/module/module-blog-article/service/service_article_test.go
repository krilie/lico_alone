package service

import (
	com_model "github.com/krilie/lico_alone/common/com-model"
	context2 "github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/utils/str_util"
	"testing"
)

var container = BuildTestContainer()

func TestBlogArticleService_QueryArticleSamplePage(t *testing.T) {
	container.MustInvoke(func(svc *BlogArticleModule) {
		ctx := context2.NewContext()
		page, count, data, err := svc.QueryArticleSamplePage(ctx, com_model.PageParams{
			PageNum:  1,
			PageSize: 10,
		}, "1")
		t.Log(err)
		t.Log(page)
		t.Log(count)
		t.Log(str_util.ToJsonPretty(data))
	})
}

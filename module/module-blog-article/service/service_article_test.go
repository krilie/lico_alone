package service

import (
	"github.com/krilie/lico_alone/common/appdig"
	com_model "github.com/krilie/lico_alone/common/com-model"
	context2 "github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/utils/jsonutil"
	"github.com/krilie/lico_alone/component"
	module_like_dislike "github.com/krilie/lico_alone/module/module-like-dislike"
	"testing"
)

var container = buildTestContainer()

// 测试用
func buildTestContainer() *appdig.AppContainer {
	var container = appdig.NewAppDig()
	container.
		MustProvides(component.DigComponentProviderAll).
		MustProvides(DigModuleBlogArticleProviderAll).
		MustProvides(module_like_dislike.DigModuleLikeDisLikeProviderAll)
	return container
}

func TestBlogArticleService_QueryArticleSamplePage(t *testing.T) {
	container.MustInvoke(func(svc *BlogArticleModule) {
		ctx := context2.EmptyAppCtx()
		page, count, data, err := svc.QueryArticleSamplePage(ctx, com_model.PageParams{
			PageNum:  1,
			PageSize: 10,
		}, "1", "sss")
		t.Log(err)
		t.Log(page)
		t.Log(count)
		t.Log(jsonutil.ToJsonPretty(data))
	})
}

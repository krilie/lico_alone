package service_common

import (
	"context"
	common_model "github.com/krilie/lico_alone/common/com-model"
	context2 "github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/component/broker/messages"
	"github.com/krilie/lico_alone/module/module-blog-article/model"
	"time"
)

func (a *CommonService) QueryArticleSamplePage(ctx context.Context, page common_model.PageParams, searchKey, uId string) (pageData *common_model.PageData, err error) {
	totalCount, totalPage, data, err := a.ModuleArticle.QueryArticleSamplePage(ctx, page, searchKey, uId)
	if err != nil {
		return nil, err
	}
	return &common_model.PageData{
		PageInfo: common_model.PageInfo{
			TotalCount: totalCount,
			TotalPage:  totalPage,
			PageNum:    page.PageNum,
			PageSize:   page.PageSize,
		},
		Data: data,
	}, nil
}

func (a *CommonService) GetArticleById(ctx context.Context, articleId, uId string) (*model.QueryArticleModel, error) {
	article, err := a.ModuleArticle.GetArticleById(ctx, articleId, uId)
	if err != nil {
		return nil, err
	}
	if article == nil {
		return nil, errs.NewNotExistsError().WithMsg("未找到")
	}
	a.broker.MustSend(ctx, &messages.BlogArticleVisitedMessage{
		Ctx:             context2.NewAppCtx(ctx, context2.MustGetAppValues(ctx).Clone(nil)),
		VisitedTime:     time.Now(),
		ArticleId:       articleId,
		VisitorIp:       context2.MustGetAppValues(ctx).RemoteIp,
		CustomerTraceId: context2.MustGetAppValues(ctx).CustomerTraceId,
		ArticleTitle:    article.Title,
	})
	return article, nil
}

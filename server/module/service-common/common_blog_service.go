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

func (a *CommonService) QueryArticleSamplePage(ctx context.Context, page common_model.PageParams, searchKey string) (pageData *common_model.PageData, err error) {
	totalCount, totalPage, data, err := a.moduleArticle.QueryArticleSamplePage(ctx, page, searchKey)
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

func (a *CommonService) GetArticleById(ctx context.Context, id string) (*model.Article, error) {
	article, err := a.moduleArticle.GetArticleById(ctx, id)
	if err != nil {
		return nil, err
	}
	if article == nil {
		return nil, errs.NewNotExistsError().WithMsg("未找到")
	}
	a.broker.MustSend(ctx, &messages.BlogArticleVisitedMessage{
		Ctx:         context2.MustGetContext(ctx).Clone().SetTx(nil),
		VisitedTime: time.Now(),
		ArticleId:   id,
		VisitorIp:   context2.MustGetContext(ctx).GetRemoteIp(),
	})
	return article, nil
}

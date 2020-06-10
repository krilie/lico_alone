package service_common

import (
	"context"
	common_model "github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/module/module-blog-article/model"
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
	return article, nil
}

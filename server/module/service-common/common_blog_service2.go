package service_common

import (
	"context"
	common_model "github.com/krilie/lico_alone/common/com-model"
)

func (a *CommonService) QueryArticleSamplePage(ctx context.Context, page common_model.PageParams, searchKey string) (pageData *common_model.PageData, err error) {
	totalPage, totalCount, data, err := a.moduleArticle.QueryArticleSamplePage(ctx, page, searchKey)
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

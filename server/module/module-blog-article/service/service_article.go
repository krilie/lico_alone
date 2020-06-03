package service

import (
	"context"
	common_model "github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/module/module-blog-article/model"
)

func (b *BlogArticleService) CreateArticle(ctx context.Context, article *model.Article) error {
	return b.Dao.CreateArticle(ctx, article)
}

func (b *BlogArticleService) DeleteArticleById(ctx context.Context, id string) (bool, error) {
	return b.Dao.DeleteArticleById(ctx, id)
}

func (b *BlogArticleService) UpdateArticle(ctx context.Context, article *model.Article) error {
	return b.Dao.UpdateArticle(ctx, article)
}

func (b *BlogArticleService) UpdateArticleSample(ctx context.Context, article *model.UpdateArticleModel) error {
	return b.Dao.UpdateArticleSample(ctx, article)
}

func (b *BlogArticleService) QueryArticleById(ctx context.Context, id string) (*model.Article, error) {
	return b.Dao.QueryArticleById(ctx, id)
}

// 分页查询
func (b *BlogArticleService) QueryArticleSamplePage(ctx context.Context, page common_model.PageParams, searchKey string) (totalPage, totalCount int, data []*model.QueryArticleModel, err error) {

	page.CheckOkOrSetDefault()

	data = make([]*model.QueryArticleModel, 0)
	db := b.Dao.GetDb(ctx).Model(new(model.Article))
	if searchKey != "" {
		db = db.Or("title like ?", "%"+searchKey+"%")
		db = db.Or("description like ?", "%"+searchKey+"%")
	}
	err = db.Count(&totalCount).Error
	if err != nil {
		return 0, 0, nil, errs.NewInternal().WithError(err)
	}
	if totalCount <= 0 {
		return 0, 0, data, nil
	}
	totalPage = totalCount / page.PageSize
	// 获取结果
	db = db.Order("created_at desc")
	db = db.Limit(page.PageSize).Offset((page.PageIndex - 1) * page.PageSize)
	err = db.Find(&data).Error
	if err != nil {
		return 0, 0, nil, errs.NewInternal().WithError(err)
	}
	return totalPage, totalCount, data, nil
}

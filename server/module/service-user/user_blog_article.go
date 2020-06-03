package service_user

import (
	"context"
	common_model "github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/module/module-blog-article/model"
)

// 文章管理

func (a *UserService) CreateArticle(ctx context.Context, article *model.Article) error {
	return a.moduleArticle.Dao.CreateArticle(ctx, article)
}

func (a *UserService) DeleteArticleById(ctx context.Context, id string) (bool, error) {
	return a.moduleArticle.Dao.DeleteArticleById(ctx, id)
}

func (a *UserService) UpdateArticle(ctx context.Context, article *model.Article) error {
	return a.moduleArticle.Dao.UpdateArticle(ctx, article)
}

func (a *UserService) QueryArticleById(ctx context.Context, id string) (*model.Article, error) {
	return a.moduleArticle.Dao.QueryArticleById(ctx, id)
}

func (a *UserService) UpdateArticleSample(ctx context.Context, article *model.UpdateArticleModel) error {
	return a.moduleArticle.Dao.UpdateArticleSample(ctx, article)
}

func (a *UserService) QueryArticleSamplePage(ctx context.Context, page common_model.PageParams, searchKey string) (totalPage, totalCount int, data []*model.QueryArticleModel, err error) {
	return a.moduleArticle.QueryArticleSamplePage(ctx, page, searchKey)
}

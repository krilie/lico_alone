package user_service

import (
	"context"
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

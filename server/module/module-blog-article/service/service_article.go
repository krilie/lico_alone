package service

import (
	"context"
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

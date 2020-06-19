package service_user

import (
	"context"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/module/module-blog-article/model"
)

// 文章管理

func (a *UserService) CreateArticle(ctx context.Context, article *model.Article) error {
	return a.ModuleArticle.Dao.CreateArticle(ctx, article)
}

func (a *UserService) DeleteArticleById(ctx context.Context, id string) (bool, error) {
	return a.ModuleArticle.Dao.DeleteArticleById(ctx, id)
}

func (a *UserService) UpdateArticle(ctx context.Context, article *model.Article) error {
	return a.ModuleArticle.Dao.UpdateArticle(ctx, article)
}

func (a *UserService) GetArticleById(ctx context.Context, id string) (*model.Article, error) {
	article, err := a.ModuleArticle.Dao.GetArticleById(ctx, id)
	if err != nil {
		return nil, err
	}
	if article == nil {
		return nil, errs.NewNotExistsError().WithMsg("not exists")
	}
	return article, err
}

func (a *UserService) UpdateArticleSample(ctx context.Context, article *model.UpdateArticleModel) error {
	return a.ModuleArticle.Dao.UpdateArticleSample(ctx, article)
}

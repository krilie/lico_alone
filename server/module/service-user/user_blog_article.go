package service_user

import (
	"context"
	com_model "github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/module/module-blog-article/model"
	"time"
)

// 文章管理

func (a *UserService) CreateArticle(ctx context.Context, article *model.CreateArticleModel) error {
	return a.ModuleArticle.Dao.CreateArticle(ctx, &model.Article{
		Model: com_model.Model{
			Id:        id_util.GetUuid(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: nil,
		},
		Title:       article.Title,
		Pv:          0,
		Content:     article.Content,
		Picture:     article.Picture,
		Description: article.Description,
		Sort:        article.Sort,
	})
}

func (a *UserService) DeleteArticleById(ctx context.Context, id string) (bool, error) {
	return a.ModuleArticle.DeleteArticleById(ctx, id)
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

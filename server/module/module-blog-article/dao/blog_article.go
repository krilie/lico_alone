package dao

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/module/module-blog-article/model"
)

type IBlogArticleDao interface {
	CreateArticle(ctx context.Context, article *model.Article) error
	DeleteArticleById(ctx context.Context, id string) (bool, error)
	UpdateArticle(ctx context.Context, article *model.Article) error
	QueryArticleById(ctx context.Context, id string) (*model.Article, error)
}

func (b *BlogArticleDao) CreateArticle(ctx context.Context, article *model.Article) error {
	if article.Id == "" {
		article.Id = id_util.GetUuid()
	}
	err := b.GetDb(ctx).Model(new(model.Article)).Create(article).Error
	return err
}

func (b *BlogArticleDao) DeleteArticleById(ctx context.Context, id string) (bool, error) {
	err := b.GetDb(ctx).Delete(&model.Article{
		Model: com_model.Model{
			Id: id,
		},
	}).Error
	if gorm.IsRecordNotFoundError(err) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func (b *BlogArticleDao) UpdateArticle(ctx context.Context, article *model.Article) error {
	panic("implement me")
}

func (b *BlogArticleDao) QueryArticleById(ctx context.Context, id string) (*model.Article, error) {
	panic("implement me")
}

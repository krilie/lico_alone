package dao

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/module/module-blog-article/model"
	"time"
)

type IBlogArticleDao interface {
	CreateArticle(ctx context.Context, article *model.Article) error
	DeleteArticleById(ctx context.Context, id string) (bool, error)
	UpdateArticle(ctx context.Context, article *model.Article) error
	UpdateArticleSample(ctx context.Context, article *model.UpdateArticleModel) error
	GetArticleById(ctx context.Context, id string) (*model.Article, error)
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
	result := b.GetDb(ctx).Model(new(model.Article)).Select("*").Update(article)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected <= 0 {
		return errs.NewNotExistsError().WithMsg("没有作出修改")
	}
	return nil
}

func (b *BlogArticleDao) UpdateArticleSample(ctx context.Context, article *model.UpdateArticleModel) error {
	result := b.GetDb(ctx).Model(new(model.Article)).Where("id=?", article.Id).
		UpdateColumns(map[string]interface{}{
			"title":       article.Title,
			"content":     article.Content,
			"picture":     article.Picture,
			"sort":        article.Sort,
			"description": article.Description,
			"updated_at":  time.Now(),
		})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected <= 0 {
		return errs.NewNotExistsError().WithMsg("没有作出修改")
	}
	return nil
}

func (b *BlogArticleDao) GetArticleById(ctx context.Context, id string) (article *model.Article, err error) {
	article = new(model.Article)
	err = b.GetDb(ctx).First(article, "id=?", id).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return article, err
}

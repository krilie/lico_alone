package service

import (
	"context"
	common_model "github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/module/module-blog-article/model"
)

func (b *BlogArticleModule) CreateArticle(ctx context.Context, article *model.Article) error {
	return b.Dao.CreateArticle(ctx, article)
}

func (b *BlogArticleModule) DeleteArticleById(ctx context.Context, id string) (bool, error) {
	return b.Dao.DeleteArticleById(ctx, id)
}

func (b *BlogArticleModule) UpdateArticle(ctx context.Context, article *model.Article) error {
	return b.Dao.UpdateArticle(ctx, article)
}

func (b *BlogArticleModule) UpdateArticleSample(ctx context.Context, article *model.UpdateArticleModel) error {
	return b.Dao.UpdateArticleSample(ctx, article)
}

func (b *BlogArticleModule) GetArticleById(ctx context.Context, id string) (*model.Article, error) {
	return b.Dao.GetArticleById(ctx, id)
}

// 分页查询
func (b *BlogArticleModule) QueryArticleSamplePage(ctx context.Context, page common_model.PageParams, searchKey string) (totalPage, totalCount int, data []*model.QueryArticleModel, err error) {

	page.CheckOkOrSetDefault()

	db := b.Dao.GetDb(ctx).Model(new(model.Article))
	if searchKey != "" {
		db = db.Or("title like ?", ndb.Like(searchKey))
		db = db.Or("description like ?", ndb.Like(searchKey))
	}
	countDb := db
	dataDb := db.Order("created_at desc")
	data = make([]*model.QueryArticleModel, 0)
	totalCount, totalPage, err = ndb.PageGetData(countDb, dataDb, page.PageNum, page.PageSize, &data)
	return totalCount, totalPage, data, err
}

// 分页查询
func (b *BlogArticleModule) QueryArticlePage(ctx context.Context, page common_model.PageParams, searchKey string) (totalPage, totalCount int, data []*model.Article, err error) {

	page.CheckOkOrSetDefault()

	db := b.Dao.GetDb(ctx).Model(new(model.Article))
	if searchKey != "" {
		db = db.Or("title like ?", ndb.Like(searchKey))
		db = db.Or("description like ?", ndb.Like(searchKey))
	}
	countDb := db
	dataDb := db.Order("created_at desc")
	data = make([]*model.Article, 0)
	totalCount, totalPage, err = ndb.PageGetData(countDb, dataDb, page.PageNum, page.PageSize, &data)
	return totalCount, totalPage, data, err
}

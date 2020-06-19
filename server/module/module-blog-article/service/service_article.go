package service

import (
	"context"
	common_model "github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/module/module-blog-article/model"
)

func (s *BlogArticleModule) CreateArticle(ctx context.Context, article *model.Article) error {
	return s.Dao.CreateArticle(ctx, article)
}

func (s *BlogArticleModule) DeleteArticleById(ctx context.Context, id string) (bool, error) {
	return s.Dao.DeleteArticleById(ctx, id)
}

func (s *BlogArticleModule) UpdateArticle(ctx context.Context, article *model.Article) error {
	return s.Dao.UpdateArticle(ctx, article)
}

func (s *BlogArticleModule) UpdateArticleSample(ctx context.Context, article *model.UpdateArticleModel) error {
	return s.Dao.UpdateArticleSample(ctx, article)
}

func (s *BlogArticleModule) GetArticleById(ctx context.Context, id string) (*model.Article, error) {
	return s.Dao.GetArticleById(ctx, id)
}

// 分页查询
func (s *BlogArticleModule) QueryArticleSamplePage(ctx context.Context, page common_model.PageParams, searchKey string) (totalCount, totalPage int, data []*model.QueryArticleModel, err error) {

	page.CheckOkOrSetDefault()

	db := s.Dao.GetDb(ctx).Model(new(model.Article))
	if searchKey != "" {
		db = db.Or("title like ?", ndb.Like(searchKey))
		db = db.Or("description like ?", ndb.Like(searchKey))
	}
	countDb := db
	dataDb := db.Order("sort desc").Order("created_at desc")
	data = make([]*model.QueryArticleModel, 0)
	totalCount, totalPage, err = ndb.PageGetData(countDb, dataDb, page.PageNum, page.PageSize, &data)
	return totalCount, totalPage, data, err
}

// 分页查询
func (s *BlogArticleModule) QueryArticlePage(ctx context.Context, page common_model.PageParams, searchKey string) (totalPage, totalCount int, data []*model.Article, err error) {

	page.CheckOkOrSetDefault()

	db := s.Dao.GetDb(ctx).Model(new(model.Article))
	if searchKey != "" {
		db = db.Or("title like ?", ndb.Like(searchKey))
		db = db.Or("description like ?", ndb.Like(searchKey))
	}
	countDb := db
	dataDb := db.Order("sort desc").Order("created_at desc")
	data = make([]*model.Article, 0)
	totalCount, totalPage, err = ndb.PageGetData(countDb, dataDb, page.PageNum, page.PageSize, &data)
	return totalCount, totalPage, data, err
}

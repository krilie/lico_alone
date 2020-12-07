package service

import (
	"context"
	"github.com/ahmetb/go-linq/v3"
	common_model "github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/module/module-blog-article/model"
	module_like_dislike "github.com/krilie/lico_alone/module/module-like-dislike"
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
func (s *BlogArticleModule) QueryArticleSamplePage(ctx context.Context, page common_model.PageParams, searchKey string) (totalCount, totalPage int, data []*model.QueryArticleModelSample, err error) {

	page.CheckOkOrSetDefault()

	db := s.Dao.GetDb(ctx).Model(new(model.Article))
	if searchKey != "" {
		db = db.Or("title like ?", ndb.Like(searchKey))
		db = db.Or("description like ?", ndb.Like(searchKey))
	}
	countDb := db
	dataDb := db.Order("sort desc").Order("created_at desc")
	data = make([]*model.QueryArticleModelSample, 0)
	totalCount, totalPage, err = ndb.PageGetData(countDb, dataDb, page.PageNum, page.PageSize, &data)
	// 查询 like dislike
	var ids []string
	linq.From(data).SelectT(func(o *model.QueryArticleModelSample) string {
		return o.Id
	}).Distinct().ToSlice(&ids)
	list, err := s.GetLikeDisLikeList(ctx, ids)
	if err != nil {
		return 0, 0, nil, err
	}
	linq.From(data).ForEachT(func(o *model.QueryArticleModelSample) {
		// like
		like := linq.From(list).FirstWithT(func(a *module_like_dislike.LikeDisLikeModelResult) bool {
			if a.BusinessId == o.Id && "like" == a.GiveType {
				return true
			}
			return false
		}).(*module_like_dislike.LikeDisLikeModelResult)
		if like == nil {
			o.Like = 0
		} else {
			o.Like = int(like.Count)
		}
		// dislike
		dislike := linq.From(list).FirstWithT(func(a *module_like_dislike.LikeDisLikeModelResult) bool {
			if a.BusinessId == o.Id && "dislike" == a.GiveType {
				return true
			}
			return false
		}).(*module_like_dislike.LikeDisLikeModelResult)
		if dislike == nil {
			o.DisLike = 0
		} else {
			o.DisLike = int(dislike.Count)
		}
	})

	return totalCount, totalPage, data, err
}

// 分页查询
func (s *BlogArticleModule) QueryArticlePage(ctx context.Context, page common_model.PageParams, searchKey string) (totalPage, totalCount int, tdata []*model.QueryArticleModel, err error) {

	page.CheckOkOrSetDefault()

	db := s.Dao.GetDb(ctx).Model(new(model.Article))
	if searchKey != "" {
		db = db.Or("title like ?", ndb.Like(searchKey))
		db = db.Or("description like ?", ndb.Like(searchKey))
	}
	countDb := db
	dataDb := db.Order("sort desc").Order("created_at desc")
	data := make([]*model.Article, 0)
	totalCount, totalPage, err = ndb.PageGetData(countDb, dataDb, page.PageNum, page.PageSize, &data)
	// 查询like dislike
	// 查询 like dislike
	var ids []string
	linq.From(data).SelectT(func(o *model.Article) string {
		return o.Id
	}).Distinct().ToSlice(&ids)
	list, err := s.GetLikeDisLikeList(ctx, ids)
	if err != nil {
		return 0, 0, nil, err
	}
	linq.From(data).
		SelectT(func(o *model.Article) *model.QueryArticleModel {
			var queryArticle = &model.QueryArticleModel{
				Id:          o.Id,
				Title:       o.Title,
				Picture:     o.Picture,
				Description: o.Description,
				Content:     o.Content,
				Pv:          o.Pv,
				Sort:        o.Sort,
				Like:        0,
				DisLike:     0,
			}
			// like
			like := linq.From(list).FirstWithT(func(a *module_like_dislike.LikeDisLikeModelResult) bool {
				if a.BusinessId == queryArticle.Id && "like" == a.GiveType {
					return true
				}
				return false
			}).(*module_like_dislike.LikeDisLikeModelResult)
			if like == nil {
				queryArticle.Like = 0
			} else {
				queryArticle.Like = int(like.Count)
			}
			// dislike
			dislike := linq.From(list).FirstWithT(func(a *module_like_dislike.LikeDisLikeModelResult) bool {
				if a.BusinessId == queryArticle.Id && "dislike" == a.GiveType {
					return true
				}
				return false
			}).(*module_like_dislike.LikeDisLikeModelResult)
			if dislike == nil {
				queryArticle.DisLike = 0
			} else {
				queryArticle.DisLike = int(dislike.Count)
			}
			return queryArticle
		}).ToSlice(&tdata)

	return totalCount, totalPage, tdata, err
}

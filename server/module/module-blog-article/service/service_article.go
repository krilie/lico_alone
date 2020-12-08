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

func (s *BlogArticleModule) GetArticleById(ctx context.Context, articleId, uId string) (*model.QueryArticleModel, error) {
	article, err := s.Dao.GetArticleById(ctx, articleId)
	if err != nil {
		return nil, err
	}
	if article == nil {
		return nil, nil
	}
	var res = &model.QueryArticleModel{
		Id:          article.Id,
		Title:       article.Title,
		Picture:     article.Picture,
		Description: article.Description,
		Content:     article.Content,
		Pv:          article.Pv,
		Sort:        article.Sort,
		Like:        0,
		DisLike:     0,
		HasLike:     false,
		HasDisLike:  false,
	}
	result, err := s.likeDislikeModule.Dao.GetLikeDiskLikeResult(ctx, "article", []string{articleId})
	if err != nil {
		return nil, err
	}
	linq.From(result).ForEachT(func(a *module_like_dislike.LikeDisLikeModelResult) {
		if a.BusinessId == articleId && a.GiveType == "like" {
			res.Like = int(a.Count)
		}
		if a.BusinessId == articleId && a.GiveType == "dislike" {
			res.DisLike = int(a.Count)
		}
	})
	userResult, err := s.likeDislikeModule.Dao.GetLikeUserLikeDisLike(ctx, "article", []string{articleId}, uId)
	if err != nil {
		return nil, err
	}
	linq.From(userResult).ForEachT(func(a *module_like_dislike.UserLikeDisLikeModelResult) {
		if a.BusinessId == articleId && a.GiveType == "like" && a.Count > 0 {
			res.HasLike = true
		}
		if a.BusinessId == articleId && a.GiveType == "dislike" && a.Count > 0 {
			res.HasDisLike = true
		}
	})
	return res, nil
}

// 分页查询
func (s *BlogArticleModule) QueryArticleSamplePage(ctx context.Context, page common_model.PageParams, searchKey string, uId string) (totalCount, totalPage int, data []*model.QueryArticleModelSample, err error) {

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
	userLikeList, err := s.likeDislikeModule.Dao.GetLikeUserLikeDisLike(ctx, "article", ids, uId)
	if err != nil {
		return 0, 0, nil, err
	}
	linq.From(data).ForEachT(func(o *model.QueryArticleModelSample) {
		// like
		like, ok := linq.From(list).FirstWithT(func(a *module_like_dislike.LikeDisLikeModelResult) bool {
			if a.BusinessId == o.Id && "like" == a.GiveType {
				return true
			}
			return false
		}).(*module_like_dislike.LikeDisLikeModelResult)
		if !ok {
			o.Like = 0
		} else {
			o.Like = int(like.Count)
		}
		hasLike, ok := linq.From(userLikeList).FirstWithT(func(a *module_like_dislike.UserLikeDisLikeModelResult) bool {
			if a.BusinessId == o.Id && "like" == a.GiveType {
				return true
			}
			return false
		}).(*module_like_dislike.UserLikeDisLikeModelResult)
		if !ok {
			o.HasLike = false
		} else {
			if hasLike.Count > 0 {
				o.HasLike = true
			} else {
				o.HasLike = false
			}
		}
		// dislike
		dislike, ok := linq.From(list).FirstWithT(func(a *module_like_dislike.LikeDisLikeModelResult) bool {
			if a.BusinessId == o.Id && "dislike" == a.GiveType {
				return true
			}
			return false
		}).(*module_like_dislike.LikeDisLikeModelResult)
		if !ok {
			o.DisLike = 0
		} else {
			o.DisLike = int(dislike.Count)
		}
		hasDisLike, ok := linq.From(userLikeList).FirstWithT(func(a *module_like_dislike.UserLikeDisLikeModelResult) bool {
			if a.BusinessId == o.Id && "dislike" == a.GiveType {
				return true
			}
			return false
		}).(*module_like_dislike.UserLikeDisLikeModelResult)
		if !ok {
			o.HasDisLike = false
		} else {
			if hasDisLike.Count > 0 {
				o.HasDisLike = true
			} else {
				o.HasDisLike = false
			}
		}
	})

	return totalCount, totalPage, data, err
}

// 分页查询
func (s *BlogArticleModule) QueryArticlePage(ctx context.Context, page common_model.PageParams, searchKey, uId string) (totalPage, totalCount int, tdata []*model.QueryArticleModel, err error) {

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
	userLikeList, err := s.likeDislikeModule.Dao.GetLikeUserLikeDisLike(ctx, "article", ids, uId)
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
				HasLike:     false,
				HasDisLike:  false,
			}
			// like
			like, ok := linq.From(list).FirstWithT(func(a *module_like_dislike.LikeDisLikeModelResult) bool {
				if a.BusinessId == queryArticle.Id && "like" == a.GiveType {
					return true
				}
				return false
			}).(*module_like_dislike.LikeDisLikeModelResult)
			if !ok {
				queryArticle.Like = 0
			} else {
				queryArticle.Like = int(like.Count)
			}
			hasLike, ok := linq.From(userLikeList).FirstWithT(func(a *module_like_dislike.UserLikeDisLikeModelResult) bool {
				if a.BusinessId == o.Id && "like" == a.GiveType {
					return true
				}
				return false
			}).(*module_like_dislike.UserLikeDisLikeModelResult)
			if !ok {
				queryArticle.HasLike = false
			} else {
				if hasLike.Count > 0 {
					queryArticle.HasLike = true
				} else {
					queryArticle.HasLike = false
				}
			}
			// dislike
			dislike, ok := linq.From(list).FirstWithT(func(a *module_like_dislike.LikeDisLikeModelResult) bool {
				if a.BusinessId == queryArticle.Id && "dislike" == a.GiveType {
					return true
				}
				return false
			}).(*module_like_dislike.LikeDisLikeModelResult)
			if !ok {
				queryArticle.DisLike = 0
			} else {
				queryArticle.DisLike = int(dislike.Count)
			}

			hasDisLike, ok := linq.From(userLikeList).FirstWithT(func(a *module_like_dislike.UserLikeDisLikeModelResult) bool {
				if a.BusinessId == o.Id && "dislike" == a.GiveType {
					return true
				}
				return false
			}).(*module_like_dislike.UserLikeDisLikeModelResult)
			if !ok {
				queryArticle.HasDisLike = false
			} else {
				if hasDisLike.Count > 0 {
					queryArticle.HasDisLike = true
				} else {
					queryArticle.HasDisLike = false
				}
			}
			return queryArticle
		}).ToSlice(&tdata)

	return totalCount, totalPage, tdata, err
}

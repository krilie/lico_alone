package service

import (
	"context"
	module_like_dislike "github.com/krilie/lico_alone/module/module-like-dislike"
)

func (s *BlogArticleModule) AddLike(ctx context.Context, uId, bId string) error {
	return s.likeDislikeModule.AddLikeRecord(ctx, uId, bId, "article")
}
func (s *BlogArticleModule) AddDisLike(ctx context.Context, uId, bId string) error {
	return s.likeDislikeModule.AddDisLikeRecord(ctx, uId, bId, "article")
}
func (s *BlogArticleModule) RemoveLike(ctx context.Context, uId, bId string) error {
	return s.likeDislikeModule.RemoveLikeRecord(ctx, uId, bId, "article")
}
func (s *BlogArticleModule) RemoveDisLike(ctx context.Context, uId, bId string) error {
	return s.likeDislikeModule.RemoveDisLikeRecord(ctx, uId, bId, "article")
}

func (s *BlogArticleModule) GetLikeDisLikeList(ctx context.Context, ids []string) ([]module_like_dislike.LikeDisLikeModelResult, error) {
	return s.likeDislikeModule.Dao.GetLikeDiskLikeResult(ctx, "article", ids)
}

package module_like_dislike

import (
	"context"
	"github.com/krilie/lico_alone/component/nlog"
)

type LikeDisLikeModule struct {
	Dao *LikeDisLikeDao
	log *nlog.NLog
}

func NewLikeDisLikeModule(dao *LikeDisLikeDao, log *nlog.NLog) *LikeDisLikeModule {
	return &LikeDisLikeModule{Dao: dao, log: log}
}

func (a *LikeDisLikeModule) AddLikeRecord(ctx context.Context, uId, bId, bType string) error {
	has, err := a.Dao.HasLikeDisLikeRecord(ctx, LikeDisLikeModelParams{
		UserId:       uId,
		BusinessType: bType,
		BusinessId:   bId,
		GiveType:     "like",
	})
	if err != nil {
		return err
	}
	if has {
		return nil
	}
	return a.Dao.AddLikeDisLikeRecord(ctx, LikeDisLikeModelParams{
		UserId:       uId,
		BusinessType: bType,
		BusinessId:   bId,
		GiveType:     "like",
	})
}
func (a *LikeDisLikeModule) AddDisLikeRecord(ctx context.Context, uId, bId, bType string) error {
	has, err := a.Dao.HasLikeDisLikeRecord(ctx, LikeDisLikeModelParams{
		UserId:       uId,
		BusinessType: bType,
		BusinessId:   bId,
		GiveType:     "dislike",
	})
	if err != nil {
		return err
	}
	if has {
		return nil
	}
	return a.Dao.AddLikeDisLikeRecord(ctx, LikeDisLikeModelParams{
		UserId:       uId,
		BusinessType: bType,
		BusinessId:   bId,
		GiveType:     "dislike",
	})
}
func (a *LikeDisLikeModule) RemoveLikeRecord(ctx context.Context, uId, bId, bType string) error {
	has, err := a.Dao.HasLikeDisLikeRecord(ctx, LikeDisLikeModelParams{
		UserId:       uId,
		BusinessType: bType,
		BusinessId:   bId,
		GiveType:     "like",
	})
	if err != nil {
		return err
	}
	if !has {
		return nil
	}
	return a.Dao.RemoveLikeDisLikeRecord(ctx, LikeDisLikeModelParams{
		UserId:       uId,
		BusinessType: bType,
		BusinessId:   bId,
		GiveType:     "like",
	})
}
func (a *LikeDisLikeModule) RemoveDisLikeRecord(ctx context.Context, uId, bId, bType string) error {
	has, err := a.Dao.HasLikeDisLikeRecord(ctx, LikeDisLikeModelParams{
		UserId:       uId,
		BusinessType: bType,
		BusinessId:   bId,
		GiveType:     "dislike",
	})
	if err != nil {
		return err
	}
	if !has {
		return nil
	}
	return a.Dao.RemoveLikeDisLikeRecord(ctx, LikeDisLikeModelParams{
		UserId:       uId,
		BusinessType: bType,
		BusinessId:   bId,
		GiveType:     "dislike",
	})
}

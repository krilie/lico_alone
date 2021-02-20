package module_like_dislike

import (
	"context"
	context_enum "github.com/krilie/lico_alone/common/com-model/context-enum"
	context2 "github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/global"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/component/nlog"
	"time"
)

type LikeDisLikeDao struct {
	*ndb.NDb
	log *nlog.NLog
}

func NewLikeDisLikeDao(db *ndb.NDb, log *nlog.NLog) *LikeDisLikeDao {
	log = log.WithField(context_enum.Module.Str(), "module like dislike")
	if global.EnableAutoMigrate {
		err := db.GetDb(context2.NewAppCtx(context.Background())).AutoMigrate(&LikeDisLikeModel{})
		if err != nil {
			panic(err)
		}
	}
	return &LikeDisLikeDao{
		NDb: db,
		log: log,
	}
}

// AddLikeDisLikeRecord 添加label like dislike shock
func (a *LikeDisLikeDao) AddLikeDisLikeRecord(ctx context.Context, params LikeDisLikeModelParams) error {
	has, err := a.HasLikeDisLikeRecord(ctx, params)
	if err != nil {
		return err
	}
	if has {
		return nil
	}
	sql := `insert into 
                tb_like_dislike(id, created_at, updated_at, deleted_at, user_id, business_type, business_id, give_type)
                VALUES (?,?,?,?,?,?,?,?)`
	exec, err := ndb.Exec(ctx, a.GetDb(ctx), sql, id_util.GetUuid(), time.Now(), time.Now(), nil, params.UserId, params.BusinessType, params.BusinessId, params.GiveType)
	if err != nil {
		return err
	}
	if exec <= 0 {
		return errs.NewInternal().WithMsg("insert failure")
	}
	return nil
}

// AddLikeDisLikeRecord 添加label like dislike shock
func (a *LikeDisLikeDao) RemoveLikeDisLikeRecord(ctx context.Context, params LikeDisLikeModelParams) error {
	has, err := a.HasLikeDisLikeRecord(ctx, params)
	if err != nil {
		return err
	}
	if !has {
		return nil
	}
	sql := `update tb_like_dislike 
            set deleted_at=? 
            where deleted_at is null and give_type=? and business_id =? and business_type=? and user_id=?`
	affected, err := ndb.Exec(ctx, a.GetDb(ctx), sql, time.Now(), params.GiveType, params.BusinessId, params.BusinessType, params.UserId)
	if err != nil {
		return err
	}
	if affected <= 0 {
		return errs.NewInternal().WithMsg("insert failure")
	}
	return nil
}

func (a *LikeDisLikeDao) HasLikeDisLikeRecord(ctx context.Context, params LikeDisLikeModelParams) (bool, error) {
	sql := `select count(id) from tb_like_dislike 
           where deleted_at is null and give_type=? and business_type=? and business_id=? and user_id=?`
	var count = 0
	err := ndb.RawQuery(ctx, a.GetDb(ctx), &count, sql, params.GiveType, params.BusinessType, params.BusinessId, params.UserId)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (a *LikeDisLikeDao) GetLikeDiskLikeResult(ctx context.Context, businessType string, businessIds []string) (res []LikeDisLikeModelResult, err error) {
	sql := `select business_id,business_type,give_type,count(give_type) as count
            from tb_like_dislike 
            where deleted_at is null and business_type=? and business_id in (?) 
            group by business_id, business_type, give_type
            order by count desc `
	err = ndb.RawQuery(ctx, a.GetDb(ctx), &res, sql, businessType, businessIds)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a *LikeDisLikeDao) GetLikeUserLikeDisLike(ctx context.Context, businessType string, businessIds []string, userId string) (res []UserLikeDisLikeModelResult, err error) {
	sql := `select business_id,business_type,give_type,count(give_type) as count
            from tb_like_dislike 
            where deleted_at is null and business_type=? and business_id in (?) and user_id=?
            group by business_id, business_type, give_type
            order by count desc `
	err = ndb.RawQuery(ctx, a.GetDb(ctx), &res, sql, businessType, businessIds, userId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

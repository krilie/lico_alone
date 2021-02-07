package dao

import (
	"context"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/module/module-dynamic-share/model"
)

func (a *DynamicShareDao) AddDynamicShare(ctx context.Context, m model.DynamicShare) error {
	err := a.GetDb(ctx).Create(m).Error
	return err
}

func (a *DynamicShareDao) DeleteDynamicShare(ctx context.Context, ids []string) error {
	err := a.GetDb(ctx).Where("id in ?", ids).Delete(&model.DynamicShare{}).Error
	return err
}

func (a *DynamicShareDao) UpdateDynamicShare(ctx context.Context, u model.UpdateDynamicShareModel) error {
	affected, err := a.Exec(ctx,
		"update tb_dynamic_share set content=?,sort=? where id=? and deleted_at is null",
		u.Content, u.Sort, u.Id)
	if err != nil {
		return err
	}
	if affected == 0 {
		return errs.NewNormal().WithMsg("no updated")
	}
	return nil
}

func (a *DynamicShareDao) QueryDynamicShare(ctx context.Context, param model.QueryDynamicShareModel) (*model.QueryDynamicShareResModel, error) {
	// 开始工作
	log := a.log.Get(ctx, "DynamicShareDao", "QueryDynamicShare")
	log.Info("begin QueryDynamicShare")
	defer log.Info("end QueryDynamicShare")
	param.PageParams.CheckOkOrSetDefault()
	// 查询
	db := a.GetDb(ctx)
	db = db.Model(new(model.DynamicShare))

	if param.ContentLike != "" {
		db.Where("content", ndb.Like(param.ContentLike))
	}
	countDb := db
	dataDb := db.Order("sort desc")

	data := make([]model.DynamicShare, 0)
	count, page, err := ndb.PageGetData(countDb, dataDb, param.PageNum, param.PageSize, &data)
	if err != nil {
		log.WithError(err).Error("sql query err")
		return nil, err
	}
	return &model.QueryDynamicShareResModel{
		TotalPage:  page,
		TotalCount: count,
		PageNum:    param.PageNum,
		PageSize:   param.PageSize,
		Data:       data,
	}, nil
}

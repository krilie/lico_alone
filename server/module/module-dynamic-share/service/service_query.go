package service

import (
	"context"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/module/module-dynamic-share/model"
)

func (a *DynamicShareModule) QueryDynamicShare(ctx context.Context, param model.QueryDynamicShareModel) (*model.QueryDynamicShareResModel, error) {
	// 开始工作
	log := a.log.Get(ctx, "DynamicShare", "QueryDynamicShare")
	log.Info("begin QueryDynamicShare")
	defer log.Info("end QueryDynamicShare")
	param.PageParams.CheckOkOrSetDefault()
	// 查询
	db := a.Dao.GetDb(ctx)
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

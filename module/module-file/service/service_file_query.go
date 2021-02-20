package service

import (
	"context"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/module/module-file/model"
)

func (a *FileModule) QueryFilePage(ctx context.Context, param model.QueryFileParam) (totalPage, totalCount, pageNum, pageSize int64, files []*model.FileMaster, err error) {
	param.PageParams.CheckOkOrSetDefault()
	db := a.dao.GetDb(ctx)
	db = db.Model(new(model.FileMaster))
	if param.UserId != "" {
		db = db.Where("user_id=?", param.UserId)
	}
	if param.BizType != "" {
		db = db.Where("biz_type=?", param.BizType)
	}
	if param.ContentType != "" {
		db = db.Where("content_type=?", param.ContentType)
	}
	if param.BucketNameLike != "" {
		db = db.Where("bucket_name like ?", ndb.Like(param.BucketNameLike))
	}
	if param.KeyNameLike != "" {
		db = db.Where("key_name like ?", ndb.Like(param.KeyNameLike))
	}
	if param.UrlLike != "" {
		db = db.Where("url like ?", ndb.Like(param.UrlLike))
	}
	if param.CreatedAtBegin != nil {
		db = db.Where("created_at >= ?", param.CreatedAtBegin)
	}
	if param.CreatedAtEnd != nil {
		db = db.Where("created_at <= ?", param.CreatedAtEnd)
	}
	countDb := db
	dataDb := db.Order("created_at desc")

	files = make([]*model.FileMaster, 0)
	totalCount, totalPage, err = ndb.PageGetData(countDb, dataDb, param.PageNum, param.PageSize, &files)
	if err != nil {
		return 0, 0, 0, 0, nil, err
	}
	return totalPage, totalCount, param.PageNum, param.PageSize, files, nil
}

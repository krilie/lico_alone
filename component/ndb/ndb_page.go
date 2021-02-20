package ndb

import (
	"context"
	"github.com/krilie/lico_alone/common/errs"
	"gorm.io/gorm"
)

func pageGetCount(db *gorm.DB, pageSize int64) (totalCount, totalPage int64, err error) {
	var totalCount64 int64
	err = db.Count(&totalCount64).Error
	if err != nil {
		return 0, 0, errs.NewInternal().WithError(err)
	}

	if totalCount64 <= 0 {
		return 0, 0, nil
	}
	totalPage = (totalCount64 + pageSize - 1) / pageSize
	return totalCount64, totalPage, nil
}

func PageGetData(countDb, dataDb *gorm.DB, pageNum, pageSize int64, data interface{}) (totalCount, totalPage int64, err error) {
	if pageNum <= 0 || pageSize <= 0 {
		return 0, 0, errs.NewParamError().WithMsg("错误的分页参数")
	}
	count, page, err := pageGetCount(countDb, pageSize)
	if err != nil {
		return 0, 0, err
	}
	if count <= 0 {
		return 0, 0, nil
	}
	err = dataDb.Limit(int(pageSize)).Offset(int((pageNum - 1) * pageSize)).Scan(data).Error
	if err != nil {
		return 0, 0, errs.NewInternal().WithError(err)
	}
	return count, page, nil
}

func Like(val string) string {
	return "%" + val + "%"
}

// 执行一个count查询 一个date查询
func GetPageDataFormSql(ctx context.Context, db *gorm.DB, countSql, dataSql string, countArgs, dataArgs []interface{}, data interface{}) (totalCount int64, err error) {

	totalCount, err = Count(ctx, db, countSql, countArgs...)
	if err != nil {
		return 0, err
	}
	if totalCount <= 0 {
		return 0, nil
	}
	err = db.Raw(dataSql, dataArgs...).Scan(data).Error

	if err != nil {
		return 0, errs.NewInternal().WithError(err)
	}
	return totalCount, nil
}

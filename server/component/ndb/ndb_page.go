package ndb

import (
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/errs"
)

func pageGetCount(db *gorm.DB, pageSize int) (totalCount, totalPage int, err error) {
	err = db.Count(&totalCount).Error
	if err != nil {
		return 0, 0, errs.NewInternal().WithError(err)
	}
	if totalCount <= 0 {
		return 0, 0, nil
	}
	totalPage = (totalCount + pageSize - 1) / pageSize
	return totalCount, totalPage, nil
}

func PageGetData(countDb, dataDb *gorm.DB, pageNum, pageSize int, data interface{}) (totalCount, totalPage int, err error) {
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
	err = dataDb.Limit(pageSize).Offset((pageNum - 1) * pageSize).Scan(data).Error
	if err != nil {
		return 0, 0, errs.NewInternal().WithError(err)
	}
	return count, page, nil
}

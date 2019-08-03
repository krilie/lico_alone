package dbutil

import (
	"github.com/jinzhu/gorm"
	"stone-men/stone/common/cmodel"
)

// FindPage 查询分页数据
func FindPage(db *gorm.DB, pageIndex, pageSize int, out interface{}) (total int, err error) {
	var count int
	result := db.Count(&count)
	if err := result.Error; err != nil {
		return 0, err
	} else if count == 0 {
		return 0, nil
	}

	// 如果分页大小小于0，则不查询数据
	if pageSize < 0 || pageIndex < 0 {
		return count, nil
	}

	if pageIndex > 0 && pageSize > 0 {
		db = db.Offset((pageIndex - 1) * pageSize)
	}
	if pageSize > 0 {
		db = db.Limit(pageSize)
	}
	result = db.Find(out)
	if err := result.Error; err != nil {
		return 0, err
	}

	return count, nil
}

// WrapPageQuery 包装带有分页的查询
func WrapPageQuery(db *gorm.DB, pp *cmodel.PaginationParam, out interface{}) (*cmodel.PaginationResult, error) {
	if pp != nil {
		total, err := FindPage(db, pp.PageIndex, pp.PageSize, out)
		if err != nil {
			return nil, err
		}
		return &cmodel.PaginationResult{
			Total: total,
		}, nil
	}

	result := db.Find(out)
	return nil, result.Error
}

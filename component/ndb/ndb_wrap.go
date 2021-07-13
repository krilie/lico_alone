package ndb

import (
	"context"
	"gorm.io/gorm"
)

// Exec 不返回结果的执行
func Exec(ctx context.Context, db *gorm.DB, sql string, values ...interface{}) (affected int64, err error) {
	resultDb := db.Exec(sql, values...)
	return resultDb.RowsAffected, resultDb.Error
}

// RawQuery 返回结果的执行
func RawQuery(ctx context.Context, db *gorm.DB, outData interface{}, sql string, values ...interface{}) error {
	raw := db.Raw(sql, values...)
	err := raw.Error
	if err != nil {
		return err
	}
	raw.Scan(outData)
	return nil
}

// Count 返回结果的执行 只返回一个数值的 如 select count(1) ...
func Count(ctx context.Context, db *gorm.DB, sql string, values ...interface{}) (count int64, err error) {
	raw := db.Raw(sql, values...)
	err = raw.Error
	if err != nil {
		return 0, err
	}
	raw.Scan(&count)
	return count, nil
}

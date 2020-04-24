package ndb

import (
	"context"
	"github.com/jinzhu/gorm"
)

// Transaction 如果tx==nil则开启新事务
// 如果 tx!=nil 则使用这个tx
func (ndb *NDb) Transaction(ctx context.Context, fc func() error) error {
	tx := GetTxFromCtx(ctx)
	if tx == nil {
		defer func() {
			if err := recover(); err != nil {
			}
		}()
		return ndb.db.Transaction(func(tx *gorm.DB) error {

			return fc()
		})
	} else {
		return fc()
	}
}

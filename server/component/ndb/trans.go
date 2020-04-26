package ndb

import (
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	context2 "github.com/krilie/lico_alone/common/context"
)

// Transaction 如果tx==nil则开启新事务
// 如果 tx!=nil 则使用这个tx
func (ndb *NDb) Transaction(ctx context.Context, fc func(ctx context.Context) error) error {
	tx := GetTxFromCtx(ctx)
	if tx == nil {
		defer func() {
			if err := recover(); err != nil {
				ndb.log.Errorf("事务中发生异常 %v", err)
			}
			ClearTxOnCtl(ctx)
		}()
		return ndb.db.Transaction(func(tx *gorm.DB) error {
			SetTxToCtx(ctx, tx)
			return fc(ctx)
		})
	} else {
		ndb.log.Debug("已经存在事务 不再重新开启事务")
		return fc(ctx)
	}
}

// 无论有没有事务都开启新事务
func (ndb *NDb) TransactionOnNewSession(ctx context.Context, fc func(ctx context.Context) error) error {
	ctx2 := context2.GetContextOrNil(ctx)
	if ctx2 == nil {
		return errors.New("新创建独立事务无效的上下文")
	}
	newCtx := ctx2.Clone()
	newCtx.Tx = nil
	defer func() {
		if err := recover(); err != nil {
			ndb.log.Errorf("新创建独立事务中发生异常已回滚 %v", err)
		}
		ClearTxOnCtl(newCtx)
	}()
	return ndb.db.Transaction(func(tx *gorm.DB) error {
		SetTxToCtx(newCtx, tx)
		return fc(newCtx)
	})
}

package ndb

import (
	"context"
	"errors"
	context2 "github.com/krilie/lico_alone/common/context"
	"gorm.io/gorm"
)

// Transaction 如果tx==nil则开启新事务
// 如果 tx!=nil 则使用这个tx
// gorm2 自动嵌套执行 同一个ctx线程不安全
func (ndb *NDb) Transaction(ctx context.Context, fc func(ctx context.Context) error) (err error) {
	ndb.log.Get(ctx).Trace("开始新的事务")
	defer ndb.log.Get(ctx).Trace("离开事务")
	db := ndb.GetSessionDb(ctx)
	if db == nil {
		db = ndb.db.WithContext(ctx)
		SetTxToCtx(ctx, db) // 临时db
		defer ClearTxOnCtl(ctx)
	}
	panicked := true
	defer func() {
		if panicked || err != nil {
			ndb.log.Get(ctx).Errorf("事务中发生panic或错误 已回滚 %v", err)
		}
	}()
	err = db.Transaction(func(tx *gorm.DB) error {
		SetTxToCtx(ctx, tx)       // 临时db
		defer SetTxToCtx(ctx, db) // 临时db
		return fc(ctx)
	})
	panicked = false
	return err
}

// 无论有没有事务都开启新事务
func (ndb *NDb) TransactionOnNewSession(ctx context.Context, fc func(ctx context.Context) error) (err error) {
	ndb.log.Get(ctx).Trace("在新的事务中执行事务代码")
	defer ndb.log.Get(ctx).Trace("离开新的事务")
	// 原来的上下文
	values := context2.GetAppValues(ctx)
	if values == nil {
		ndb.log.Get(ctx).Error("应用上下文为nil 非法")
		return errors.New("新创建独立事务无效的上下文")
	}
	// 新的上下文
	valuesNew := context2.NewAppCtxValues()
	valuesNew.CopyFrom(values)
	appCtxNew := context2.NewAppCtx(ctx, valuesNew)
	valuesNew.Tx = ndb.db.WithContext(appCtxNew) // 新的session 新的事务
	// 开始事务
	return ndb.Transaction(appCtxNew, fc)
}

type IGetNDb interface {
	GetNDb(ctx context.Context) *NDb
}

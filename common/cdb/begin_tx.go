package cdb

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/ccontext"
	"github.com/krilie/lico_alone/common/clog"
	"github.com/krilie/lico_alone/common/errs"
)

// 如果存在事务 则使用已经存在的事务 如果不存在事务 则使用新的事务
const contextTxKey = "gorm.db.tx"

var txActionOnTrans = txOptionRequired

// 事务处理方式
const txOptionRequired = 1    // 如果有上个事务 则使用上个事务 没有则新创建
const txOptionRequiresNew = 2 // 如果有上个事务 则挂起并开启新事务

func init() {
	txActionOnTrans = txOptionRequired
}

func GetDbFromCtx(ctx context.Context) *gorm.DB {
	var localCtx = ctx.(*ccontext.Context)
	if localCtx == nil {
		return nil
	}
	db, ok := localCtx.Db.(*gorm.DB)
	if ok {
		return db
	} else {
		return nil
	}
}

func SetDbOnCtx(ctx context.Context, db *gorm.DB) {
	localCtx, ok := ctx.(*ccontext.Context)
	if !ok {
		return
	}
	localCtx.Db = db
}

func GetTxFromCtx(ctx context.Context) *gorm.DB {
	localCtx, ok := ctx.(*ccontext.Context)
	if !ok {
		return nil
	}
	tx, ok := localCtx.GetTx().(*gorm.DB)
	if ok {
		return tx
	} else {
		return nil
	}
}

func SetTxOnCtx(ctx context.Context, tx *gorm.DB) {
	localCtx, ok := ctx.(*ccontext.Context)
	if !ok {
		return
	}
	localCtx.SetTx(tx)
}

func IsTxExistsOnCtx(ctx context.Context) bool {
	localCtx, ok := ctx.(*ccontext.Context)
	if !ok {
		return false
	}
	return localCtx.GetTx() != nil
}

func RemoveTxOnCtx(ctx context.Context) {
	localCtx, ok := ctx.(*ccontext.Context)
	if !ok {
		return
	}
	localCtx.SetTx(nil)
}

type Service interface {
	NewWithTx(ctx context.Context, tx *gorm.DB) (service Service, err error) // trans tx优先
	GetDb(ctx context.Context) *gorm.DB
}

func WithTrans(ctx context.Context, oriService Service, txFunc func(ctx context.Context, service Service) error, transOptions ...int) (err error) {
	log = clog.NewLog(ctx, "common.cdb.with_trans", "WithTrans")
	// 事务选项
	var transOption = txOptionRequired
	if len(transOptions) > 0 {
		transOption = transOptions[0]
	}
	// 开启事务
	var thisTx, outTx *gorm.DB
	outTx = GetTxFromCtx(ctx)
	if transOption == txOptionRequired {
		// 需要事务的情况
		if outTx == nil {
			// 开启新事务 并在结束时提交或回滚
			thisTx = oriService.GetDb(ctx)
			thisTx = thisTx.Begin()
			if err := thisTx.Error; err != nil {
				return err
			}
			// 设置当前事务到上下文
			SetTxOnCtx(ctx, thisTx)
			defer SetTxOnCtx(ctx, nil)
			// 开启事务
			service, err := oriService.NewWithTx(ctx, thisTx)
			defer func() {
				if recoverErr := recover(); recoverErr != nil {
					log.Errorln("事务panic:", recoverErr)
					err = fmt.Errorf("%v", recoverErr)
					if err := oriService.GetDb(ctx).Rollback().Error; err != nil {
						log.Error(err.Error())
					}
					return
				}
			}()
			err = txFunc(ctx, service)
			if err != nil {
				log.Debugln("事务回滚")
				if err := service.GetDb(ctx).Rollback().Error; err != nil {
					log.Error(err.Error())
				}
			} else {
				log.Debugln("事务提交")
				if err := service.GetDb(ctx).Commit().Error; err != nil {
					log.Error(err.Error())
				}
			}
			return err
		} else {
			thisTx = outTx
			// 开启事务
			service, err := oriService.NewWithTx(ctx, thisTx)
			defer func() {
				if recoverErr := recover(); recoverErr != nil {
					log.Errorln("事务panic:", recoverErr)
					panic(recoverErr)
				}
			}()
			err = txFunc(ctx, service)
			return err
		}

	} else if transOption == txOptionRequiresNew {
		// 需要新开事务的情况
		// 新上下文
		ctx := ccontext.CloneContext(ctx)
		// 开启新事务 并在结束时提交或回滚
		thisTx = GetDbFromCtx(ctx).Begin()
		if err := thisTx.Error; err != nil {
			return err
		}
		// 设置当前事务到上下文
		SetTxOnCtx(ctx, thisTx)
		defer SetTxOnCtx(ctx, nil)
		// 开启事务
		service, err := oriService.NewWithTx(ctx, thisTx)
		defer func() {
			if recoverErr := recover(); recoverErr != nil {
				log.Errorln("事务panic:", recoverErr)
				err = fmt.Errorf("%v", recoverErr)
				if err := oriService.GetDb(ctx).Rollback().Error; err != nil {
					log.Error(err.Error())
				}
				return
			}
		}()
		err = txFunc(ctx, service)
		if err != nil {
			log.Debugln("事务回滚")
			if err := service.GetDb(ctx).Rollback().Error; err != nil {
				log.Error(err.Error())
			}
		} else {
			log.Debugln("事务提交")
			if err := service.GetDb(ctx).Commit().Error; err != nil {
				log.Error(err.Error())
			}
		}
		return err
	} else {
		panic(errs.NewInternal().WithMsg("db trans option error"))
	}
}

package cdb

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/clog"
)

type Service interface {
	SetTx(ctx context.Context, tx *gorm.DB) (service Service, err error)
	GetDb(ctx context.Context) *gorm.DB
}

func WithTrans(ctx context.Context, oriService Service, txFunc func(service Service) error) (err error) {
	log = clog.NewLog(ctx, "common.cdb.with_trans", "WithTrans")
	gDb := oriService.GetDb(ctx)
	var tx *gorm.DB
	var needRollBackCommit bool
	var service Service
	if IsInTx(gDb) {
		// 已经在事务中不要转化
		service = oriService
		needRollBackCommit = false
		tx = gDb
	} else {
		needRollBackCommit = true
		tx = gDb.Begin()
		err = tx.Error
		if err != nil {
			log.Errorln("事务开启失败:", err.Error())
			return err
		}
		log.Errorln("事务开启成功")
		// 新开事务要set tx 转化
		service, _ = oriService.SetTx(ctx, tx)
	}
	defer func() {
		if rerr := recover(); rerr != nil {
			log.Errorln("事务panic:", rerr)
			err = fmt.Errorf("%v", rerr)
			if needRollBackCommit {
				tx.Rollback()
			}
			return
		}
	}()
	err = txFunc(service)
	if err != nil {
		log.Infoln("事务回滚")
		if needRollBackCommit {
			tx.Rollback()
		}
	} else {
		log.Infoln("事务提交")
		if needRollBackCommit {
			tx.Commit()
		}
	}
	return err
}

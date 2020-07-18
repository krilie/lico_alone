package ndb

import (
	"context"
	"errors"
	context2 "github.com/krilie/lico_alone/common/context"
	"gorm.io/gorm"
)

func GetTxFromCtx(ctx context.Context) *gorm.DB {
	nCtx := context2.GetContextOrNil(ctx)
	if nCtx == nil {
		return nil
	} else {
		if nCtx.Tx == nil {
			return nil
		} else {
			return nCtx.Tx.(*gorm.DB)
		}
	}
}

func SetTxToCtx(ctx context.Context, tx *gorm.DB) {
	nCtx := context2.GetContextOrNil(ctx)
	if nCtx == nil {
		panic(errors.New("无效的上下文"))
	} else {
		nCtx.Tx = tx
	}
}

func ClearTxOnCtl(ctx context.Context) {
	SetTxToCtx(ctx, nil)
}

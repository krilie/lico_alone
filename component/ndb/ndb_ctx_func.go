package ndb

import (
	"context"
	"errors"
	context2 "github.com/krilie/lico_alone/common/context"
	"gorm.io/gorm"
)

func GetTxFromCtx(ctx context.Context) *gorm.DB {
	values := context2.GetAppValues(ctx)
	if values == nil {
		return nil
	} else {
		if values.Tx == nil {
			return nil
		} else {
			return values.Tx.(*gorm.DB)
		}
	}
}

func SetTxToCtx(ctx context.Context, tx *gorm.DB) {
	values := context2.GetAppValues(ctx)
	if values == nil {
		panic(errors.New("无效的上下文"))
	} else {
		values.Tx = tx
	}
}

func ClearTxOnCtl(ctx context.Context) {
	SetTxToCtx(ctx, nil)
}

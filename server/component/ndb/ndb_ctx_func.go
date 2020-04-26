package ndb

import (
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	context2 "github.com/krilie/lico_alone/common/context"
)

func GetTxFromCtx(ctx context.Context) *gorm.DB {
	orNil := context2.GetContextOrNil(ctx)
	if orNil == nil {
		return nil
	} else {
		return orNil.Tx.(*gorm.DB)
	}
}

func SetTxToCtx(ctx context.Context, tx *gorm.DB) {
	orNil := context2.GetContextOrNil(ctx)
	if orNil == nil {
		panic(errors.New("not a app context find"))
	} else {
		orNil.Tx = tx
	}
}

func ClearTxOnCtl(ctx context.Context) {
	SetTxToCtx(ctx, nil)
}

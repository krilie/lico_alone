package service

import (
	"github.com/krilie/lico_alone/common/context"
)

func (Account) DeleteBill(ctx *context.Context, billId string) error {
	// 标记删除
	return nil
}

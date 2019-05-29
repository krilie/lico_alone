package service

import (
	"github.com/krilie/lico_alone/common/context"
	"time"
)

// 总结 统计

// 统计此用户月份的信息
func (Account) GetMonthSummary(ctx *context.Context, userId string, time time.Time) (string, error) {

}

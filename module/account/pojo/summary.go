package pojo

import (
	"github.com/shopspring/decimal"
	"time"
)

// 时区统计数据
type AccountSummary struct {
	BeginTime time.Time       `json:"begin_time"`
	EndTime   time.Time       `json:"end_time"`
	Credit    decimal.Decimal `json:"credit"`
	Debit     decimal.Decimal `json:"debit"`
	Accounts  []AccountItem   `json:"accounts"`
	Bills     []BillItem      `json:"bills"`
}

type AccountItem struct {
	Id         string          `json:"id"`
	Name       string          `json:"name"`
	Num        string          `json:"num"`
	Credit     decimal.Decimal `json:"credit"`
	Debit      decimal.Decimal `json:"debit"`
	Amount     decimal.Decimal `json:"amount"`
	NowBalance decimal.Decimal `json:"now_balance"`
}

type BillItem struct {
	Id         string          `json:"id"`
	CreateTime time.Time       `json:"create_time"`
	Note       string          `json:"note"`
	Amount     decimal.Decimal `json:"amount"`
	Image      string          `json:"image"`
	IsValid    bool            `json:"is_valid"`
}

package model

import (
	"github.com/shopspring/decimal"
)

type BillDetail struct {
	AccountId string          `json:"account_id" form:"account_id" binding:"required"`
	Debit     decimal.Decimal `json:"debit" form:"debit" binding:"required"`
	Credit    decimal.Decimal `json:"credit" form:"credit" binding:"required"`
	Note      string          `json:"note" form:"note" binding:"-"`
}

type AccountInfo struct {
	Id      string          `json:"id"`
	Name    string          `json:"name"`
	Num     string          `json:"num"`
	Balance decimal.Decimal `json:"balance"`
	Image   string          `json:"image"`
}

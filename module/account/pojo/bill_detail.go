package pojo

import (
	"github.com/shopspring/decimal"
)

type BillDetail struct {
	AccountId string          `json:"account_id"`
	Debit     decimal.Decimal `json:"debit"`
	Credit    decimal.Decimal `json:"credit"`
	Note      string          `json:"note"`
}

type AccountInfo struct {
	Id      string          `json:"id"`
	Name    string          `json:"name"`
	Num     string          `json:"num"`
	Balance decimal.Decimal `json:"balance"`
	Image   string          `json:"image"`
}

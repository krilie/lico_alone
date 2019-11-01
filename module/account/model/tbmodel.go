package model

import (
	"github.com/krilie/lico_alone/common/cmodel"
	"time"
)
import "github.com/shopspring/decimal"

type AccountItem struct {
	cmodel.Model
	UpdateTime  time.Time       `gorm:"type:datetime;not null" json:"update_time"`
	UserId      string          `gorm:"type:varchar(32);not null" json:"user_id"`
	Name        string          `gorm:"type:varchar(50);not null" json:"name"`
	Code        string          `gorm:"type:varchar(50);not null" json:"code"`
	Debit       decimal.Decimal `gorm:"type:decimal(14,2);not null;default 0" json:"debit"`
	Credit      decimal.Decimal `gorm:"type:decimal(14,2);not null;default 0" json:"credit"`
	Balance     decimal.Decimal `gorm:"type:decimal(14,2);not null;default 0" json:"balance"`
	Description string          `gorm:"type:varchar(100);default null" json:"description"`
	Image       string          `gorm:"type:varchar(200);not null" json:"image"`
}

func (AccountItem) TableName() string {
	return "tb_account_item"
}

type AccountBill struct {
	cmodel.Model
	UserId  string          `gorm:"type:varchar(32);not null" json:"user_id"`
	Amount  decimal.Decimal `gorm:"type:decimal(14,2);not null;default 0"json:"amount"` // 发生额
	IsValid bool            `gorm:"type:boolean;not null" json:"is_valid"`
	Image   string          `gorm:"type:varchar(500);not null" json:"image"`
	Note    string          `gorm:"type:varchar(100);default null" json:"note"`
}

func (AccountBill) TableName() string {
	return "tb_account_bill"
}

type AccountBillDetail struct {
	cmodel.Model
}

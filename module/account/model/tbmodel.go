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
	PreOpLogIndex int64           `json:"pre_op_log_index" gorm:"column:pre_op_log_index"`
	UserId        string          `gorm:"type:varchar(32);not null" json:"user_id"`
	Amount        decimal.Decimal `gorm:"type:decimal(14,2);not null;default 0"json:"amount"` // 发生额
	IsValid       bool            `gorm:"type:boolean;not null" json:"is_valid"`
	Image         string          `gorm:"type:varchar(500);not null" json:"image"`
	Note          string          `gorm:"type:varchar(100);default null" json:"note"`
}

func (AccountBill) TableName() string {
	return "tb_account_bill"
}

type AccountBillDetail struct {
	cmodel.Model
	BillId        string          `gorm:"type:varchar(32);not null" json:"bill_id"`
	AccountItemId string          `gorm:"type:varchar(32);not null" json:"account_item_id"`
	Amount        decimal.Decimal `gorm:"type:decimal(14,2);not null;default 0"json:"amount"` // 值 可正可负 负借 正贷
}

func (AccountBillDetail) TableName() string {
	return "tb_account_bill_detail"
}

// 操作记录 可撤销最后一次记录不留痕迹
type AccountOperatorLog struct {
	cmodel.Model
	OpIndex       int64           `json:"op_index" gorm:"column:op_index"` // 操作序号
	BillId        string          `gorm:"type:varchar(32);not null" json:"bill_id"`
	AccountItemId string          `gorm:"type:varchar(32);not null" json:"account_item_id"`
	Amount        decimal.Decimal `gorm:"type:decimal(14,2);not null;default 0"json:"amount"` // 值 可正可负 负借 正贷
	Message       string          `gorm:"type:varchar(500);not null" json:"message"`
}

func (AccountOperatorLog) TableName() string {
	return "tb_account_operator_log"
}

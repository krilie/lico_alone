package model

import "time"
import "github.com/shopspring/decimal"

type Account struct {
	Id          string          `gorm:"type:varchar(32);primary_key" json:"id"`
	UserId      string          `gorm:"type:varchar(32);not null" json:"user_id"`
	Name        string          `gorm:"type:varchar(50);not null" json:"name"`
	CreateTime  time.Time       `gorm:"type:DATETIME;not null" json:"create_time"`
	UpdateTime  time.Time       `gorm:"type:DATETIME;not null" json:"update_time"`
	Num         string          `gorm:"type:varchar(50);not null" json:"num"`
	Debit       decimal.Decimal `gorm:"type:decimal(10,2);not null;default 0" json:"debit"`
	Credit      decimal.Decimal `gorm:"type:decimal(10,2);not null;default 0" json:"credit"`
	Balance     decimal.Decimal `gorm:"type:decimal(10,2);not null;default 0" json:"balance"`
	Description string          `gorm:"type:varchar(100);default null" json:"description"`
	Image       string          `gorm:"type:varchar(200);not null" json:"image"`
}

func (Account) TableName() string {
	return "tb_account"
}

type Bill struct {
	Id         string          `gorm:"type:varchar(32);primary_key" json:"id"`
	UserId     string          `gorm:"type:varchar(32);not null" json:"user_id"`
	CreateTime time.Time       `gorm:"type:DATETIME;not null" json:"create_time"`
	Note       string          `gorm:"type:varchar(100);default null" json:"note"`
	Amount     decimal.Decimal `gorm:"type:decimal(10,2);not null;default 0"json:"amount"`
	Image      string          `gorm:"type:varchar(500);not null" json:"image"`
	IsValid    bool            `gorm:"type:boolean;not null" json:"is_valid"`
}

func (Bill) TableName() string {
	return "tb_bill"
}

type BillDetail struct {
	Id          string          `gorm:"type:varchar(32);primary_key" json:"id"`
	UserId      string          `gorm:"type:varchar(32);not null" json:"user_id"`
	BillId      string          `gorm:"type:varchar(32);not null" json:"bill_id"`
	CreateTime  time.Time       `gorm:"type:DATETIME;not null" json:"create_time"`
	AccountId   string          `gorm:"type:varchar(32);not null" json:"account_id"`
	AccountNum  string          `gorm:"type:varchar(50);not null" json:"account_num"`
	AccountName string          `gorm:"type:varchar(50);not null" json:"account_name"`
	Note        string          `gorm:"type:varchar(100);default null" json:"note"`
	Debit       decimal.Decimal `gorm:"type:decimal(10,2);not null;default 0" json:"debit"`
	Credit      decimal.Decimal `gorm:"type:decimal(10,2);not null;default 0" json:"credit"`
	Balance     decimal.Decimal `gorm:"type:decimal(10,2);not null;default 0" json:"balance"`
	IsValid     bool            `gorm:"type:boolean;not null" json:"is_valid"`
}

func (BillDetail) TableName() string {
	return "tb_bill_detail"
}

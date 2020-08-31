package model

import (
	com_model "github.com/krilie/lico_alone/common/com-model"
)

// CustomerAccount email 发送记录
// 記錄訪問者的一些信息
type CustomerAccount struct {
	com_model.Model
	CustomerTraceId string `json:"customer_trace_id" gorm:"column:customer_trace_id;type:char(36);uniqueIndex;not null"`
	LoginName       string `json:"login_name" gorm:"column:login_name;size:256;index;not null"` // 生成唯一代碼
	Password        string `json:"password" gorm:"column:password;size:256;index;not null"`     // 可用登錄名做鹽
	LastAccessIp    string `json:"last_access_ip" gorm:"column:last_access_ip;not null;index"`
	LastAccessAddr  string `json:"last_access_addr" gorm:"column:last_access_addr;not null;index"`
	Mobile          string `json:"mobile" gorm:"column:mobile;size:32;not null;index"`
	Email           string `json:"email" gorm:"column:email;size:32;not null;index"`
	Other           string `json:"other" gorm:"column:other;size:1024;not null"` // 其它
	AccessTimes     int    `json:"access_times" gorm:"column:access_times;type:int;not null"`
}

func (CustomerAccount) TableName() string {
	return "tb_customer_account"
}

type CreateCustomerAccountModel struct {
	CustomerTraceId string `json:"customer_trace_id" gorm:"column:customer_trace_id;type:char(36);uniqueIndex;not null"`
	LoginName       string `json:"login_name" gorm:"column:login_name;size:256;index;not null"` // 生成唯一代碼
	Password        string `json:"password" gorm:"column:password;size:256;index;not null"`     // 可用登錄名做鹽
	LastAccessIp    string `json:"last_access_ip" gorm:"column:last_access_ip;not null;index"`
	Mobile          string `json:"mobile" gorm:"column:mobile;size:32;not null;index"`
	Email           string `json:"email" gorm:"column:email;size:32;not null;index"`
	Other           string `json:"other" gorm:"column:other;size:1024;not null"` // 其它
}

package model

import (
	"github.com/krilie/lico_alone/common/com-model"
	context_enum "github.com/krilie/lico_alone/common/com-model/context-enum"
	"time"
)

// MessageEmail email 发送记录
type MessageEmail struct {
	com_model.Model
	SendTime  time.Time `json:"send_time" gorm:"column:send_time;type:datetime;not null"` // 发送时间
	From      string    `json:"from" gorm:"column:from;size:256;not null"`                // from
	To        string    `json:"to" gorm:"column:to;size:256;not null"`                    // to
	Subject   string    `json:"subject" gorm:"column:subject;size:256;not null"`          // subject
	Content   string    `json:"content" gorm:"column:content;size:1024;not null"`         // 内容
	IsSuccess bool      `json:"is_success" gorm:"column:is_success;not null"`             // 是否成功
	Other     string    `json:"other" gorm:"column:other;size:1024;not null"`             // 其它
}

func (MessageEmail) TableName() string {
	return "tb_message_email"
}

// MessageSms 短信发送记录
type MessageSms struct {
	com_model.Model
	SendTime  time.Time `json:"send_time" gorm:"column:send_time;type:datetime;not null"` // 发送时间
	Name      string    `json:"name" gorm:"column:name;not null;size:32"`                 // 名称
	To        string    `json:"to" gorm:"column:to;not null;size:32"`                     // to
	Message   string    `json:"message" gorm:"column:message;size:512;not null"`          // 消息
	IsSuccess bool      `json:"is_success" gorm:"column:is_success;not null"`             // 是否成功
	Other     string    `json:"other" gorm:"column:other;size:1024;not null"`             // 其它
}

func (MessageSms) TableName() string {
	return "tb_message_sms"
}

// MessageValidCode 短信验证码表 可以是短信也可是邮件的验证码
type MessageValidCode struct {
	com_model.Model
	SendTime time.Time `json:"send_time" gorm:"column:send_time;type:datetime;not null"` // 发送时间
	PhoneNum string    `json:"phone_num" gorm:"column:phone_num;size:32;not null"`       // 手机号
	Code     string    `json:"code" gorm:"column:code;size:16;not null"`                 // 验证码
	Type     int       `json:"type" gorm:"column:type;not null;type:int"`                // 1->登录 2->注册 3->改密码
}

func (MessageValidCode) TableName() string {
	return "tb_message_valid_code"
}

type ValidCodeType = context_enum.IntEnum

const (
	MessageValidCodeTypeLogin        ValidCodeType = 1 // 登录
	MessageValidCodeTypeRegister     ValidCodeType = 2 // 注册
	MessageValidCodeTypeChangePassWd ValidCodeType = 3 // 改密码
)

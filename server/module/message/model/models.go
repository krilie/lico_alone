package model

import (
	"github.com/krilie/lico_alone/common/common-model"
	"time"
)

// MessageEmail email 发送记录
type MessageEmail struct {
	common_model.Model
	SendTime  time.Time `json:"send_time" gorm:"column:send_time;type:datetime"` // 发送时间
	From      string    `json:"from" gorm:"column:from;size:256"`                // from
	To        string    `json:"to" gorm:"column:to;size:256"`                    // to
	Subject   string    `json:"subject" gorm:"column:subject;size:256"`          // subject
	Content   string    `json:"content" gorm:"column:content;size:1024"`         // 内容
	IsSuccess bool      `json:"is_success" gorm:"column:is_success"`             // 是否成功
	Other     string    `json:"other" gorm:"column:other;size:1024"`             // 其它
}

func (MessageEmail) TableName() string {
	return "tb_message_email"
}

// MessageSms 短信发送记录
type MessageSms struct {
	common_model.Model
	SendTime  time.Time `json:"send_time" gorm:"column:send_time;type:datetime"` // 发送时间
	Name      string    `json:"name" gorm:"column:name"`                         // 名称
	To        string    `json:"to" gorm:"column:to"`                             // to
	Message   string    `json:"message" gorm:"column:message;size:512"`          // 消息
	IsSuccess bool      `json:"is_success" gorm:"column:is_success"`             // 是否成功
	Other     string    `json:"other" gorm:"column:other;size:1024"`             // 其它
}

func (MessageSms) TableName() string {
	return "tb_message_sms"
}

// MessageValidCode 短信验证码表 可以是短信也可是邮件的验证码
type MessageValidCode struct {
	common_model.Model
	SendTime time.Time `json:"send_time" gorm:"column:send_time;type:datetime"` // 发送时间
	PhoneNum string    `json:"phone_num" gorm:"column:phone_num;size:32"`       // 手机号
	Code     string    `json:"code" gorm:"column:code;size:16"`                 // 验证码
	Type     int       `json:"type" gorm:"column:type"`                         // 1->登录 2->注册 3->改密码
}

func (MessageValidCode) TableName() string {
	return "tb_message_valid_code"
}

const (
	MessageValidCodeTypeLogin        = 1 // 登录
	MessageValidCodeTypeRegister     = 2 // 注册
	MessageValidCodeTypeChangePassWd = 3 // 改密码
)

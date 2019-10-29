package model

import (
	"github.com/krilie/lico_alone/common/cmodel"
	"time"
)

// MessageEmail email 发送记录
type MessageEmail struct {
	cmodel.Model
	SendTime  time.Time
	From      string
	To        string
	Subject   string
	Content   string
	IsSuccess bool
	Other     string
}

// MessageSms 短信发送记录
type MessageSms struct {
	cmodel.Model
	SendTime  time.Time
	Name      string
	To        string
	Message   string
	IsSuccess bool
	Other     string
}

// MessageValidCode 短信验证码表 可以是短信也可是邮件的验证码
type MessageValidCode struct {
	cmodel.Model
	SendTime time.Time
	PhoneNum string
	Code     string
	Type     int // 1->登录 2->注册 3->改密码
}

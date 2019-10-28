package model

import (
	"github.com/krilie/lico_alone/common/cmodel"
	"time"
)

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

type MessageSms struct {
	cmodel.Model
	SendTime  time.Time
	Name      string
	To        string
	Message   string
	IsSuccess bool
	Other     string
}

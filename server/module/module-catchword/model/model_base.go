package model

import (
	"github.com/krilie/lico_alone/common/com-model"
)

type Catchword struct {
	com_model.Model
	Title   string `json:"title" gorm:"column:title;type:varchar(128);not null;default:''"`      // 消息
	Content string `json:"content" gorm:"column:content;type:varchar(2048);not null;default:''"` // 图片地址
}

func (Catchword) TableName() string {
	return "tb_catchword"
}

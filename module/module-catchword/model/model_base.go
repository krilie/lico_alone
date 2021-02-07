package model

import (
	"github.com/krilie/lico_alone/common/com-model"
	"time"
)

type Catchword struct {
	com_model.Model
	Sort    int64  `json:"sort" gorm:"column:sort;type:int;not null;default:0"`                  // 自增排序字段
	Title   string `json:"title" gorm:"column:title;type:varchar(128);not null;default:''"`      // 消息
	Content string `json:"content" gorm:"column:content;type:varchar(2048);not null;default:''"` // 图片地址
}

func (Catchword) TableName() string {
	return "tb_catchword"
}

type CatchwordVo struct {
	Id        string    `gorm:"column:id;primaryKey;type:char(36)" json:"id"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;not null;type:datetime(3);default:CURRENT_TIMESTAMP(3)"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;not null;type:datetime(3);default:CURRENT_TIMESTAMP(3)"`
	Sort      int64     `json:"sort" gorm:"column:sort;type:long;auto increment;default:0"`           // 自增排序字段
	Title     string    `json:"title" gorm:"column:title;type:varchar(128);not null;default:''"`      // 消息
	Content   string    `json:"content" gorm:"column:content;type:varchar(2048);not null;default:''"` // 图片地址
}

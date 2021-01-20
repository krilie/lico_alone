package model

import (
	"time"
)

type CreateCarouselModel struct {
	Message  string `json:"message" gorm:"column:message;type:varchar(128);not null"` // 消息
	Url      string `json:"url" gorm:"column:url;type:varchar(512);not null"`         // 图片地址
	IsOnShow bool   `json:"is_on_show" gorm:"column:is_on_show;not null"`             // 是否显示
}

// Carousel 轮播图
type CarouselVo struct {
	Id        string    `gorm:"column:id;primaryKey;type:char(36)" json:"id"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;not null;type:datetime(3)"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;not null;type:datetime(3)"`
	Message   string    `json:"message" gorm:"column:message;type:varchar(128);not null"` // 消息
	Url       string    `json:"url" gorm:"column:url;type:varchar(512);not null"`         // 图片地址
	IsOnShow  bool      `json:"is_on_show" gorm:"column:is_on_show;not null"`             // 是否显示
}

package model

import "github.com/krilie/lico_alone/common/com-model"

// Carousel 轮播图
type Carousel struct {
	com_model.Model
	Message  string `json:"message" gorm:"column:message;type:varchar(128);not null"` // 消息
	Url      string `json:"url" gorm:"column:url;type:varchar(512);not null"`         // 图片地址
	IsOnShow bool   `json:"is_on_show" gorm:"column:is_on_show;not null"`             // 是否显示
}

func (Carousel) TableName() string {
	return "tb_carousel_master"
}

type UpdateCarouselModel struct {
	Id       string `json:"id"`         // id
	Message  string `json:"message"`    // 消息
	Url      string `json:"url"`        // 图片地址
	IsOnShow bool   `json:"is_on_show"` // 是否显示
}

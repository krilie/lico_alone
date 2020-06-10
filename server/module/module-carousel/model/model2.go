package model

type CreateCarouselModel struct {
	Id       string `gorm:"column:id;primary_key;type:char(36)" json:"id"`            // id
	Message  string `json:"message" gorm:"column:message;type:varchar(128);not null"` // 消息
	Url      string `json:"url" gorm:"column:url;type:varchar(512);not null"`         // 图片地址
	IsOnShow bool   `json:"is_on_show" gorm:"column:is_on_show;not null"`             // 是否显示
}

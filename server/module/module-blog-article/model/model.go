package model

import (
	common_model "github.com/krilie/lico_alone/common/com-model"
)

type Article struct {
	common_model.Model
	Title   string `json:"title" gorm:"column:title;type:varchar(256);not null"`
	Pv      int    `json:"pv" gorm:"column:pv;type:int;not null"`
	Content string `json:"content" gorm:"column:content;type:text;not null"`
	Picture string `json:"picture" gorm:"column:picture;type:varchar(512);not null"`
}

func (Article) TableName() string {
	return "tb_article_master"
}

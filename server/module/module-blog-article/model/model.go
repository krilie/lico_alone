package model

import (
	"github.com/krilie/lico_alone/common/com-model"
)

type Article struct {
	com_model.Model
	Title       string `json:"title" gorm:"column:title;type:varchar(256);not null"`
	Pv          int    `json:"pv" gorm:"column:pv;type:int;not null"`
	Content     string `json:"content" gorm:"column:content;type:text;not null"`
	Picture     string `json:"picture" gorm:"column:picture;type:varchar(512);not null"`
	Description string `json:"description" gorm:"column:description;type:varchar(512);not null"` // 描述+关键值
	Sort        int    `json:"sort" gorm:"column:sort;type:int;not null;index"`
}

func (Article) TableName() string {
	return "tb_article_master"
}

type UpdateArticleModel struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Picture     string `json:"picture"`
	Description string `json:"description"`
	Sort        int    `json:"sort"`
}

// QueryArticleModel 分页查询简单结果
type QueryArticleModel struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Picture     string `json:"picture"`
	Description string `json:"description"`
	Pv          int    `json:"pv"`
	Sort        int    `json:"sort"`
}

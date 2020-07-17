package model

import (
	com_model "github.com/krilie/lico_alone/common/com-model"
	"time"
)

// StatArticleVisitorLogs 文档访问记录
type StatArticleVisitorLogs struct {
	com_model.Model
	AccessTime      time.Time `json:"access_time" gorm:"column:access_time;type:datetime;index;not null"`
	Ip              string    `json:"ip" gorm:"column:ip;type:nvarchar(64);index;not null"`
	CustomerTraceId string    `json:"customer_trace_id" gorm:"column:customer_trace_id;type:nvarchar(64);index;not null"`
	ArticleId       string    `json:"article_id" gorm:"column:article_id;type:nvarchar(36);not null"`
	ArticleTitle    string    `json:"article_title" gorm:"column:article_title;type:nvarchar(256);not null"`
	// 区域信息
	RegionName string `json:"region_name" gorm:"column:region_name;type:nvarchar(128);index;not null"`
	City       string `json:"city" gorm:"column:city;type:nvarchar(128);index;not null"`
	Memo       string `json:"memo" gorm:"column:memo;type:nvarchar(512);not null"`
}

func (StatArticleVisitorLogs) TableName() string {
	return "tb_stat_article_visitor_logs"
}

type AddStatArticleVisitorModel struct {
	AccessTime      time.Time `json:"access_time"`
	Ip              string    `json:"ip"`
	CustomerTraceId string    `json:"customer_trace_id"`
	ArticleId       string    `json:"article_id"`
	ArticleTitle    string    `json:"article_title"`
	RegionName      string    `json:"region_name"`
	CityName        string    `json:"city_name"`
	Memo            string    `json:"memo"`
}

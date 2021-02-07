package model

import (
	"github.com/krilie/lico_alone/common/com-model"
	"time"
)

// StatVisitorLogs 用户访问记录
type StatVisitorLogs struct {
	com_model.Model
	AccessTime time.Time `json:"access_time" gorm:"column:access_time;type:datetime(3);index;not null"`
	Ip         string    `json:"ip" gorm:"column:ip;type:nvarchar(64);index;not null"`
	TraceId    string    `json:"trace_id" gorm:"column:trace_id;type:nvarchar(64);index;not null"`
	RegionName string    `json:"region_name" gorm:"column:region_name;type:nvarchar(128);index;not null"`
	City       string    `json:"city" gorm:"column:city;type:nvarchar(128);index;not null"`
	Memo       string    `json:"memo" gorm:"column:memo;type:nvarchar(512);not null"`
}

func (StatVisitorLogs) TableName() string {
	return "tb_stat_visitor_logs"
}

type AddStatVisitorLogsModel struct {
	AccessTime time.Time
	Ip         string
	TraceId    string
	RegionName string
	CityName   string
	Memo       string
}

package model

import (
	"time"
)

type Config struct {
	Name       string    `json:"name" gorm:"column:name;size:255;primary_key"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;type:datetime;not null" json:"create_time"`
	Value      string    `json:"value" gorm:"column:value;type:text;not null"`
}

func (Config) TableName() string {
	return "tb_config_master"
}

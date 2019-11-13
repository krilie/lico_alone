package model

import (
	"time"
)

type Config struct {
	Name       string    `gorm:"column:name;size:255;primary_key"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null" json:"create_time"`
	Value      string    `gorm:"column:value;size:5000"`
}

func (Config) TableName() string {
	return "tb_config"
}

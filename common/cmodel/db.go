package cmodel

import "time"

type Model struct {
	Id         string    `gorm:"column:id;primary_key;type:varchar(32)" json:"id"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null" json:"create_time"`
}

package model

import "time"

type Model struct {
	Id         string    `gorm:"column:id;primary_key;type:char(36)" json:"id"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null" json:"create_time"`
}

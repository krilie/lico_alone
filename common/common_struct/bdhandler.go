package common_struct

import "time"

type DbHandler struct {
	ID         string    `gorm:"primary_key;type:varchar(32)" json:"id"` // 用户id uuid
	CreateTime time.Time `gorm:"type:DATETIME;not null" json:"create_time"`
}

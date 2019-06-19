package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID         string         `gorm:"primary_key;type:varchar(32)" json:"id"` // 用户id uuid
	CreateTime time.Time      `gorm:"type:DATETIME;not null" json:"create_time"`
	LoginName  string         `gorm:"type:varchar(50);unique_index;not null"`
	NickName   string         `gorm:"type:varchar(50)"`               // 呢称
	Phone      sql.NullString `gorm:"type:varchar(20);unique_index"`  // 电话
	Email      sql.NullString `gorm:"type:varchar(100);unique_index"` // email
	Password   string         `gorm:"type:varchar(64);not null"`      // 密码md5
	Picture    sql.NullString `gorm:"type:varchar(500)"`              //用户头像
	Salt       string         `gorm:"type:varchar(8);not null"`       // 盐值
}

func (User) TableName() string {
	return "tb_user"
}

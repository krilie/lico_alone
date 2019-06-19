package model

import (
	"time"
)

// 上传的文件的内容
type File struct {
	ID          string    `gorm:"primary_key;type:varchar(32)" json:"id"` // 用户id uuid
	CreateTime  time.Time `gorm:"type:DATETIME;not null" json:"create_time"`
	ObjKey      string    `gorm:"type:varchar(200);unique_index;not null"`
	UserId      string    `gorm:"type:varchar(32);not null"`
	ContentType string    `gorm:"type:varchar(50)"`
	BizType     string    `gorm:"type:varchar(50)"`
	Size        int       `gorm:"type:int;not null"`
}

func (File) TableName() string {
	return "tb_file"
}

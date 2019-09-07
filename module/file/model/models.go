package model

import (
	"time"
)

// 上传的文件的内容
type FileMaster struct {
	Id          string    `gorm:"column:id;primary_key;type:varchar(32)" json:"id"` // 用户id uuid
	CreateTime  time.Time `gorm:"column:create_time;type:DATETIME;not null" json:"create_time"`
	KeyName     string    `gorm:"column:key_name;type:varchar(200);unique_index;not null"`
	BucketName  string    `gorm:"column:bucket_name;type:varchar(200);unique_index;not null"`
	UserId      string    `gorm:"column:user_id;type:varchar(32);not null"`
	ContentType string    `gorm:"column:content_type;type:varchar(50)"`
	BizType     string    `gorm:"column:biz_type;type:varchar(50)"`
	Size        int       `gorm:"column:size;type:int;not null"`
}

func (FileMaster) TableName() string {
	return "tb_file_master"
}

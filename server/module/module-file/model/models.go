package model

import common_model "github.com/krilie/lico_alone/common/com-model"

// 上传的文件的内容
type FileMaster struct {
	common_model.Model
	KeyName     string `gorm:"column:key_name;type:varchar(200);unique_index;not null"`
	BucketName  string `gorm:"column:bucket_name;type:varchar(200);not null"`
	Url         string `gorm:"column:url;type:varchar(200);unique_index;not null"`
	UserId      string `gorm:"column:user_id;type:char(36);not null"`
	ContentType string `gorm:"column:content_type;type:varchar(50);not null"`
	BizType     string `gorm:"column:biz_type;type:varchar(50);not null"`
	Size        int    `gorm:"column:size;type:int;not null"`
}

func (FileMaster) TableName() string {
	return "tb_file_master"
}

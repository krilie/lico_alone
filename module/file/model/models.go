package model

import "github.com/krilie/lico_alone/common/common_struct"

// 上传的文件的内容
type File struct {
	common_struct.DbHandler
	ObjKey      string `gorm:"type:varchar(200);unique_index;not null"`
	ContentType string `gorm:"type:varchar(50)"`
	BizType     string `gorm:"type:varchar(50)"`
	Size        int    `gorm:"type:int;not null"`
}

func (File) TableName() string {
	return "tb_file"
}

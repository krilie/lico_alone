package model

import "github.com/krilie/lico_alone/common/comstruct"

// 上传的文件的内容
type File struct {
	comstruct.DbHandler
	ObjKey      string `gorm:"type:varchar(200);unique_index;not null"`
	UserId      string `gorm:"type:varchar(32);not null"`
	ContentType string `gorm:"type:varchar(50)"`
	BizType     string `gorm:"type:varchar(50)"`
	Size        int    `gorm:"type:int;not null"`
}

func (File) TableName() string {
	return "tb_file"
}

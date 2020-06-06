package model

import (
	"github.com/krilie/lico_alone/common/com-model"
	"time"
)

// 上传的文件的内容
type FileMaster struct {
	com_model.Model
	KeyName     string `json:"key_name" gorm:"column:key_name;type:varchar(200);unique_index;not null"`
	BucketName  string `json:"bucket_name" gorm:"column:bucket_name;type:varchar(200);not null"`
	Url         string `json:"url" gorm:"column:url;type:varchar(200);unique_index;not null"`
	UserId      string `json:"user_id" gorm:"column:user_id;type:char(36);not null"`
	ContentType string `json:"content_type" gorm:"column:content_type;type:varchar(50);not null"`
	BizType     string `json:"biz_type" gorm:"column:biz_type;type:varchar(50);not null"`
	Size        int    `json:"size" gorm:"column:size;type:int;not null"`
}

func (FileMaster) TableName() string {
	return "tb_file_master"
}

// @Param page_num query int true "page_num页索引"
// @Param page_size query int true "page_size页大小"
// @Param key_name_like formData string true "key_name_like"
// @Param bucket_name_like formData string true "bucket_name_like"
// @Param url_like formData string true "url_like"
// @Param user_id formData string true "user_id"
// @Param biz_type formData string true "biz_type"
// @Param content_type formData string true "content_type"
// @Param created_at_begin formData string true "created_at_begin"
// @Param created_at_end formData string true "created_at_end"
type QueryFileParam struct {
	com_model.PageParams
	KeyNameLike    string     `json:"key_name_like" form:"key_name_like" xml:"key_name_like" `
	BucketNameLike string     `json:"bucket_name_like" form:"bucket_name_like" xml:"bucket_name_like" `
	UrlLike        string     `json:"url_like" form:"url_like" xml:"url_like" `
	UserId         string     `json:"user_id" form:"user_id" xml:"user_id" `
	BizType        string     `json:"biz_type" form:"biz_type" xml:"biz_type" `
	ContentType    string     `json:"content_type" form:"content_type" xml:"content_type" `
	CreatedAtBegin *time.Time `json:"created_at_begin" form:"created_at_begin" xml:"created_at_begin" `
	CreatedAtEnd   *time.Time `json:"created_at_end" form:"created_at_end" xml:"created_at_end" `
}

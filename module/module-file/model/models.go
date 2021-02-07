package model

import (
	"github.com/krilie/lico_alone/common/com-model"
	"time"
)

// 上传的文件的内容
type FileMaster struct {
	com_model.Model
	KeyName     string `json:"key_name" gorm:"column:key_name;type:varchar(200);uniqueIndex;not null"`
	BucketName  string `json:"bucket_name" gorm:"column:bucket_name;type:varchar(200);not null"`
	Url         string `json:"url" gorm:"column:url;type:varchar(200);uniqueIndex;not null"`
	UserId      string `json:"user_id" gorm:"column:user_id;type:char(36);not null"`
	ContentType string `json:"content_type" gorm:"column:content_type;type:varchar(50);not null"`
	BizType     string `json:"biz_type" gorm:"column:biz_type;type:varchar(50);not null"`
	Size        int    `json:"size" gorm:"column:size;type:int(11);not null"`
}

func (FileMaster) TableName() string {
	return "tb_file_master"
}

func (f FileMaster) ToDto() FileMasterDto {
	return FileMasterDto{
		ModelVo: com_model.ModelVo{
			Id:        f.Model.Id,
			CreatedAt: f.Model.CreatedAt,
			UpdatedAt: f.Model.UpdatedAt,
			DeletedAt: func() *time.Time {
				if f.Model.DeletedAt.Valid {
					return &f.Model.DeletedAt.Time
				} else {
					return nil
				}
			}(),
		},
		KeyName:     f.KeyName,
		BucketName:  f.BucketName,
		Url:         f.Url,
		UserId:      f.UserId,
		ContentType: f.ContentType,
		BizType:     f.BizType,
		Size:        f.Size,
	}
}

type QueryFileParam struct {
	com_model.PageParams
	KeyNameLike    string     `json:"key_name_like" form:"key_name_like" xml:"key_name_like" `
	BucketNameLike string     `json:"bucket_name_like" form:"bucket_name_like" xml:"bucket_name_like" `
	UrlLike        string     `json:"url_like" form:"url_like" xml:"url_like" `
	UserId         string     `json:"user_id" form:"user_id" xml:"user_id" `
	BizType        string     `json:"biz_type" form:"biz_type" xml:"biz_type" `
	ContentType    string     `json:"content_type" form:"content_type" xml:"content_type" `
	CreatedAtBegin *time.Time `json:"created_at_begin" form:"created_at_begin" xml:"created_at_begin" time_format:"2006-01-02T15:04:05Z07:00"` // rfc3339
	CreatedAtEnd   *time.Time `json:"created_at_end" form:"created_at_end" xml:"created_at_end" time_format:"2006-01-02T15:04:05Z07:00"`       // rfc3339
}

// 上传的文件的内容
type FileMasterDto struct {
	com_model.ModelVo
	KeyName     string `json:"key_name" gorm:"column:key_name;type:varchar(200);uniqueIndex;not null"`
	BucketName  string `json:"bucket_name" gorm:"column:bucket_name;type:varchar(200);not null"`
	Url         string `json:"url" gorm:"column:url;type:varchar(200);uniqueIndex;not null"`
	UserId      string `json:"user_id" gorm:"column:user_id;type:char(36);not null"`
	ContentType string `json:"content_type" gorm:"column:content_type;type:varchar(50);not null"`
	BizType     string `json:"biz_type" gorm:"column:biz_type;type:varchar(50);not null"`
	Size        int    `json:"size" gorm:"column:size;type:int(11);not null"`
}

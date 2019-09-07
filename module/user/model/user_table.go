package model

import (
	"github.com/krilie/lico_alone/common/cmodel"
	"time"
)

type UserMaster struct {
	cmodel.Model
	UpdateTime time.Time `gorm:"column:update_time;type:DATETIME;not null" json:"create_time"` // 创建时间
	LoginName  string    `gorm:"column:login_name;type:varchar(50)"`                           // 呢称
	PhoneNum   *string   `gorm:"column:phone_num;type:varchar(20);unique_index"`               // 电话
	Email      *string   `gorm:"column:email;type:varchar(100);unique_index"`                  // email
	Password   string    `gorm:"column:password;type:varchar(64);not null"`                    // 密码md5
	Picture    *string   `gorm:"column:picture;type:varchar(500)"`                             // 用户头像
	Salt       string    `gorm:"column:salt;type:varchar(8);not null"`                         // 盐值
}

func (user UserMaster) TableName() string {
	return "tb_user_master"
}

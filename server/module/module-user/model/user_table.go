package model

import (
	"github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/utils/time_util"
	"time"
)

type UserMaster struct {
	com_model.Model
	LoginName string `gorm:"column:login_name;type:varchar(50);not null"`            // 呢称
	PhoneNum  string `gorm:"column:phone_num;type:varchar(20);uniqueIndex;not null"` // 电话
	Email     string `gorm:"column:email;type:varchar(100);not null"`                // email
	Password  string `gorm:"column:password;type:varchar(64);not null"`              // 密码md5
	Picture   string `gorm:"column:picture;type:varchar(500);not null"`              // 用户头像
	Salt      string `gorm:"column:salt;type:varchar(8);not null"`                   // 盐值
}

func (user UserMaster) TableName() string {
	return "tb_user_master"
}

func NewUserMaster(Id string, CreatedAt time.Time, UpdatedAt time.Time, DeletedAt *time.Time, UpdateTime time.Time,
	LoginName, PhoneNum, Email, Password, Picture, Salt string) *UserMaster {
	return &UserMaster{
		Model: com_model.Model{
			Id:        Id,
			CreatedAt: CreatedAt,
			UpdatedAt: UpdatedAt,
			DeletedAt: time_util.SqlNullTime(DeletedAt),
		},
		LoginName: LoginName,
		PhoneNum:  PhoneNum,
		Email:     Email,
		Password:  Password,
		Picture:   Picture,
		Salt:      Salt,
	}
}

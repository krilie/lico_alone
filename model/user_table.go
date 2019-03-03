package model

type User struct {
	DbHandler        //共有字段
	LoginName string `gorm:"type:varchar(20);unique_index"`
	NickName  string `gorm:"type:varchar(20)"`               // 呢称
	Phone     string `gorm:"type:varchar(20);unique_index"`  // 电话
	Email     string `gorm:"type:varchar(100);unique_index"` // email
	Password  string `gorm:"type:varchar(64)"`               // 密码md5
	Salt      string `gorm:"type:varchar(8)"`                // 盐值
}

func (User) TableName() string {
	return "tb_user"
}

package model

import (
	"github.com/jinzhu/gorm"
	"github.com/lico603/lico-my-site-user/common/log"
)

var Db *gorm.DB

func init() {
	if db, err := gorm.Open("mysql", "root:9@tcp(localhost:3306)/my?charset=utf8&parseTime=True&loc=Local"); err != nil {
		log.Log.Panicln(err)
		return
	} else {
		Db = db
		Db.CreateTable(&User{})
		Db.CreateTable(&Role{})
		Db.CreateTable(&Permission{})
		Db.CreateTable(&RolePermission{})
		Db.CreateTable(&UserRole{})
	}
}

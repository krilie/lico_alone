package model

import (
	"github.com/jinzhu/gorm"
	"github.com/lico603/lico-my-site-user/common/log"
	"time"
)

var Db *gorm.DB

func init() {
	if db, err := gorm.Open("mysql", "root:123456@tcp(192.168.31.238:3306)/test?charset=utf8&parseTime=True&loc=Local"); err != nil {
		log.Log.Panicln(err)
		return
	} else {
		Db = db
		Db.CreateTable(&User{})
		Db.CreateTable(&Role{})
		Db.CreateTable(&Permission{})
		Db.CreateTable(&RolePermission{})
		Db.CreateTable(&UserRole{})

		Db.DB().SetMaxOpenConns(1)
		Db.DB().SetMaxIdleConns(1)
		Db.DB().SetConnMaxLifetime(time.Hour * 7)
	}
}

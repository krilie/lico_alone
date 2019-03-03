package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/lico603/lico-my-site-user/common/config"
	"github.com/lico603/lico-my-site-user/common/log"
	"time"
)

type DbHandler struct {
	ID         string    `gorm:"primary_key;type:varchar(32)"` // 用户id uuid
	CreateTime time.Time `gorm:"type:DATETIME;not null"`
	Version    int       `gorm:"not null;default:0"`
}

var Db *gorm.DB

func init() {
	connStr := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		config.GetString("db.user"),
		config.GetString("db.password"),
		config.GetString("db.ip"),
		config.GetInt("db.port"),
		config.GetString("db.database"),
	)

	if db, err := gorm.Open("mysql", connStr); err != nil {
		log.Log.Panicln(err)
		return
	} else {
		Db = db

		if !Db.HasTable(&User{}) {
			if err := Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&User{}).Error; err != nil {
				log.Log.Panicln(err)
			}
		}
		if !Db.HasTable(&Role{}) {
			if err := Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&Role{}).Error; err != nil {
				log.Log.Panicln(err)
			}
		}
		if !Db.HasTable(&Permission{}) {
			if err := Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&Permission{}).Error; err != nil {
				log.Log.Panicln(err)
			}
		}
		if !Db.HasTable(&RolePermission{}) {
			if err := Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&RolePermission{}).Error; err != nil {
				log.Log.Panicln(err)
			}
		}
		if !Db.HasTable(&UserRole{}) {
			if err := Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&UserRole{}).Error; err != nil {
				log.Log.Panicln(err)
			}
		}
		if !Db.HasTable(&AppUserAccessToken{}) {
			if err := Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&AppUserAccessToken{}).Error; err != nil {
				log.Log.Panicln(err)
			}
		}

		Db.DB().SetMaxOpenConns(config.GetInt("db.max_open_conn"))
		Db.DB().SetMaxIdleConns(config.GetInt("db.max_idle_conn"))
		Db.DB().SetConnMaxLifetime(time.Second * time.Duration(config.GetInt("db.conn_max_left_time")))
	}
}

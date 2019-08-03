package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/krilie/lico_alone/common/comlog"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/common/context"
	"time"
)

var Db *gorm.DB
var log = comlog.NewLog(context.NewContext(), "alone.module.userbase.model", "init")

func init() {
	connStr := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True",
		config.GetString("db.user"),
		config.GetString("db.password"),
		config.GetString("db.ip"),
		config.GetInt("db.port"),
		config.GetString("db.database"),
	)

	if db, err := gorm.Open("mysql", connStr); err != nil {
		log.Panicln(err)
		return
	} else {
		Db = db

		if !Db.HasTable(&User{}) {
			if err := Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&User{}).Error; err != nil {
				log.Panicln(err)
			}
		}
		if !Db.HasTable(&Role{}) {
			if err := Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&Role{}).Error; err != nil {
				log.Panicln(err)
			}
		}
		if !Db.HasTable(&Permission{}) {
			if err := Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&Permission{}).Error; err != nil {
				log.Panicln(err)
			}
		}
		if !Db.HasTable(&RolePermission{}) {
			if err := Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&RolePermission{}).Error; err != nil {
				log.Panicln(err)
			}
		}
		if !Db.HasTable(&UserRole{}) {
			if err := Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&UserRole{}).Error; err != nil {
				log.Panicln(err)
			}
		}
		if !Db.HasTable(&ClientUserAccessToken{}) {
			if err := Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&ClientUserAccessToken{}).Error; err != nil {
				log.Panicln(err)
			}
		}

		Db.DB().SetMaxOpenConns(config.GetInt("db.max_open_conn"))
		Db.DB().SetMaxIdleConns(config.GetInt("db.max_idle_conn"))
		Db.DB().SetConnMaxLifetime(time.Second * time.Duration(config.GetInt("db.conn_max_left_time")))
	}
}

package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/comlog"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/common/context"
	"github.com/sirupsen/logrus"
	"runtime/debug"
	"time"
)
import _ "github.com/jinzhu/gorm/dialects/mysql"

var Db *gorm.DB
var log *logrus.Entry

func init() {
	log = comlog.NewLog(context.NewContext(), "br_go.common.db")

	connStr := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True",
		config.GetString("db.user"),
		config.GetString("db.password"),
		config.GetString("db.ip"),
		config.GetInt("db.port"),
		config.GetString("db.database"),
	)
	var err error
	if Db, err = gorm.Open("mysql", connStr+"&loc=Asia%2FShanghai"); err != nil {
		log.Fatal(err, debug.Stack()) // 报错退出程序
		return
	} else {
		Db.DB().SetMaxOpenConns(config.GetInt("db.max_open_conn"))
		Db.DB().SetMaxIdleConns(config.GetInt("db.max_idle_conn"))
		Db.DB().SetConnMaxLifetime(time.Second * time.Duration(config.GetInt("db.conn_max_left_time")))
		log.Info("db init done. params:", connStr+"&loc=Asia%2FShanghai") // 数据库初始化成功
	}
}

func Close() {
	err := Db.Close()
	if err != nil {
		log.Warn(err)
	} else {
		log.Info("db closed.")
	}
}

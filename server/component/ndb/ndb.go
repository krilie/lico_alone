package ndb

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/component/nlog"
	"runtime/debug"
	"sync"
	"time"
)

type NDb struct {
	cfg         config.DB
	log         *nlog.NLog
	onceStartDb sync.Once
	onceStopDb  sync.Once
	Db          *gorm.DB
}

func (ndb *NDb) Start() {
	ndb.onceStartDb.Do(func() {
		connStr := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True",
			ndb.cfg.User,
			ndb.cfg.Password,
			ndb.cfg.Host,
			ndb.cfg.Port,
			ndb.cfg.DbName,
		)
		var err error
		if ndb.Db, err = gorm.Open("mysql", connStr+"&loc=Asia%2FShanghai"); err != nil {
			fmt.Println(err.Error())
			ndb.log.Fatal(err, string(debug.Stack())) // 报错退出程序
			return
		} else {
			ndb.Db.DB().SetMaxOpenConns(ndb.cfg.MaxOpenConn)
			ndb.Db.DB().SetMaxIdleConns(ndb.cfg.MaxIdleConn)
			ndb.Db.DB().SetConnMaxLifetime(time.Second * time.Duration(ndb.cfg.ConnMaxLeftTime))
			ndb.log.Info("db init done. params:", connStr+"&loc=Asia%2FShanghai") // 数据库初始化成功
			ndb.Db = ndb.Db.Debug()
		}
	})
}

func (ndb *NDb) CloseDb() {
	ndb.onceStopDb.Do(func() {
		err := ndb.Db.Close()
		if err != nil {
			ndb.log.Warn(err)
		} else {
			ndb.log.Info("db closed.")
		}
	})
}

func NewNDb(dbCfg config.DB, log *nlog.NLog) (closer func(), ndb *NDb) {
	ndb = &NDb{log: log, cfg: dbCfg}
	ndb.Start()
	return ndb.CloseDb, ndb
}

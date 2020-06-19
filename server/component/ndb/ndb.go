package ndb

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	context2 "github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/component/ncfg"
	"github.com/krilie/lico_alone/component/nlog"
	"runtime/debug"
	"sync"
	"time"
)

const gormTransConDb = "gormTransConDb"

type NDb struct {
	cfg struct {
		ConnStr         string
		MaxOpenConn     int
		MaxIdleConn     int
		ConnMaxLeftTime int
	}
	log         *nlog.NLog
	onceStartDb sync.Once
	onceStopDb  sync.Once
	db          *gorm.DB
}

func (ndb *NDb) GetDb(ctx context.Context) *gorm.DB {
	nCtx := context2.GetContextOrNil(ctx)
	if nCtx == nil {
		return ndb.db
	} else {
		if nCtx.Tx == nil {
			return ndb.db
		} else {
			return nCtx.Tx.(*gorm.DB)
		}
	}
}

func (ndb *NDb) Ping() error {
	return ndb.db.DB().Ping()
}

func (ndb *NDb) Start() {
	ndb.onceStartDb.Do(func() {
		var err error
		if ndb.db, err = gorm.Open("mysql", ndb.cfg.ConnStr); err != nil {
			fmt.Println(err.Error())
			ndb.log.Fatal(err, string(debug.Stack())) // 报错退出程序
			return
		} else {
			ndb.db.DB().SetMaxOpenConns(ndb.cfg.MaxOpenConn)
			ndb.db.DB().SetMaxIdleConns(ndb.cfg.MaxIdleConn)
			ndb.db.DB().SetConnMaxLifetime(time.Second * time.Duration(ndb.cfg.ConnMaxLeftTime))
			ndb.log.Info("db init done. params:", ndb.cfg.ConnStr) // 数据库初始化成功
			ndb.db = ndb.db.Debug()
			ndb.db.LogMode(true)
			ndb.db.SetLogger(ndb.log)
		}
	})
}

func (ndb *NDb) CloseDb() {
	ndb.onceStopDb.Do(func() {
		err := ndb.db.Close()
		if err != nil {
			ndb.log.Warn(err)
		} else {
			ndb.log.Info("db closed.")
		}
	})
}

func NewNDb(dbCfg *ncfg.NConfig, log *nlog.NLog) (ndb *NDb) {
	ctx := context2.NewContext()
	ctx.Module = "ndb"
	ctx.Function = "dbfunc"
	log = log.Get(ctx)
	log.Info("no ndb created")
	ndb = &NDb{log: log}
	ndb.cfg.ConnStr = dbCfg.Cfg.DB.ConnStr
	ndb.cfg.MaxOpenConn = dbCfg.Cfg.DB.MaxOpenConn
	ndb.cfg.MaxIdleConn = dbCfg.Cfg.DB.MaxIdleConn
	ndb.cfg.ConnMaxLeftTime = dbCfg.Cfg.DB.ConnMaxLeftTime
	ndb.Start()
	return ndb
}

package ndb

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	context2 "github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/component/dbmigrate"
	"github.com/krilie/lico_alone/component/ncfg"
	"github.com/krilie/lico_alone/component/nlog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"runtime/debug"
	"strings"
	"sync"
	"time"
)

const gormTransConDb = "gormTransConDb"
const dbVersion = 210001

type NDb struct {
	cfg         ncfg.DB
	log         *nlog.NLog
	onceStartDb sync.Once
	onceStopDb  sync.Once
	db          *gorm.DB
}

func (ndb *NDb) GetDb(ctx context.Context) *gorm.DB {
	values := context2.GetAppValues(ctx)
	if values == nil {
		return ndb.db
	}
	if values.Tx == nil {
		return ndb.db
	} else {
		return values.Tx.(*gorm.DB)
	}
}

// GetSessionDb 获取上下文中的数据库连接 可以为nil
func (ndb *NDb) GetSessionDb(ctx context.Context) *gorm.DB {
	values := context2.GetAppValues(ctx)
	if values == nil {
		return nil
	}
	if values.Tx == nil {
		return nil
	} else {
		return values.Tx.(*gorm.DB)
	}
}

func (ndb *NDb) Ping() error {
	db, err := ndb.db.DB()
	if err != nil {
		panic(err)
	}
	return db.Ping()
}

func (ndb *NDb) Start() {
	ndb.onceStartDb.Do(func() {
		// 数据库迁移
		ndb.MigrationDb()
		// 正常开启数据库
		var err error
		if ndb.db, err = gorm.Open(mysql.Open(ndb.cfg.ConnStr), &gorm.Config{}); err != nil {
			fmt.Println(err.Error())
			ndb.log.Fatal(err, string(debug.Stack())) // 报错退出程序
			return
		} else {
			db, err := ndb.db.DB()
			if err != nil {
				panic(err)
				return
			}
			db.SetMaxOpenConns(ndb.cfg.MaxOpenConn)
			db.SetMaxIdleConns(ndb.cfg.MaxIdleConn)
			db.SetConnMaxLifetime(time.Second * time.Duration(ndb.cfg.ConnMaxLeftTime))
			ndb.log.Info("migrationDb init done. params:", "connect string") // 数据库初始化成功
			ndb.db.Logger = &ndbLogger{NLog: ndb.log.WithField("gorm", "gorm-inner")}
			ndb.db.Logger.LogMode(logger.Info)
			ndb.db = ndb.db.Debug()
		}
	})
}

func (ndb *NDb) CloseDb() {
	ndb.onceStopDb.Do(func() {
		db, err2 := ndb.db.DB()
		if err2 != nil {
			panic(err2)
		}
		err := db.Close()
		if err != nil {
			ndb.log.Warn(err)
		} else {
			ndb.log.Info("db closed.")
		}
	})
}

func NewNDb(cfg *ncfg.NConfig, log *nlog.NLog) (ndb *NDb) {

	var dbCfg = cfg.GetDbCfg()

	values := context2.NewAppCtxValues()
	values.Module = "ndb"
	values.Function = "NewNDb"
	ctx := context2.NewAppCtx(context.Background(), values)
	log = log.Get(ctx)
	log.Info("no ndb created")
	ndb = &NDb{log: log}
	ndb.cfg = *dbCfg
	if ndb.cfg.MigrationPath == "" {
		ndb.cfg.MigrationPath = "/migrations"
	}
	ndb.Start()
	return ndb
}

func (ndb *NDb) MigrationDb() {
	defer func() {
		if err := recover(); err != nil {
			ndb.log.WithField("err", err).WithField("db_version", dbVersion).WithField("migration_path", ndb.cfg.MigrationPath).Info("migrations failure")
			panic(err)
		}
	}()
	// 数据库迁移
	var dbName = GetDbNameFromConnectStr(ndb.cfg.ConnStr)
	var connectStrForMigration = strings.Replace(ndb.cfg.ConnStr, dbName, "", 1)
	migrationDb, err := sql.Open("mysql", connectStrForMigration)
	if err != nil {
		panic(err)
	}
	defer migrationDb.Close()
	_, err = migrationDb.Exec("CREATE DATABASE IF NOT EXISTS " + dbName + " DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci;")
	if err != nil {
		panic(err)
	}
	_, err = migrationDb.Exec("USE " + dbName)
	if err != nil {
		panic(err)
	}
	if ndb.cfg.MigrationPath == "" {
		ndb.cfg.MigrationPath = "/migrations"
	}
	// 如果没有则创建数据库
	dbmigrate.Migrate(migrationDb, "file://"+ndb.cfg.MigrationPath, dbVersion) // 指定数据库版本
	ndb.log.WithField("db_name", dbName).WithField("db_version", dbVersion).WithField("migration_path", ndb.cfg.MigrationPath).Info("migrations ok")
}

func GetDbNameFromConnectStr(connectStr string) (dbName string) {
	begin := strings.Index(connectStr, "/")
	if begin == -1 {
		return ""
	}
	end := strings.Index(connectStr, "?")
	if end == -1 {
		return ""
	}
	dbName = connectStr[begin+1 : end]
	return dbName
}

package ndb

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/krilie/lico_alone/component/dbmigrate"
	"testing"
)

func TestMigrate3(t *testing.T) {
	db, err := sql.Open("mysql", "root:123123@tcp(local:3306)/myapp_test?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai&multiStatements=true")
	if err != nil {
		panic(err)
	}
	dbmigrate.Migrate(db, "test", "file://c://sqls", 20210206140300)
}

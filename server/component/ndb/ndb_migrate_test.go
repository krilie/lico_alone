package ndb

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/krilie/lico_alone/component/dbmigrate"
	"testing"
)

func TestMigrate3(t *testing.T) {
	db, err := sql.Open("mysql", "test:123456@tcp(lizo.top:3306)/?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai&multiStatements=true")
	if err != nil {
		panic(err)
	}
	dbmigrate.Migrate(db, "file://../dbmigrate/migrations", 20210206140300)
}

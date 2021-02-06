package dbmigrate

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(db *sql.DB, dbName string, path string, version uint) {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		path, // "file://component/dbmigrate/migrations", // 相对目录
		dbName, driver)
	if err != nil {
		panic(err)
	}

	err = m.Migrate(version) // <0 down,=0 noop,>0 up
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic(err)
	}
}

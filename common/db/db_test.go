package db

import (
	"testing"
)

func TestDb(T *testing.T) {
	defer Close()
	db := Db.New()
	T.Log(db)
	_ = db.Close()
}

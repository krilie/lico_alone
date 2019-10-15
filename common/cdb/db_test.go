package cdb

import (
	"testing"
)

func TestDb(T *testing.T) {
	defer CloseDb()
	db := Db.New()
	T.Log(db)
	_ = db.Close()
}

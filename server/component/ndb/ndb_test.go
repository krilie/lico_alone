package ndb

import "testing"

func TestNDb_Start(t *testing.T) {
	db, f, err := InitNDb()
	t.Log(db, err)
	f()
}

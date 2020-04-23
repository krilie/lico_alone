// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package ndb

import (
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/component/nlog"
)

// Injectors from wire.go:

func InitNDb() (*NDb, func(), error) {
	config := _wireConfigValue
	db := config.DB
	nLog := nlog.InitNLog()
	nDb, cleanup := NewNDb(db, nLog)
	return nDb, func() {
		cleanup()
	}, nil
}

var (
	_wireConfigValue = config.Cfg
)

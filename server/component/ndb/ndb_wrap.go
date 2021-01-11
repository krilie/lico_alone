package ndb

import (
	"context"
)

func (ndb *NDb) Exec(ctx context.Context, sql string, values ...interface{}) (affected int64, err error) {
	resultDb := ndb.GetDb(ctx).Exec(sql, values...)
	return resultDb.RowsAffected, resultDb.Error
}

func (ndb *NDb) RawQuery(ctx context.Context, outData interface{}, sql string, values ...interface{}) error {
	raw := ndb.GetDb(ctx).Raw(sql, values...)
	err := raw.Error
	if err != nil {
		return err
	}
	raw.Scan(outData)
	return nil
}

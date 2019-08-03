package cdb


// 事务
type CDbTx interface {
	Commit() error
	RollBack() error
}

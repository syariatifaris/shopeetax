package db

import (
	"context"
	"github.com/jmoiron/sqlx"
)

//ExecuteTransactionFunc contract
type ExecuteTransactionFunc func(context.Context, *sqlx.Tx, func() error) error

//RepoDB base repository database contract
type RepoDB interface {
	BeginTransaction() (*sqlx.Tx, error)
}

//ExecuteInTx executes in transaction
func ExecuteInTx(ctx context.Context, tx *sqlx.Tx, fn func() error) (err error) {
	defer func() {
		if err == nil {
			_ = tx.Commit()
		} else {
			_ = tx.Rollback()
		}
	}()
	err = fn()
	return err
}

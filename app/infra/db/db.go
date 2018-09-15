package db

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/syariatifaris/shopeetax/app/infra/config"
)

const dcsFormatPG = `user=%s password=%s host=%s dbname=%s port=%d sslmode=disable`
const dcsFormatMySQL = `%s:%s@(%s:%d)/%s`

//NsqSqlxConnection createa new sqlx db connection
func NewSqlxConnection(cfg config.Database) *sqlx.DB {
	connectionString := fmt.Sprintf(dcsFormatMySQL,
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)
	db, err := sqlx.Connect(cfg.Type, connectionString)
	if err != nil {
		panic(fmt.Sprint("cannot create db connection ", err.Error()))
		return nil
	}

	return db
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

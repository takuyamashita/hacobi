package repository

import (
	"context"
	"database/sql"
)

type Transaction struct {
	db *sql.DB
}

func (t Transaction) Exec(f func(*sql.Tx) error, ctx context.Context) error {

	opt := &sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  false,
	}

	tx, err := t.db.BeginTx(ctx, opt)
	if err != nil {
		return err
	}

	if err := f(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

package repository

import (
	"context"
	"database/sql"

	"github.com/takuyamashita/hacobi/src/api/pkg/db"
)

type Transaction struct {
	db *db.MySQL
}

func (t Transaction) Begin(ctx context.Context) (func() error, error) {

	opt := &sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  false,
	}

	tx, err := t.db.BeginTx(ctx, opt)
	if err != nil {
		return nil, err
	}

	return tx.Rollback, nil

}

package repository

import (
	"context"
	"database/sql"

	"github.com/takuyamashita/hacobi/src/api/pkg/db"
	"github.com/takuyamashita/hacobi/src/api/pkg/usecase"
)

var _ usecase.TransationRepositoryIntf = (*Transaction)(nil)

type Transaction struct {
	db *db.MySQL
}

func NewTransaction(db *db.MySQL) *Transaction {
	return &Transaction{
		db: db,
	}
}

func makeCommitFunc(tx *db.MySQL) func() error {
	return func() error {
		return tx.Commit()
	}
}

func (t Transaction) Begin(ctx context.Context) (func() error, func() error, error) {

	opt := &sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  false,
	}

	tx, err := t.db.BeginTx(ctx, opt)
	if err != nil {
		return nil, nil, err
	}

	return makeCommitFunc(tx), tx.Rollback, nil

}

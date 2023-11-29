package db

import (
	"context"
	"database/sql"
	"errors"
	"log"
)

type MySQL struct {
	*sql.DB
	tx       *sql.Tx
	txCount  int
	Begin    interface{}
	Exec     interface{}
	Query    interface{}
	QueryRow interface{}
}

func NewMySQL(db *sql.DB) MySQL {
	return MySQL{
		DB: db,
	}
}

type txFunc func(tx *MySQL) error

func (m *MySQL) BeginTx(ctx context.Context, opts *sql.TxOptions) (*MySQL, error) {

	log.Println("Begin", m.txCount)

	m.txCount++

	if m.tx != nil {
		return m, nil
	}

	tx, err := m.DB.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}

	m.tx = tx

	return m, nil
}

func (m *MySQL) Commit() error {

	log.Println("Commit", m.txCount)

	if m.txCount > 1 {
		m.txCount--
		return nil
	}

	if m.tx == nil {
		return errors.New("transaction has already been committed")
	}

	err := m.tx.Commit()
	if err != nil {
		return err
	}

	m.txCount = 0
	m.tx = nil

	return nil
}

func (m *MySQL) Rollback() error {

	log.Println("Rollback", m.txCount)

	if m.tx != nil {
		m.txCount = 0
		err := m.tx.Rollback()
		m.tx = nil
		return err
	}

	return nil
}

func (m *MySQL) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {

	if m.tx != nil {
		return m.tx.ExecContext(ctx, query, args...)
	}

	return m.DB.ExecContext(ctx, query, args...)
}

func (m *MySQL) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {

	if m.tx != nil {
		return m.tx.QueryContext(ctx, query, args...)
	}

	return m.DB.QueryContext(ctx, query, args...)
}

func (m *MySQL) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {

	if m.tx != nil {
		return m.tx.QueryRowContext(ctx, query, args...)
	}

	return m.DB.QueryRowContext(ctx, query, args...)
}

package daohelper

import (
	"context"
	"database/sql"
	golog "log"

	"github.com/doug-martin/goqu/v9"
	"github.com/jtorz/phoenix-backend/app/shared/base"
	"github.com/jtorz/phoenix-backend/app/shared/ctxinfo"
)

// QueryHelper holds helper functions.
type QueryHelper struct{}

// NewSelect initiates the sql query.
func (QueryHelper) NewSelect(table ...interface{}) *goqu.SelectDataset {
	return goqu.Dialect("postgres").From(table...)
}

// NewInsert initiates the sql insert statement.
func (QueryHelper) NewInsert(table interface{}) *goqu.InsertDataset {
	return goqu.Dialect("postgres").Insert(table)
}

// NewUpdate initiates the sql update statement.
func (QueryHelper) NewUpdate(table interface{}) *goqu.UpdateDataset {
	return goqu.Dialect("postgres").Update(table)
}

// NewDelete initiates the sql delete statement.
func (QueryHelper) NewDelete(table interface{}) *goqu.DeleteDataset {
	return goqu.Dialect("postgres").Delete(table)
}

// QueryContext executes a query that returns rows, typically a SELECT.
func (QueryHelper) QueryContext(ctx context.Context, exe base.Executor, s *goqu.SelectDataset) (*sql.Rows, error) {
	query, args, err := s.Prepared(true).ToSQL()
	if err != nil {
		return nil, err
	}
	if ctxinfo.PrintLog(ctx) {
		golog.Println(query, "\n", args)
	}
	return exe.QueryContext(ctx, query, args...)
}

// QueryRowContext executes a query that is expected to return at most one row.
// If the query selects no rows, the *Row's Scan will return ErrNoRows.
// Otherwise, the *Row's Scan scans the first selected row and discards
// the rest.
func (QueryHelper) QueryRowContext(ctx context.Context, exe base.Executor, s *goqu.SelectDataset) (*sql.Row, error) {
	query, args, err := s.Prepared(true).ToSQL()
	if err != nil {
		return nil, err
	}
	if ctxinfo.PrintLog(ctx) {
		golog.Println(query, "\n", args)
	}
	return exe.QueryRowContext(ctx, query, args...), nil
}

// DoUpdate executes a query without returning any rows.
func (QueryHelper) DoUpdate(ctx context.Context, exe base.Executor, s *goqu.UpdateDataset) (sql.Result, error) {
	query, args, err := s.Prepared(true).ToSQL()
	if err != nil {
		return nil, err
	}
	if ctxinfo.PrintLog(ctx) {
		golog.Println(query, "\n", args)
	}
	return exe.ExecContext(ctx, query, args...)
}

// DoUpdateReturningRow executes a query  returning a row.
func (QueryHelper) DoUpdateReturningRow(ctx context.Context, exe base.Executor, s *goqu.UpdateDataset, returning ...interface{}) (*sql.Row, error) {
	query, args, err := s.Prepared(true).Returning(returning...).ToSQL()
	if err != nil {
		return nil, err
	}
	if ctxinfo.PrintLog(ctx) {
		golog.Println(query, "\n", args)
	}
	return exe.QueryRowContext(ctx, query, args...), nil
}

// DoInsert executes a query without returning any rows.
func (QueryHelper) DoInsert(ctx context.Context, exe base.Executor, s *goqu.InsertDataset) (sql.Result, error) {
	query, args, err := s.Prepared(true).ToSQL()
	if err != nil {
		return nil, err
	}
	if ctxinfo.PrintLog(ctx) {
		golog.Println(query, "\n", args)
	}
	return exe.ExecContext(ctx, query, args...)
}

// DoDelete executes a query without returning any rows.
func (QueryHelper) DoDelete(ctx context.Context, exe base.Executor, s *goqu.DeleteDataset) (sql.Result, error) {
	query, args, err := s.Prepared(true).ToSQL()
	if err != nil {
		return nil, err
	}
	if ctxinfo.PrintLog(ctx) {
		golog.Println(query, "\n", args)
	}
	return exe.ExecContext(ctx, query, args...)
}

// DoInsertReturning executes a query without returning any rows.
func (QueryHelper) DoInsertReturning(ctx context.Context, exe base.Executor, s *goqu.InsertDataset, returning ...interface{}) (*sql.Row, error) {
	query, args, err := s.Prepared(true).Returning(returning...).ToSQL()
	if err != nil {
		return nil, err
	}
	if ctxinfo.PrintLog(ctx) {
		golog.Println(query, "\n", args)
	}
	return exe.QueryRowContext(ctx, query, args...), nil
}

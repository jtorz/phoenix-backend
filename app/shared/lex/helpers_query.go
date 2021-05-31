package lex

import (
	"context"
	"database/sql"
	"log"

	"github.com/doug-martin/goqu/v9"
	"github.com/jtorz/phoenix-backend/app/config"
	"github.com/jtorz/phoenix-backend/app/shared/base"
	"github.com/jtorz/phoenix-backend/app/shared/ctxinfo"

	// goqudialect
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/doug-martin/goqu/v9/exp"
)

// CoalesceStr coalesce empty string.
func CoalesceStr(col string) exp.SQLFunctionExpression {
	return goqu.COALESCE(goqu.C(col), goqu.L("''"))
}

// NewSelect initiates the sql query.
func NewSelect(columns ...interface{}) *goqu.SelectDataset {
	return goqu.Dialect("postgres").Select(columns...)
}

// NewInsert initiates the sql insert statement.
func NewInsert(table interface{}) *goqu.InsertDataset {
	return goqu.Dialect("postgres").Insert(table)
}

// NewUpdate initiates the sql update statement.
func NewUpdate(table interface{}) *goqu.UpdateDataset {
	return goqu.Dialect("postgres").Update(table)
}

// NewDelete initiates the sql delete statement.
func NewDelete(table interface{}) *goqu.DeleteDataset {
	return goqu.Dialect("postgres").Delete(table)
}

// QueryContext executes a query that returns rows, typically a SELECT.
func QueryContext(ctx context.Context, exe base.Executor, s *goqu.SelectDataset) (*sql.Rows, error) {
	if err := s.Error(); err != nil {
		return nil, err
	}
	query, args, err := s.Prepared(true).ToSQL()
	if err != nil {
		return nil, err
	}
	if ctxinfo.LogginAllowed(ctx, config.LogDebug) {
		log.Println(query, "\n", args)
	}
	return exe.QueryContext(ctx, query, args...)
}

// QueryRowContext executes a query that is expected to return at most one row.
// If the query selects no rows, the *Row's Scan will return ErrNoRows.
// Otherwise, the *Row's Scan scans the first selected row and discards
// the rest.
func QueryRowContext(ctx context.Context, exe base.Executor, s *goqu.SelectDataset) (*sql.Row, error) {
	if err := s.Error(); err != nil {
		return nil, err
	}
	query, args, err := s.Prepared(true).ToSQL()
	if err != nil {
		return nil, err
	}
	if ctxinfo.LogginAllowed(ctx, config.LogDebug) {
		log.Println(query, "\n", args)
	}
	return exe.QueryRowContext(ctx, query, args...), nil
}

// DoUpdate executes a query without returning any rows.
func DoUpdate(ctx context.Context, exe base.Executor, s *goqu.UpdateDataset) (sql.Result, error) {
	if err := s.Error(); err != nil {
		return nil, err
	}
	query, args, err := s.Prepared(true).ToSQL()
	if err != nil {
		return nil, err
	}
	if ctxinfo.LogginAllowed(ctx, config.LogDebug) {
		log.Println(query, "\n", args)
	}
	return exe.ExecContext(ctx, query, args...)
}

// DoUpdateReturningRow executes a query  returning a row.
func DoUpdateReturningRow(ctx context.Context, exe base.Executor, s *goqu.UpdateDataset, returning ...interface{}) (*sql.Row, error) {
	if err := s.Error(); err != nil {
		return nil, err
	}
	query, args, err := s.Prepared(true).Returning(returning...).ToSQL()
	if err != nil {
		return nil, err
	}
	if ctxinfo.LogginAllowed(ctx, config.LogDebug) {
		log.Println(query, "\n", args)
	}
	return exe.QueryRowContext(ctx, query, args...), nil
}

// DoInsert executes a query without returning any rows.
func DoInsert(ctx context.Context, exe base.Executor, s *goqu.InsertDataset) (sql.Result, error) {
	if err := s.Error(); err != nil {
		return nil, err
	}
	query, args, err := s.Prepared(true).ToSQL()
	if err != nil {
		return nil, err
	}
	if ctxinfo.LogginAllowed(ctx, config.LogDebug) {
		log.Println(query, "\n", args)
	}
	return exe.ExecContext(ctx, query, args...)
}

// DoDelete executes a query without returning any rows.
func DoDelete(ctx context.Context, exe base.Executor, s *goqu.DeleteDataset) (sql.Result, error) {
	if err := s.Error(); err != nil {
		return nil, err
	}
	query, args, err := s.Prepared(true).ToSQL()
	if err != nil {
		return nil, err
	}
	if ctxinfo.LogginAllowed(ctx, config.LogDebug) {
		log.Println(query, "\n", args)
	}
	return exe.ExecContext(ctx, query, args...)
}

// DoInsertReturning executes a query without returning any rows.
func DoInsertReturning(ctx context.Context, exe base.Executor, s *goqu.InsertDataset, returning ...interface{}) (*sql.Row, error) {
	if err := s.Error(); err != nil {
		return nil, err
	}
	query, args, err := s.Prepared(true).Returning(returning...).ToSQL()
	if err != nil {
		return nil, err
	}
	if ctxinfo.LogginAllowed(ctx, config.LogDebug) {
		log.Println(query, "\n", args)
	}
	return exe.QueryRowContext(ctx, query, args...), nil
}

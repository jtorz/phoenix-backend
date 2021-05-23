package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/doug-martin/goqu/v9"
	"github.com/jtorz/phoenix-backend/app/shared/base"
)

// QueryContext executes a query that returns rows, typically a SELECT.
func QueryContext(ctx context.Context, exe base.Executor, s *goqu.SelectDataset) (*sql.Rows, error) {
	if err := s.Error(); err != nil {
		return nil, err
	}
	query, args, err := s.Prepared(true).ToSQL()
	if err != nil {
		return nil, err
	}
	log.Println(query, "\n", args)
	return exe.QueryContext(ctx, query, args...)
}

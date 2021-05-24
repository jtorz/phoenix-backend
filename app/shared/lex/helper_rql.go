package lex

import (

	//"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/jtorz/phoenix-backend/app/shared/base"
	rqlgq "github.com/jtorz/phoenix-backend/utils/rql-goqu"
)

func parser(model interface{}) (*rqlgq.Parser, error) {
	return rqlgq.NewParser(rqlgq.Config{Model: model, ColumnFn: func(s string) string { return s }})
}

// ParseClientFilter parser the client query for the given table.
func ParseClientFilter(qry base.ClientQuery, model interface{}) (rqlParams *rqlgq.Params, err error) {
	if len(qry.RQL) == 0 {
		return &rqlgq.Params{
			FilterExp: exp.NewExpressionList(exp.AndType),
		}, nil
	}
	p, err := parser(model)
	if err != nil {
		return nil, err
	}
	return p.Parse(qry.RQL)
}

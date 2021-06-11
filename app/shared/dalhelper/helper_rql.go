package dalhelper

import (

	//"github.com/doug-martin/goqu/v9"
	"fmt"

	"github.com/doug-martin/goqu/v9/exp"
	"github.com/jtorz/phoenix-backend/app/shared/base"
	"github.com/jtorz/phoenix-backend/app/shared/baseerrors"
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
		return nil, fmt.Errorf("%w %s", baseerrors.ErrInvalidData, err.Error())
	}
	rqlParams, err = p.Parse(qry.RQL)
	if err != nil {
		return nil, fmt.Errorf("%w %s", baseerrors.ErrInvalidData, err.Error())
	}
	return
}

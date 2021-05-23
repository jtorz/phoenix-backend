// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package {{.PackageName}}

import (
	"fmt"

	//"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/jtorz/phoenix-backend/app/shared/base"
	rqlgq "github.com/jtorz/phoenix-backend/utils/rql-goqu"
)

var rql map[string]*rqlgq.Parser

func init() {
	rql = make(map[string]*rqlgq.Parser)

	{{range .Tables}}
	initParser(T.{{.GoCase}}, {{.GoCase}}) {{end}}

	{{range .Views}}
	initParser(V.{{.GoCase}}, {{.GoCase}}) {{end}}
}

func initParser(tablename string, model interface{}) {
	p, err := rqlgq.NewParser(rqlgq.Config{Model: model})
	if err != nil {
		panic(fmt.Sprintf("can't init rql config %s", err.Error()))
	}
	rql[tablename] = p
}

{{range $Obj := .Tables}}
	// ParseFilter parser the client query for the table {{.Name}}.
	func (Table{{$Obj.GoCase}}) ParseFilter(qry base.ClientQuery) (rqlParams *rqlgq.Params, err error) {
		return ParseFilter(T.{{$Obj.GoCase}}, qry)
	}
{{end}}

{{range $Obj := .Views}}
	// ParseFilter parser the client query for the table {{.Name}}.
	func (View{{$Obj.GoCase}}) ParseFilter(qry base.ClientQuery) (rqlParams *rqlgq.Params, err error) {
		return ParseFilter(V.{{$Obj.GoCase}}, qry)
	}
{{end}}


// ParseFilter parser the client query for the given table.
func ParseFilter(tablename string, qry base.ClientQuery) (rqlParams *rqlgq.Params, err error) {
	if len(qry.RQL) == 0 {
		return &rqlgq.Params{
			FilterExp: exp.NewExpressionList(exp.AndType),
		}, nil
	}
	p, err := rql[tablename].Parse(qry.RQL)
	if err != nil {
		return nil, err
	}
	return p, nil
}
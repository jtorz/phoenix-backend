package main

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/jtorz/phoenix-backend/app/shared/daohelper"
	"github.com/jtorz/phoenix-backend/utils/stringset"
	"github.com/volatiletech/sqlboiler/strmangle"
)

func (g *Generator) getObjects() (err error) {
	if g.TPLData.Tables, err = g.getDbObjects("tablename", "pg_tables"); err != nil {
		return err
	}
	if g.TPLData.Views, err = g.getDbObjects("viewname", "pg_views"); err != nil {
		return err
	}
	return nil
}

func (g *Generator) getDbObjects(objname, objtype string) ([]TplObject, error) {
	var objects []TplObject
	h := daohelper.QueryHelper{}
	where := []exp.Expression{goqu.C("schemaname").Eq(g.Schema)}
	if g.FilterPrefix != "" {
		where = append(where, goqu.C(objname).Like(g.FilterPrefix+"%"))
	}
	query := h.NewSelect(objtype).Select(objname).Where(where...)
	rows, err := h.QueryContext(context.Background(), g.DB, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		obj := TplObject{}
		if err := rows.Scan(&obj.Name); err != nil {
			return nil, err
		}
		obj.GoCase = stringset.UpperFirst(strmangle.CamelCase(obj.Name))

		obj.Columns, err = g.getDbObjectColumns(obj.Name)
		if err != nil {
			return nil, err
		}
		objects = append(objects, obj)
	}
	return objects, nil
}

func (g *Generator) getDbObjectColumns(objname string) ([]TplColumn, error) {
	var columns []TplColumn
	h := daohelper.QueryHelper{}
	query := h.NewSelect("information_schema.columns").Select("column_name").
		Where(
			goqu.C("table_schema").Eq(g.Schema),
			goqu.C("table_name").Eq(objname),
		)
	rows, err := h.QueryContext(context.Background(), g.DB, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		col := TplColumn{}
		if err := rows.Scan(&col.Name); err != nil {
			return nil, err
		}
		col.GoCase = stringset.UpperFirst(strmangle.CamelCase(col.Name))
		columns = append(columns, col)
	}
	return columns, nil
}

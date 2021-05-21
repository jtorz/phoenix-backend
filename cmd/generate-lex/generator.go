package main

import (
	"context"
	"fmt"
	"strings"

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
		obj.GoCase = goCase(obj.Name)

		obj.Columns, err = g.getDbObjectColumns(obj.Name)
		if err != nil {
			return nil, err
		}

		obj.Fks, err = g.getFK(obj.Name)
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
	query := h.NewSelect("information_schema.columns").
		Select("column_name", "is_nullable", "data_type").
		Where(
			goqu.C("table_schema").Table("columns").Eq(g.Schema),
			goqu.C("table_name").Table("columns").Eq(objname),
		)

	rows, err := h.QueryContext(context.Background(), g.DB, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		col := TplColumn{Nullable: "-"}
		var nullable string
		if err := rows.Scan(&col.Name, &nullable, &col.DataType); err != nil {
			return nil, err
		}
		if nullable == "YES" {
			col.Nullable = "N"
		}
		col.DataType = strings.ReplaceAll(col.DataType, `"`, "")
		col.DataType = strings.ReplaceAll(col.DataType, "`", "")
		col.DataType = strings.ReplaceAll(col.DataType, ",", "")
		col.GoCase = stringset.UpperFirst(strmangle.CamelCase(col.Name))
		columns = append(columns, col)
	}
	return columns, nil
}

func (g *Generator) getFK(objname string) ([]Fk, error) {
	var fks []Fk
	qry := `SELECT
		tc.constraint_name,
		string_agg(distinct kcu.column_name, ', ') AS column_names,
		ccu.table_name AS foreign_table_name,
		string_agg(distinct ccu.column_name, ', ') AS foreign_column_names
	FROM
		information_schema.table_constraints AS tc
		JOIN information_schema.key_column_usage AS kcu
		ON tc.constraint_name = kcu.constraint_name
		JOIN information_schema.constraint_column_usage AS ccu
		ON ccu.constraint_name = tc.constraint_name
	WHERE constraint_type = 'FOREIGN KEY'
		AND tc.table_name = $1
	GROUP BY tc.constraint_name, tc.table_name, ccu.table_name;`
	rows, err := g.DB.Query(qry, objname)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		fk := Fk{}
		var columnNames, fcolumnNames string
		if err := rows.Scan(&fk.ConstraintName, &columnNames, &fk.FTable, &fcolumnNames); err != nil {
			return nil, err
		}
		fk.ConstraintName = goCase(fk.ConstraintName)
		fk.FTable = goCase(fk.FTable)
		columns := strings.Split(columnNames, ",")
		fcolumns := strings.Split(fcolumnNames, ",")
		fk.Columns = make([]FKCol, len(columns))
		for i := range columns {
			fk.Columns[i] = FKCol{
				Orig: goCase(columns[i]),
				Dest: goCase(fcolumns[i]),
			}
			fmt.Printf("%#v\n", fk.Columns[i])
		}
		fks = append(fks, fk)
	}
	return fks, nil
}

func goCase(s string) string {
	return stringset.UpperFirst(strmangle.CamelCase(strings.TrimSpace(s)))
}

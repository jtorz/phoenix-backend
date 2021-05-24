package codegen

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

type Entity struct {
	DBName   string // database table name (Examples: fnd_user, fnd_role)
	DBGoCase string // database table name in Go Case (Examples: FndUser, FndRole)
	Columns  []Attribute
	Fks      []Fk
}

type Attribute struct {
	DBName     string // database column name (Examples: use_id, use_name)
	DBGoCase   string // database column name in Go case (Examples: UseID, UseName)
	DBNullable bool   // database value is nullable
	DBDataType string // database data type
	GoField    string // Go field name (Examples: ID, Name)
	GoDataType string // Go field name (Examples: ID, Name)
}

type Fk struct {
	DBName       string // database fk constraint name in Go Case (Examples: FndActionFkFndModule)
	DBGoCase     string // database fk constraint name (Examples: fnd_user, fnd_action_fk_fnd_module)
	FTableGoCase string // database foreing table name in Go Case (Examples: FndModule)
	FTable       string // database foreing table name (Examples: fnd_module)
	Columns      []FKColumns
}

type FKColumns struct {
	O Attribute // origin
	D Attribute // destiny
}

func GetEntities(ctx context.Context, db *sql.DB, schema, filterprefix string) (tables, views []Entity, err error) {
	tables, err = GetEntitiesFromType(ctx, db, TypeTables, schema, filterprefix)
	if err != nil {
		return nil, nil, err
	}
	views, err = GetEntitiesFromType(ctx, db, TypeViews, schema, filterprefix)
	if err != nil {
		return nil, nil, err
	}
	return
}

type EntityType string

const (
	TypeTables EntityType = "T"
	TypeViews  EntityType = "V"
)

func GetEntitiesFromType(ctx context.Context, db *sql.DB, entType EntityType, schema, filterprefix string) ([]Entity, error) {
	var entityType, datasource string
	switch entType {
	case TypeTables:
		entityType = "tablename"
		datasource = "pg_tables"
	case TypeViews:
		entityType = "viewname"
		datasource = "pg_views"
	}

	var entities []Entity
	where := []exp.Expression{goqu.C("schemaname").Eq(schema)}
	if filterprefix != "" {
		where = append(where, goqu.C(entityType).Like(filterprefix+"%"))
	}
	query := goqu.Dialect("postgres").Select(entityType).From(datasource).Where(where...).Order(goqu.I(entityType).Asc())
	rows, err := QueryContext(context.Background(), db, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		obj := Entity{}
		if err := rows.Scan(&obj.DBName); err != nil {
			return nil, err
		}
		obj.DBGoCase = goCase(obj.DBName)

		obj.Columns, err = GetAttributes(ctx, db, schema, obj.DBName)
		if err != nil {
			return nil, err
		}

		obj.Fks, err = GetFK(ctx, db, obj.DBName)
		if err != nil {
			return nil, err
		}
		entities = append(entities, obj)
	}
	return entities, nil
}

func GetAttributes(ctx context.Context, db *sql.DB, schema, objname string) ([]Attribute, error) {
	var columns []Attribute
	query := goqu.Dialect("postgres").Select("column_name", "is_nullable", "data_type", "domain_name").
		From("information_schema.columns").
		Where(
			goqu.C("table_schema").Table("columns").Eq(schema),
			goqu.C("table_name").Table("columns").Eq(objname),
		).
		Order(goqu.I("ordinal_position").Asc())

	rows, err := QueryContext(ctx, db, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		col := Attribute{}
		var dm sql.NullString
		var nullable string
		err := rows.Scan(
			&col.DBName,
			&nullable,
			&col.DBDataType,
			&dm,
		)
		if err != nil {
			return nil, err
		}
		col.DBNullable = nullable == "YES"
		col.DBGoCase = goCase(col.DBName)
		col.DBDataType = strings.ReplaceAll(col.DBDataType, `"`, "")
		col.DBDataType = strings.ReplaceAll(col.DBDataType, "`", "")
		col.DBDataType = strings.ReplaceAll(col.DBDataType, ",", " ")
		col.GoDataType = TranslateColumnType(col.DBDataType)
		col.GoField = col.DBGoCase[3:]

		columns = append(columns, col)
	}
	return columns, nil
}

func GetFK(ctx context.Context, db *sql.DB, objname string) ([]Fk, error) {
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
	rows, err := db.QueryContext(ctx, qry, objname)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		fk := Fk{}
		var columnNames, fcolumnNames string
		if err := rows.Scan(&fk.DBName, &columnNames, &fk.FTable, &fcolumnNames); err != nil {
			return nil, err
		}
		fk.DBGoCase = goCase(fk.DBName)
		fk.FTableGoCase = goCase(fk.FTable)
		columns := strings.Split(columnNames, ",")
		fcolumns := strings.Split(fcolumnNames, ",")
		fk.Columns = make([]FKColumns, len(columns))
		for i := range columns {
			fk.Columns[i] = FKColumns{
				O: Attribute{
					DBName:   columns[i],
					DBGoCase: goCase(columns[i]),
				},
				D: Attribute{
					DBName:   fcolumns[i],
					DBGoCase: goCase(fcolumns[i]),
				},
			}
			fmt.Printf("%#v\n", fk.Columns[i])
		}
		fks = append(fks, fk)
	}
	return fks, nil
}

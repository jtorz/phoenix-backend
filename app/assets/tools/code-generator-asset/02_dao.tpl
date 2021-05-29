package {{$.ServiceAbbr | lowercase}}dao

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/jtorz/phoenix-backend/app/services/{{$.ServiceAbbr | lowercase}}/{{$.ServiceAbbr | lowercase}}model"
	"github.com/jtorz/phoenix-backend/app/shared/base"
	"github.com/jtorz/phoenix-backend/app/shared/baseerrors"

	//lint:ignore ST1001 dot import allowed only in dao packages for
	. "github.com/jtorz/phoenix-backend/app/shared/lex"
)

// Dao{{$.Entity.GoStruct}} Data Access structure.
type Dao{{$.Entity.GoStruct}} struct{}

// GetByID retrives the record information using its ID.
func (dao *Dao{{$.Entity.GoStruct}}) GetByID(ctx context.Context, exe base.Executor,
	{{range $Col := $.Entity.Columns}}
	{{- if $Col.IsPK}} {{$Col.GoVarName}} {{$Col.GoDataType}}, {{end}}
	{{- end}}
) (*{{$.ServiceAbbr | lowercase}}model.{{$.Entity.GoStruct}}, error) {
	rec := {{$.ServiceAbbr | lowercase}}model.{{$.Entity.GoStruct}}{}
	query := NewSelect(
			{{- range $Col := $.Entity.Columns}}
			{{- if not $Col.IsPK}}
				{{$.Entity.DBGoCase}}.{{$Col.DBGoCase}},
			{{- end}}
			{{- end}}
		).
		From(goqu.T(T.{{$.Entity.DBGoCase}})).
		Where(
			{{- range $Col := $.Entity.Columns}}
			{{- if $Col.IsPK}}
				goqu.C({{$.Entity.DBGoCase}}.{{$Col.DBGoCase}}).Eq({{$Col.GoVarName}}),
			{{- end}}
			{{- end}}
		)

	row, err := QueryRowContext(ctx, exe, query)
	if err != nil {
		DebugErr(ctx, err)
		return nil, err
	}

	{{- range $Col := $.Entity.Columns}}
	{{- if $Col.DBNullable}}
		var {{$Col.GoVarName}} Zero{{$Col.GoDataType | upperfirst}}
	{{- end}}
	{{- end}}
	err = row.Scan(
		{{- range $Col := $.Entity.Columns}}
		{{- if not $Col.IsPK}}
			{{- if $Col.DBNullable}}
				&{{$Col.GoVarName}},
			{{- else}}
					&rec.{{$Col.GoField}},
			{{- end}}
		{{- end}}
		{{- end}}
	)
	if err != nil {
		DebugErr(ctx, err)
		return nil, err
	}

	{{- range $Col := $.Entity.Columns}}
	{{- if $Col.DBNullable}}
		if {{$Col.GoVarName}} != "" {
			rec.{{$Col.GoField}} = {{$Col.GoDataType}}({{$Col.GoVarName}} )
		}
	{{- end}}
	{{- end}}
	{{range $Col := $.Entity.Columns}}
	{{- if $Col.IsPK}}
		rec.{{$Col.GoField}} = {{$Col.GoVarName}}
	{{- end}}
	{{- end}}
	return &rec, nil
}

// List returns the list of records that can be filtered by the user.
func (dao *Dao{{$.Entity.GoStruct}}) List(ctx context.Context, exe base.Executor,
	qry base.ClientQuery,
) ({{$.ServiceAbbr | lowercase}}model.{{$.Entity.GoSlice}}, error) {
	res := make({{$.ServiceAbbr | lowercase}}model.{{$.Entity.GoSlice}}, 0)
	params, err := ParseClientFilter(qry, {{$.ServiceAbbr | lowercase}}model.{{$.Entity.GoStruct}}{})
	if err != nil {
		return nil, err
	}
	if qry.OnlyActive {
		{{- range $Col := $.Entity.Columns}}
		{{- if eq $Col.GoField "Status"}}
			params.FilterExp = params.FilterExp.Append(goqu.C({{$.Entity.DBGoCase}}.{{$Col.DBGoCase}}).Eq(base.StatusActive))
		{{- end}}
		{{- end}}
	}
	query := NewSelect(
			{{- range $Col := $.Entity.Columns}}
				{{$.Entity.DBGoCase}}.{{$Col.DBGoCase}},
			{{- end}}
		).From(T.{{$.Entity.DBGoCase}}).
		Where(params.FilterExp).
		Limit(params.Limit).
		Offset(params.Offset).
		Order(params.Sort...)

	rows, err := QueryContext(ctx, exe, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, nil
		}
		DebugErr(ctx, err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		{{- range $Col := $.Entity.Columns}}
		{{- if $Col.DBNullable}}
			var {{$Col.GoVarName}} Zero{{$Col.GoDataType | upperfirst}}
		{{- end}}
		{{- end}}
		rec := {{$.ServiceAbbr | lowercase}}model.{{$.Entity.GoStruct}}{}
		err := rows.Scan(
			{{- range $Col := $.Entity.Columns}}
			{{- if $Col.DBNullable}}
				&{{$Col.GoVarName}},
			{{- else}}
				&rec.{{$Col.GoField}},
			{{- end}}
			{{- end}}
		)
		if err != nil {
			DebugErr(ctx, err)
			return nil, err
		}
		{{- range $Col := $.Entity.Columns}}
		{{- if $Col.DBNullable}}
			if {{$Col.GoVarName}} != "" {
				rec.{{$Col.GoField}} = {{$Col.GoDataType}}({{$Col.GoVarName}} )
			}
		{{- end}}
		{{- end}}
		res = append(res, rec)
	}
	return res, nil
}

// New creates a new record.
func (dao *Dao{{$.Entity.GoStruct}}) New(ctx context.Context, tx *sql.Tx,
	rec *{{$.ServiceAbbr | lowercase}}model.{{$.Entity.GoStruct}},
) error {
	now := time.Now()
	record := goqu.Record{
		{{- range $Col := $.Entity.Columns}}
		{{- if not $Col.IsPK}}
			{{- if eq $Col.GoField "UpdatedAt"}}
				{{$.Entity.DBGoCase}}.{{$Col.DBGoCase}}:now,
			{{- else if $Col.DBNullable}}
				{{$.Entity.DBGoCase}}.{{$Col.DBGoCase}}:Zero{{$Col.GoDataType | upperfirst}}(rec.{{$Col.GoField}}),
			{{- else}}
				{{$.Entity.DBGoCase}}.{{$Col.DBGoCase}}:rec.{{$Col.GoField}},
			{{- end}}
		{{- end}}
		{{- end}}
	}
	ins := NewInsert(T.{{$.Entity.DBGoCase}}).Rows(record)
	row, err := DoInsertReturning(ctx, tx, ins,
		{{- range $Col := $.Entity.Columns}}
		{{- if $Col.IsPK}}
			{{$.Entity.DBGoCase}}.{{$Col.DBGoCase}},
		{{- end}}
		{{- end}}
	)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	err = row.Scan(&rec.ID)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	rec.UpdatedAt = now
	return nil
}

// Edit edits the record.
func (dao *Dao{{$.Entity.GoStruct}}) Edit(ctx context.Context, tx *sql.Tx,
	rec *{{$.ServiceAbbr | lowercase}}model.{{$.Entity.GoStruct}},
) error {
	now := time.Now()
	record := goqu.Record{
		{{- range $Col := $.Entity.Columns}}
		{{- if not $Col.IsPK}}
			{{- if eq $Col.GoField "UpdatedAt"}}
				{{$.Entity.DBGoCase}}.{{$Col.DBGoCase}}:now,
			{{- else if $Col.DBNullable}}
				{{$.Entity.DBGoCase}}.{{$Col.DBGoCase}}:Zero{{$Col.GoDataType | upperfirst}}(rec.{{$Col.GoField}}),
			{{- else}}
				{{$.Entity.DBGoCase}}.{{$Col.DBGoCase}}:rec.{{$Col.GoField}},
			{{- end}}
		{{- end}}
		{{- end}}
	}

	query := NewUpdate(T.{{$.Entity.DBGoCase}}).
		Set(record).
		Where(
			{{- range $Col := $.Entity.Columns}}
			{{- if or ($Col.IsPK) (eq $Col.GoField "UpdatedAt")}}
				goqu.C({{$.Entity.DBGoCase}}.{{$Col.DBGoCase}}).Eq(rec.{{$Col.GoField}}),
			{{else if eq $Col.GoField "Status"}}
			{{- end}}
			{{- end}}
		)
	res, err := DoUpdate(ctx, tx, query)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	rec.UpdatedAt = now
	return CheckOneRowUpdated(ctx, T.{{$.Entity.DBGoCase}}, res)
}

// SetStatus updates the logical status of the record.
func (dao *Dao{{$.Entity.GoStruct}}) SetStatus(ctx context.Context, tx *sql.Tx,
	rec *{{$.ServiceAbbr | lowercase}}model.{{$.Entity.GoStruct}},
) error {
	now := time.Now()
	query := NewUpdate(T.{{$.Entity.DBGoCase}}).
		Set(goqu.Record{
			{{- range $Col := $.Entity.Columns}}
				{{- if eq $Col.GoField "UpdatedAt"}}
					{{$.Entity.DBGoCase}}.{{$Col.DBGoCase}}:now,
				{{- else if eq $Col.GoField "Status"}}
					{{$.Entity.DBGoCase}}.{{$Col.DBGoCase}}:rec.{{$Col.GoField}},
				{{- else}}
				{{- end}}
			{{- end}}
		}).
		Where(
			{{- range $Col := $.Entity.Columns}}
			{{- if or ($Col.IsPK) (eq $Col.GoField "UpdatedAt")}}
				goqu.C({{$.Entity.DBGoCase}}.{{$Col.DBGoCase}}).Eq(rec.{{$Col.GoField}}),
			{{- end}}
			{{- end}}
		)
	res, err := DoUpdate(ctx, tx, query)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	rec.UpdatedAt = now
	return CheckOneRowUpdated(ctx, T.{{$.Entity.DBGoCase}}, res)
}

// Delete performs a physical delete of the record.
func (dao *Dao{{$.Entity.GoStruct}}) Delete(ctx context.Context, tx *sql.Tx,
	rec *{{$.ServiceAbbr | lowercase}}model.{{$.Entity.GoStruct}},
) error {
	query := NewDelete(T.{{$.Entity.DBGoCase}}).
		Where(
			{{- range $Col := $.Entity.Columns}}
			{{- if or ($Col.IsPK) (eq $Col.GoField "UpdatedAt")}}
				goqu.C({{$.Entity.DBGoCase}}.{{$Col.DBGoCase}}).Eq(rec.{{$Col.GoField}}),
			{{- end}}
			{{- end}}
		)
	res, err := DoDelete(ctx, tx, query)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	return CheckOneRowUpdated(ctx, T.{{$.Entity.DBGoCase}}, res)
}

package {{$.ServiceAbbr | lowercase}}biz

import (
	"context"
	"database/sql"

	"github.com/jtorz/phoenix-backend/app/services/{{$.ServiceAbbr | lowercase}}/{{$.ServiceAbbr | lowercase}}dao"
	"github.com/jtorz/phoenix-backend/app/services/{{$.ServiceAbbr | lowercase}}/{{$.ServiceAbbr | lowercase}}model"
	"github.com/jtorz/phoenix-backend/app/shared/base"
)

// Biz{{$.Entity.GoStruct}} business component.
type Biz{{$.Entity.GoStruct}} struct {
	Exe base.Executor
	dao *{{$.ServiceAbbr | lowercase}}dao.Dao{{$.Entity.GoStruct}}
}

// NewBiz{{$.Entity.GoStruct}} creates the business component.
func NewBiz{{$.Entity.GoStruct}}() Biz{{$.Entity.GoStruct}} {
	return Biz{{$.Entity.GoStruct}}{
		dao: &{{$.ServiceAbbr | lowercase}}dao.Dao{{$.Entity.GoStruct}}{},
	}
}

// GetByID retrives the record information using its ID.
func (biz *Biz{{$.Entity.GoStruct}}) GetByID(ctx context.Context, exe base.Executor,
	{{range $Col := $.Entity.Columns}}
	{{- if $Col.IsPK}} {{$Col.GoVarName}} {{$Col.GoDataType}}, {{end}}
	{{- end}}
) (*{{$.ServiceAbbr | lowercase}}model.{{$.Entity.GoStruct}}, error) {
	rec, err := biz.dao.GetByID(ctx, exe,
	{{- range $Col := $.Entity.Columns}}
	{{- if $Col.IsPK}} {{$Col.GoVarName}}, {{end}}
	{{- end}})
	if err != nil {
		return nil, err
	}
	biz.setRecordActions(ctx, rec)
	return rec, nil
}

// List returns the list of records that can be filtered by the user.
func (biz *Biz{{$.Entity.GoStruct}}) List(ctx context.Context, exe base.Executor,
	qry base.ClientQuery,
) ({{$.ServiceAbbr | lowercase}}model.{{$.Entity.GoSlice}}, error) {
	recs, err := biz.dao.List(ctx, exe, qry)
	if err != nil {
		return nil, err
	}
	biz.setRecordActions{{$.Entity.GoSlice}}(ctx, recs)
	return recs, nil
}

// New creates a new record.
func (biz *Biz{{$.Entity.GoStruct}}) New(ctx context.Context, tx *sql.Tx,
	rec *{{$.ServiceAbbr | lowercase}}model.{{$.Entity.GoStruct}},
) error {
	rec.Status = base.StatusCaptured
	if err := biz.dao.New(ctx, tx, rec); err != nil {
		return err
	}
	biz.setRecordActions(ctx, rec)
	return nil
}

// Edit edits the record.
func (biz *Biz{{$.Entity.GoStruct}}) Edit(ctx context.Context, tx *sql.Tx,
	rec *{{$.ServiceAbbr | lowercase}}model.{{$.Entity.GoStruct}},
) error {
	return biz.dao.Edit(ctx, tx, rec)
}

// SetStatus updates the logical status of the record.
func (biz *Biz{{$.Entity.GoStruct}}) SetStatus(ctx context.Context, tx *sql.Tx,
	rec *{{$.ServiceAbbr | lowercase}}model.{{$.Entity.GoStruct}},
) error {
	if err := biz.dao.SetStatus(ctx, tx, rec); err != nil {
		return err
	}
	biz.setRecordActions(ctx, rec)
	return nil
}

// Delete performs a physical delete of the record.
func (biz *Biz{{$.Entity.GoStruct}}) Delete(ctx context.Context, tx *sql.Tx,
	rec *{{$.ServiceAbbr | lowercase}}model.{{$.Entity.GoStruct}},
) error {
	if err := biz.dao.Delete(ctx, tx, rec); err != nil {
		return err
	}
	return nil
}

// setRecordActions{{$.Entity.GoSlice}} sets the record actiosn to every element in the {{$.Entity.GoSlice}} slice.
func (biz *Biz{{$.Entity.GoStruct}}) setRecordActions{{$.Entity.GoSlice}}(ctx context.Context,
	recs {{$.ServiceAbbr | lowercase}}model.{{$.Entity.GoSlice}},
) {
	for i := range recs {
		biz.setRecordActions(ctx, &recs[i])
	}
}

// setRecordActions sets the record action sto {{$.Entity.GoStruct}} record.
func (biz *Biz{{$.Entity.GoStruct}}) setRecordActions(ctx context.Context,
	rec *{{$.ServiceAbbr | lowercase}}model.{{$.Entity.GoStruct}},
) {
	rec.RecordActions = base.NewRecordActionsCommon(rec.Status)
}

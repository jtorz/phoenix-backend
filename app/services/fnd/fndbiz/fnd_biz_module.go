package fndbiz

import (
	"context"
	"database/sql"

	"github.com/jtorz/phoenix-backend/app/services/fnd/fnddao"
	"github.com/jtorz/phoenix-backend/app/services/fnd/fndmodel"
	"github.com/jtorz/phoenix-backend/app/shared/base"
)

// BizModule business component.
type BizModule struct {
	Exe base.Executor
	dao *fnddao.DaoModule
}

// NewBizModule creates the business component.
func NewBizModule() BizModule {
	return BizModule{
		dao: &fnddao.DaoModule{},
	}
}

// GetByID retrives the record information using its ID.
func (biz *BizModule) GetByID(ctx context.Context, exe base.Executor,
	id string,
) (*fndmodel.Module, error) {
	rec, err := biz.dao.GetByID(ctx, exe, id)
	if err != nil {
		return nil, err
	}
	biz.setRecordActions(ctx, rec)
	return rec, nil
}

// List returns the list of records that can be filtered by the user.
func (biz *BizModule) List(ctx context.Context, exe base.Executor,
	qry base.ClientQuery,
) (fndmodel.Modules, error) {
	recs, err := biz.dao.List(ctx, exe, qry)
	if err != nil {
		return nil, err
	}
	biz.setRecordActionsModules(ctx, recs)
	return recs, nil
}

// New creates a new record.
func (biz *BizModule) New(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.Module,
) error {
	rec.Status = base.StatusCaptured
	if err := biz.dao.New(ctx, tx, rec); err != nil {
		return err
	}
	biz.setRecordActions(ctx, rec)
	return nil
}

// Edit edits the record.
func (biz *BizModule) Edit(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.Module,
) error {
	return biz.dao.Edit(ctx, tx, rec)
}

// SetStatus updates the logical status of the record.
func (biz *BizModule) SetStatus(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.Module,
) error {
	if err := biz.dao.SetStatus(ctx, tx, rec); err != nil {
		return err
	}
	biz.setRecordActions(ctx, rec)
	return nil
}

// Delete performs a physical delete of the record.
func (biz *BizModule) Delete(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.Module,
) error {
	if err := biz.dao.Delete(ctx, tx, rec); err != nil {
		return err
	}
	return nil
}

// setRecordActionsModules sets the record actiosn to every element in the Modules slice.
func (biz *BizModule) setRecordActionsModules(ctx context.Context,
	recs fndmodel.Modules,
) {
	for i := range recs {
		biz.setRecordActions(ctx, &recs[i])
	}
}

// setRecordActions sets the record action sto Module record.
func (biz *BizModule) setRecordActions(ctx context.Context,
	rec *fndmodel.Module,
) {
	rec.RecordActions = base.NewRecordActionsCommon(rec.Status)
}

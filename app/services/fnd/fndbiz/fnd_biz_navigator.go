package fndbiz

import (
	"context"
	"database/sql"

	"github.com/jtorz/phoenix-backend/app/services/fnd/fnddao"
	"github.com/jtorz/phoenix-backend/app/services/fnd/fndmodel"
	"github.com/jtorz/phoenix-backend/app/shared/base"
)

// BizNavigator business component.
type BizNavigator struct {
	Exe base.Executor
	dao *fnddao.DaoNavigator
}

// NewBizNavigator creates the business component.
func NewBizNavigator() BizNavigator {
	return BizNavigator{
		dao: &fnddao.DaoNavigator{},
	}
}

// GetByID retrives the record information using its ID.
func (biz *BizNavigator) GetByID(ctx context.Context, exe base.Executor,
	id string,
) (*fndmodel.Navigator, error) {
	rec, err := biz.dao.GetByID(ctx, exe, id)
	if err != nil {
		return nil, err
	}
	rec.RecordActions.SimpleActions(rec.Status)
	return rec, nil
}

// List returns the list of records that can be filtered by the user.
func (biz *BizNavigator) List(ctx context.Context, exe base.Executor,
	qry base.ClientQuery,
) (fndmodel.Navigators, error) {
	rows, err := biz.dao.List(ctx, exe, qry)
	if err != nil {
		return nil, err
	}
	for i := range rows {
		rows[i].RecordActions.SimpleActions(rows[i].Status)
	}
	return rows, nil
}

// New creates a new record.
func (biz *BizNavigator) New(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.Navigator,
) error {
	rec.Status = base.StatusCaptured
	if err := biz.dao.New(ctx, tx, rec); err != nil {
		return err
	}
	rec.RecordActions.SimpleActions(rec.Status)
	return nil
}

// Edit edits the record.
func (biz *BizNavigator) Edit(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.Navigator,
) error {
	return biz.dao.Edit(ctx, tx, rec)
}

// SetStatus updates the logical status of the record.
func (biz *BizNavigator) SetStatus(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.Navigator,
) error {
	if err := biz.dao.SetStatus(ctx, tx, rec); err != nil {
		return err
	}
	rec.RecordActions.SimpleActions(rec.Status)
	return nil
}

// Delete performs a physical delete of the record.
func (biz *BizNavigator) Delete(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.Navigator,
) error {
	if err := biz.dao.Delete(ctx, tx, rec); err != nil {
		return err
	}
	return nil
}

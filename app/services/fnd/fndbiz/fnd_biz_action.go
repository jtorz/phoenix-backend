package fndbiz

import (
	"context"
	"database/sql"

	"github.com/jtorz/phoenix-backend/app/services/fnd/fnddao"
	"github.com/jtorz/phoenix-backend/app/services/fnd/fndmodel"
	"github.com/jtorz/phoenix-backend/app/shared/base"
)

// BizAction business component.
type BizAction struct {
	Exe base.Executor
	dao *fnddao.DaoAction
}

// NewBizAction creates the business component.
func NewBizAction() BizAction {
	return BizAction{
		dao: &fnddao.DaoAction{},
	}
}

// GetByID retrives the record information using its ID.
func (biz *BizAction) GetByID(ctx context.Context, exe base.Executor,
	moduleID string, actionID string,
) (*fndmodel.Action, error) {
	rec, err := biz.dao.GetByID(ctx, exe, moduleID, actionID)
	if err != nil {
		return nil, err
	}
	biz.setRecordActions(ctx, rec)
	return rec, nil
}

// List returns the list of records that can be filtered by the user.
func (biz *BizAction) List(ctx context.Context, exe base.Executor,
	qry base.ClientQuery, moduleID string,
) (fndmodel.Actions, error) {
	recs, err := biz.dao.List(ctx, exe, qry, moduleID)
	if err != nil {
		return nil, err
	}
	biz.setRecordActionsActions(ctx, recs)
	return recs, nil
}

// New creates a new record.
func (biz *BizAction) New(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.Action,
) error {
	rec.Status = base.StatusCaptured
	if err := biz.dao.New(ctx, tx, rec); err != nil {
		return err
	}
	biz.setRecordActions(ctx, rec)
	return nil
}

// Edit edits the record.
func (biz *BizAction) Edit(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.Action,
) error {
	return biz.dao.Edit(ctx, tx, rec)
}

// SetStatus updates the logical status of the record.
func (biz *BizAction) SetStatus(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.Action,
) error {
	if err := biz.dao.SetStatus(ctx, tx, rec); err != nil {
		return err
	}
	biz.setRecordActions(ctx, rec)
	return nil
}

// Delete performs a physical delete of the record.
func (biz *BizAction) Delete(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.Action,
) error {
	if err := biz.dao.Delete(ctx, tx, rec); err != nil {
		return err
	}
	return nil
}

// setRecordActionsActions sets the record actions to every element in the Actions slice.
func (biz *BizAction) setRecordActionsActions(ctx context.Context,
	recs fndmodel.Actions,
) {
	for i := range recs {
		biz.setRecordActions(ctx, &recs[i])
	}
}

// setRecordActions sets the record actions to Action record.
func (biz *BizAction) setRecordActions(ctx context.Context,
	rec *fndmodel.Action,
) {
	rec.RecordActions = base.NewRecordActionsCommon(rec.Status)
}

package fndbiz

import (
	"context"
	"database/sql"

	"github.com/jtorz/phoenix-backend/app/services/fnd/fnddao"
	"github.com/jtorz/phoenix-backend/app/services/fnd/fndmodel"
	"github.com/jtorz/phoenix-backend/app/shared/base"
)

// BizNavElement business component.
type BizNavElement struct {
	Exe base.Executor
	dao *fnddao.DaoNavElement
}

// NewBizNavElement creates the business component.
func NewBizNavElement() BizNavElement {
	return BizNavElement{
		dao: &fnddao.DaoNavElement{},
	}
}

// GetByID retrives the record information using its ID.
func (biz *BizNavElement) GetByID(ctx context.Context, exe base.Executor,
	id string,
) (*fndmodel.NavElement, error) {
	rec, err := biz.dao.GetByID(ctx, exe, id)
	if err != nil {
		return nil, err
	}
	rec.RecordActions.SimpleActions(rec.Status)
	return rec, nil
}

// ListAll returns the all the navigator records with the flag IsAssigned
// if the navigator element is assigned to the role.
func (biz *BizNavElement) ListActive(ctx context.Context, exe base.Executor,
	rolID string,
) (fndmodel.Navigator, error) {
	rows, err := biz.dao.ListAll(ctx, exe, rolID)
	if err != nil {
		return nil, err
	}
	return rows.Tree(), nil
}

// New creates a new record.
func (biz *BizNavElement) New(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.NavElement,
) error {
	rec.Status = base.StatusCaptured
	if err := biz.dao.New(ctx, tx, rec); err != nil {
		return err
	}
	rec.RecordActions.SimpleActions(rec.Status)
	return nil
}

// Edit edits the record.
func (biz *BizNavElement) Edit(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.NavElement,
) error {
	return biz.dao.Edit(ctx, tx, rec)
}

// SetStatus updates the logical status of the record.
func (biz *BizNavElement) SetStatus(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.NavElement,
) error {
	if err := biz.dao.SetStatus(ctx, tx, rec); err != nil {
		return err
	}
	rec.RecordActions.SimpleActions(rec.Status)
	return nil
}

// Delete performs a physical delete of the record.
func (biz *BizNavElement) Delete(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.NavElement,
) error {
	if err := biz.dao.Delete(ctx, tx, rec); err != nil {
		return err
	}
	return nil
}

// AssociateRole associates the nav element to the role.
func (biz *BizNavElement) AssociateRole(ctx context.Context, tx *sql.Tx,
	elementID, roleID string,
) error {
	return biz.dao.AssociateRole(ctx, tx, elementID, roleID)
}

// DissociateRole dissociates the nav element from the role.
func (biz *BizNavElement) DissociateRole(ctx context.Context, tx *sql.Tx,
	elementID, roleID string,
) error {
	return biz.dao.DissociateRole(ctx, tx, elementID, roleID)
}

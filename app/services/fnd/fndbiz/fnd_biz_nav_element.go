package fndbiz

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jtorz/phoenix-backend/app/services/fnd/fnddao"
	"github.com/jtorz/phoenix-backend/app/services/fnd/fndmodel"
	"github.com/jtorz/phoenix-backend/app/shared/base"
	"github.com/jtorz/phoenix-backend/app/shared/baseerrors"
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

// New creates a new record.
func (biz *BizNavElement) UpsertAll(ctx context.Context, tx *sql.Tx,
	nav fndmodel.Navigator,
) error {
	if err := biz.upsertAll(ctx, tx, nav, ""); err != nil {
		return err
	}
	return nil
}

// New creates a new record.
func (biz *BizNavElement) upsertAll(ctx context.Context, tx *sql.Tx,
	nav fndmodel.Navigator, parentID string,
) error {
	if len(nav) == 0 {
		return nil
	}
	for i := range nav {
		nav[i].ParentID = parentID
		nav[i].Order = i
		if err := biz.upsert(ctx, tx, &nav[i]); err != nil {
			return err
		}
		if err := biz.upsertAll(ctx, tx, nav[i].Children, nav[i].ID); err != nil {
			return err
		}
	}
	return nil
}

// New creates a new record.
func (biz *BizNavElement) upsert(ctx context.Context, tx *sql.Tx,
	navElem *fndmodel.NavElement,
) error {
	exists := true
	oldRec, err := biz.GetByID(ctx, tx, navElem.ID)
	if err != nil {
		if baseerrors.IsErrNotFound(err) {
			exists = false
		} else {
			return err
		}
	}
	if navElem.Deleted {
		if exists {
			navElem.UpdatedAt = oldRec.UpdatedAt
			if err := biz.Delete(ctx, tx, navElem); err != nil {
				return err
			}
		}
		navElem.ID = ""
		return nil
	}

	if exists {
		navElem.UpdatedAt = oldRec.UpdatedAt
		err = biz.Edit(ctx, tx, navElem)
	} else {
		err = biz.New(ctx, tx, navElem)
	}
	return err
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
func (biz *BizNavElement) ListAll(ctx context.Context, exe base.Executor,
	rolID string,
) (fndmodel.Navigator, error) {
	rows, err := biz.dao.ListAll(ctx, exe, rolID)
	if err != nil {
		return nil, err
	}
	tree := rows.Tree()
	biz.setRecordActionsNavigator(ctx, tree)
	return tree, nil
}

// New creates a new record.
func (biz *BizNavElement) New(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.NavElement,
) error {
	if rec.ID == "" {
		return fmt.Errorf("%w (nav element id)", baseerrors.ErrInvalidData)
	}
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
	biz.setRecordActions(ctx, rec)
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

func (biz *BizNavElement) setRecordActionsNavigator(ctx context.Context,
	nav fndmodel.Navigator,
) {
	for i := range nav {
		biz.setRecordActions(ctx, &nav[i])
		biz.setRecordActionsNavigator(ctx, nav[i].Children)
	}
}

func (biz *BizNavElement) setRecordActions(ctx context.Context,
	navElem *fndmodel.NavElement,
) {
	switch navElem.Status {
	case base.StatusActive:
		navElem.RecordActions = base.RecordActions{"invalidate"}
	default:
		navElem.RecordActions = base.RecordActions{"validate"}
	}
	biz.setRecordActionsNavigator(ctx, navElem.Children)
}

package fndbiz

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jtorz/phoenix-backend/app/services/fnd/fnddal"
	"github.com/jtorz/phoenix-backend/app/services/fnd/fndmodel"
	"github.com/jtorz/phoenix-backend/app/shared/base"
	"github.com/jtorz/phoenix-backend/app/shared/baseerrors"
)

// BizNavElement business component.
type BizNavElement struct {
	Exe base.Executor
	dal *fnddal.DalNavElement
}

// NewBizNavElement creates the business component.
func NewBizNavElement() BizNavElement {
	return BizNavElement{
		dal: &fnddal.DalNavElement{},
	}
}

// UpsertOrDeleteAll processes the all the records with the Create, Update or Delete actions.
//
// Delete: the record is deleted if the field NavElement.Deleted is true.
// Create: the record is created if doesn't exist.
// Update: the record is updated if already exists.
func (biz *BizNavElement) UpsertOrDeleteAll(ctx context.Context, tx *sql.Tx,
	nav fndmodel.Navigator,
) error {
	if err := biz.upsertOrDeleteAll(ctx, tx, nav, ""); err != nil {
		return err
	}
	return nil
}

// upsertOrDeleteAll processes the all the records with the Create, Update or Delete actions.
//
// Delete: the record is deleted if the field NavElement.Deleted is true.
// Create: the record is created if doesn't exist.
// Update: the record is updated if already exists.
func (biz *BizNavElement) upsertOrDeleteAll(ctx context.Context, tx *sql.Tx,
	nav fndmodel.Navigator, parentID string,
) error {
	if len(nav) == 0 {
		return nil
	}
	for i := range nav {
		nav[i].ParentID = parentID
		nav[i].Order = i
		if err := biz.upsertOrDelete(ctx, tx, &nav[i]); err != nil {
			return err
		}
		if err := biz.upsertOrDeleteAll(ctx, tx, nav[i].Children, nav[i].ID); err != nil {
			return err
		}
	}
	return nil
}

// upsertOrDelete processes the record with the Create, Update or Delete actions.
//
// Delete: the record is deleted if the field NavElement.Deleted is true.
// Create: the record is created if doesn't exist.
// Update: the record is updated if already exists.
func (biz *BizNavElement) upsertOrDelete(ctx context.Context, tx *sql.Tx,
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
	rec, err := biz.dal.GetByID(ctx, exe, id)
	if err != nil {
		return nil, err
	}
	biz.setRecordActions(ctx, rec)
	return rec, nil
}

// ListAll returns the all the navigator records with the flag IsAssigned
// if the navigator element is assigned to the role.
func (biz *BizNavElement) ListAll(ctx context.Context, exe base.Executor,
	rolID string,
) (fndmodel.Navigator, error) {
	rows, err := biz.dal.ListAll(ctx, exe, rolID)
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
	if err := biz.dal.New(ctx, tx, rec); err != nil {
		return err
	}
	biz.setRecordActions(ctx, rec)
	return nil
}

// Edit edits the record.
func (biz *BizNavElement) Edit(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.NavElement,
) error {
	return biz.dal.Edit(ctx, tx, rec)
}

// SetStatus updates the logical status of the record.
func (biz *BizNavElement) SetStatus(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.NavElement,
) error {
	if err := biz.dal.SetStatus(ctx, tx, rec); err != nil {
		return err
	}
	biz.setRecordActions(ctx, rec)
	return nil
}

// Delete performs a physical delete of the record.
func (biz *BizNavElement) Delete(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.NavElement,
) error {
	if err := biz.dal.Delete(ctx, tx, rec); err != nil {
		return err
	}
	return nil
}

// AssociateRole associates the nav element to the role.
func (biz *BizNavElement) AssociateRole(ctx context.Context, tx *sql.Tx,
	elementID, roleID string,
) error {
	return biz.dal.AssociateRole(ctx, tx, elementID, roleID)
}

// DissociateRole dissociates the nav element from the role.
func (biz *BizNavElement) DissociateRole(ctx context.Context, tx *sql.Tx,
	elementID, roleID string,
) error {
	return biz.dal.DissociateRole(ctx, tx, elementID, roleID)
}

// setRecordActionsNavElements sets the records action to every element in the Navigator slice.
func (biz *BizNavElement) setRecordActionsNavigator(ctx context.Context,
	nav fndmodel.Navigator,
) {
	for i := range nav {
		biz.setRecordActions(ctx, &nav[i])
	}
}

// setRecordActions sets the record actions to NavElement record.
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

package fndbiz

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jtorz/phoenix-backend/app/services/fnd/fnddao"
	"github.com/jtorz/phoenix-backend/app/services/fnd/fndmodel"
	"github.com/jtorz/phoenix-backend/app/shared/base"
	"github.com/jtorz/phoenix-backend/app/shared/baseerrors"
	"github.com/jtorz/phoenix-backend/app/shared/baseservice"
)

// BizRole business component.
type BizRole struct {
	Exe base.Executor
	dao *fnddao.DaoRole
}

// NewBizRole creates the business component.
func NewBizRole() BizRole {
	return BizRole{
		dao: &fnddao.DaoRole{},
	}
}

// GetByID retrives the record information using its ID.
func (biz *BizRole) GetByID(ctx context.Context, exe base.Executor,
	id string,
) (*fndmodel.Role, error) {
	rec, err := biz.dao.GetByID(ctx, exe, id)
	if err != nil {
		return nil, err
	}
	biz.setRecordActions(ctx, rec)
	return rec, nil
}

// List returns the list of records that can be filtered by the user.
func (biz *BizRole) List(ctx context.Context, exe base.Executor,
	OnlyActive bool,
) (fndmodel.Roles, error) {
	recs, err := biz.dao.List(ctx, exe, OnlyActive)
	if err != nil {
		return nil, err
	}
	biz.setRecordActionsRoles(ctx, recs)
	return recs, nil
}

// New creates a new record.
func (biz *BizRole) New(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.Role,
) error {
	rec.Status = base.StatusCaptured
	if err := biz.dao.New(ctx, tx, rec); err != nil {
		return err
	}
	biz.setRecordActions(ctx, rec)
	return nil
}

// Edit edits the record.
func (biz *BizRole) Edit(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.Role,
) error {
	return biz.dao.Edit(ctx, tx, rec)
}

// SetStatus updates the logical status of the record.
func (biz *BizRole) SetStatus(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.Role,
) error {
	if rec.ID == baseservice.RoleAdmin {
		return fmt.Errorf("can't change admin role status: %w", baseerrors.ErrPrivilege)
	}
	if err := biz.dao.SetStatus(ctx, tx, rec); err != nil {
		return err
	}
	biz.setRecordActions(ctx, rec)
	return nil
}

// Delete performs a physical delete of the record.
func (biz *BizRole) Delete(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.Role,
) error {
	if rec.ID == baseservice.RoleAdmin {
		return fmt.Errorf("can't delete admin role: %w", baseerrors.ErrPrivilege)
	}
	if err := biz.dao.Delete(ctx, tx, rec); err != nil {
		return err
	}
	return nil
}

// setRecordActionsRoles sets the record actiosn to every element in the Roles slice.
func (biz *BizRole) setRecordActionsRoles(ctx context.Context,
	recs fndmodel.Roles,
) {
	for i := range recs {
		biz.setRecordActions(ctx, &recs[i])
	}
}

// setRecordActions sets the record action sto Role record.
func (biz *BizRole) setRecordActions(ctx context.Context,
	rec *fndmodel.Role,
) {
	rec.RecordActions = base.NewRecordActionsCommon(rec.Status)
}

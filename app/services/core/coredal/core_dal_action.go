package coredal

import (
	"context"
	"database/sql"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/jtorz/phoenix-backend/app/services/core/coremodel"
	"github.com/jtorz/phoenix-backend/app/shared/base"

	//lint:ignore ST1001 dot import allowed only in dal packages for
	. "github.com/jtorz/phoenix-backend/app/shared/lex"
)

// DalAction Data Access structure.
type DalAction struct{}

// GetByID retrives the record information using its ID.
func (dal *DalAction) GetByID(ctx context.Context, exe base.Executor,
	moduleID string, actionID string,
) (*coremodel.Action, error) {
	rec := coremodel.Action{}
	query := NewSelect(
		CoreAction.ActName,
		CoreAction.ActDescription,
		CoreAction.ActOrder,
		CoreAction.ActCreatedAt,
		CoreAction.ActUpdatedAt,
		CoreAction.ActStatus,
	).
		From(goqu.T(T.CoreAction)).
		Where(
			goqu.C(CoreAction.ActModuleID).Eq(moduleID),
			goqu.C(CoreAction.ActActionID).Eq(actionID),
		)

	row, err := QueryRowContext(ctx, exe, query)
	if err != nil {
		DebugErr(ctx, err)
		return nil, err
	}
	err = row.Scan(
		&rec.Name,
		&rec.Description,
		&rec.Order,
		&rec.CreatedAt,
		&rec.UpdatedAt,
		&rec.Status,
	)
	if err != nil {
		DebugErr(ctx, err)
		return nil, WrapNotFound(ctx, T.CoreAction, err)
	}

	rec.ModuleID = moduleID
	rec.ActionID = actionID

	return &rec, nil
}

// List returns the list of records that can be filtered by the user.
func (dal *DalAction) List(ctx context.Context, exe base.Executor,
	qry base.ClientQuery, moduleID string,
) (coremodel.Actions, error) {
	res := make(coremodel.Actions, 0)
	params, err := ParseClientFilter(qry, coremodel.Action{})
	if err != nil {
		return nil, err
	}
	if qry.OnlyActive {
		params.FilterExp = params.FilterExp.Append(goqu.C(CoreAction.ActStatus).Eq(base.StatusActive))
	}
	query := NewSelect(
		CoreAction.ActActionID,
		CoreAction.ActName,
		CoreAction.ActDescription,
		CoreAction.ActOrder,
		CoreAction.ActCreatedAt,
		CoreAction.ActUpdatedAt,
		CoreAction.ActStatus,
	).From(T.CoreAction).
		Where(params.FilterExp.Append(
			goqu.C(CoreAction.ActModuleID).Eq(moduleID),
		)).
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
		rec := coremodel.Action{ModuleID: moduleID}
		err := rows.Scan(
			&rec.ActionID,
			&rec.Name,
			&rec.Description,
			&rec.Order,
			&rec.CreatedAt,
			&rec.UpdatedAt,
			&rec.Status,
		)
		if err != nil {
			DebugErr(ctx, err)
			return nil, err
		}
		res = append(res, rec)
	}
	return res, nil
}

// New creates a new record.
func (dal *DalAction) New(ctx context.Context, tx *sql.Tx,
	rec *coremodel.Action,
) error {
	now := time.Now()
	record := goqu.Record{
		CoreAction.ActModuleID:    rec.ModuleID,
		CoreAction.ActActionID:    rec.ActionID,
		CoreAction.ActName:        rec.Name,
		CoreAction.ActDescription: rec.Description,
		CoreAction.ActOrder:       rec.Order,
		CoreAction.ActCreatedAt:   rec.CreatedAt,
		CoreAction.ActUpdatedAt:   now,
		CoreAction.ActStatus:      rec.Status,
	}
	ins := NewInsert(T.CoreAction).Rows(record)
	_, err := DoInsert(ctx, tx, ins)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	rec.UpdatedAt = now
	return nil
}

// Edit edits the record.
func (dal *DalAction) Edit(ctx context.Context, tx *sql.Tx,
	rec *coremodel.Action,
) error {
	now := time.Now()
	record := goqu.Record{
		CoreAction.ActName:        rec.Name,
		CoreAction.ActDescription: rec.Description,
		CoreAction.ActOrder:       rec.Order,
		CoreAction.ActUpdatedAt:   now,
		CoreAction.ActStatus:      rec.Status,
	}

	query := NewUpdate(T.CoreAction).
		Set(record).
		Where(
			goqu.C(CoreAction.ActModuleID).Eq(rec.ModuleID),

			goqu.C(CoreAction.ActActionID).Eq(rec.ActionID),

			goqu.C(CoreAction.ActUpdatedAt).Eq(rec.UpdatedAt),
		)
	res, err := DoUpdate(ctx, tx, query)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	rec.UpdatedAt = now
	return CheckOneRowUpdated(ctx, T.CoreAction, res)
}

// SetStatus updates the logical status of the record.
func (dal *DalAction) SetStatus(ctx context.Context, tx *sql.Tx,
	rec *coremodel.Action,
) error {
	now := time.Now()
	query := NewUpdate(T.CoreAction).
		Set(goqu.Record{
			CoreAction.ActUpdatedAt: now,
			CoreAction.ActStatus:    rec.Status,
		}).
		Where(
			goqu.C(CoreAction.ActModuleID).Eq(rec.ModuleID),
			goqu.C(CoreAction.ActActionID).Eq(rec.ActionID),
			goqu.C(CoreAction.ActUpdatedAt).Eq(rec.UpdatedAt),
		)
	res, err := DoUpdate(ctx, tx, query)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	rec.UpdatedAt = now
	return CheckOneRowUpdated(ctx, T.CoreAction, res)
}

// Delete performs a physical delete of the record.
func (dal *DalAction) Delete(ctx context.Context, tx *sql.Tx,
	rec *coremodel.Action,
) error {
	query := NewDelete(T.CoreAction).
		Where(
			goqu.C(CoreAction.ActModuleID).Eq(rec.ModuleID),
			goqu.C(CoreAction.ActActionID).Eq(rec.ActionID),
			goqu.C(CoreAction.ActUpdatedAt).Eq(rec.UpdatedAt),
		)
	res, err := DoDelete(ctx, tx, query)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	return CheckOneRowUpdated(ctx, T.CoreAction, res)
}

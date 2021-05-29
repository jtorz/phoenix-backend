package fnddao

import (
	"context"
	"database/sql"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/jtorz/phoenix-backend/app/services/fnd/fndmodel"
	"github.com/jtorz/phoenix-backend/app/shared/base"

	//lint:ignore ST1001 dot import allowed only in dao packages for
	. "github.com/jtorz/phoenix-backend/app/shared/lex"
)

// DaoAction Data Access structure.
type DaoAction struct{}

// GetByID retrives the record information using its ID.
func (dao *DaoAction) GetByID(ctx context.Context, exe base.Executor,
	moduleID string, actionID string,
) (*fndmodel.Action, error) {
	rec := fndmodel.Action{}
	query := NewSelect(
		FndAction.ActName,
		FndAction.ActDescription,
		FndAction.ActOrder,
		FndAction.ActRoute,
		FndAction.ActMethod,
		FndAction.ActCreatedAt,
		FndAction.ActUpdatedAt,
		FndAction.ActStatus,
	).
		From(goqu.T(T.FndAction)).
		Where(
			goqu.C(FndAction.ActModuleID).Eq(moduleID),
			goqu.C(FndAction.ActActionID).Eq(actionID),
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
		&rec.Route,
		&rec.Method,
		&rec.CreatedAt,
		&rec.UpdatedAt,
		&rec.Status,
	)
	if err != nil {
		DebugErr(ctx, err)
		return nil, err
	}

	rec.ModuleID = moduleID
	rec.ActionID = actionID

	return &rec, nil
}

// List returns the list of records that can be filtered by the user.
func (dao *DaoAction) List(ctx context.Context, exe base.Executor,
	qry base.ClientQuery, moduleID string,
) (fndmodel.Actions, error) {
	res := make(fndmodel.Actions, 0)
	params, err := ParseClientFilter(qry, fndmodel.Action{})
	if err != nil {
		return nil, err
	}
	if qry.OnlyActive {
		params.FilterExp = params.FilterExp.Append(goqu.C(FndAction.ActStatus).Eq(base.StatusActive))
	}
	query := NewSelect(
		FndAction.ActActionID,
		FndAction.ActName,
		FndAction.ActDescription,
		FndAction.ActOrder,
		FndAction.ActRoute,
		FndAction.ActMethod,
		FndAction.ActCreatedAt,
		FndAction.ActUpdatedAt,
		FndAction.ActStatus,
	).From(T.FndAction).
		Where(params.FilterExp.Append(
			goqu.C(FndAction.ActModuleID).Eq(moduleID),
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
		rec := fndmodel.Action{ModuleID: moduleID}
		err := rows.Scan(
			&rec.ActionID,
			&rec.Name,
			&rec.Description,
			&rec.Order,
			&rec.Route,
			&rec.Method,
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
func (dao *DaoAction) New(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.Action,
) error {
	now := time.Now()
	record := goqu.Record{
		FndAction.ActModuleID:    rec.ModuleID,
		FndAction.ActActionID:    rec.ActionID,
		FndAction.ActName:        rec.Name,
		FndAction.ActDescription: rec.Description,
		FndAction.ActOrder:       rec.Order,
		FndAction.ActRoute:       rec.Route,
		FndAction.ActMethod:      rec.Method,
		FndAction.ActCreatedAt:   rec.CreatedAt,
		FndAction.ActUpdatedAt:   now,
		FndAction.ActStatus:      rec.Status,
	}
	ins := NewInsert(T.FndAction).Rows(record)
	_, err := DoInsert(ctx, tx, ins)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	rec.UpdatedAt = now
	return nil
}

// Edit edits the record.
func (dao *DaoAction) Edit(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.Action,
) error {
	now := time.Now()
	record := goqu.Record{
		FndAction.ActName:        rec.Name,
		FndAction.ActDescription: rec.Description,
		FndAction.ActOrder:       rec.Order,
		FndAction.ActRoute:       rec.Route,
		FndAction.ActMethod:      rec.Method,
		FndAction.ActCreatedAt:   rec.CreatedAt,
		FndAction.ActUpdatedAt:   now,
		FndAction.ActStatus:      rec.Status,
	}

	query := NewUpdate(T.FndAction).
		Set(record).
		Where(
			goqu.C(FndAction.ActModuleID).Eq(rec.ModuleID),

			goqu.C(FndAction.ActActionID).Eq(rec.ActionID),

			goqu.C(FndAction.ActUpdatedAt).Eq(rec.UpdatedAt),
		)
	res, err := DoUpdate(ctx, tx, query)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	rec.UpdatedAt = now
	return CheckOneRowUpdated(ctx, T.FndAction, res)
}

// SetStatus updates the logical status of the record.
func (dao *DaoAction) SetStatus(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.Action,
) error {
	now := time.Now()
	query := NewUpdate(T.FndAction).
		Set(goqu.Record{
			FndAction.ActUpdatedAt: now,
			FndAction.ActStatus:    rec.Status,
		}).
		Where(
			goqu.C(FndAction.ActModuleID).Eq(rec.ModuleID),
			goqu.C(FndAction.ActActionID).Eq(rec.ActionID),
			goqu.C(FndAction.ActUpdatedAt).Eq(rec.UpdatedAt),
		)
	res, err := DoUpdate(ctx, tx, query)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	rec.UpdatedAt = now
	return CheckOneRowUpdated(ctx, T.FndAction, res)
}

// Delete performs a physical delete of the record.
func (dao *DaoAction) Delete(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.Action,
) error {
	query := NewDelete(T.FndAction).
		Where(
			goqu.C(FndAction.ActModuleID).Eq(rec.ModuleID),
			goqu.C(FndAction.ActActionID).Eq(rec.ActionID),
			goqu.C(FndAction.ActUpdatedAt).Eq(rec.UpdatedAt),
		)
	res, err := DoDelete(ctx, tx, query)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	return CheckOneRowUpdated(ctx, T.FndAction, res)
}

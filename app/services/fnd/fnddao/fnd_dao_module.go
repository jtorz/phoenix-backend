package fnddao

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/jtorz/phoenix-backend/app/services/fnd/fndmodel"
	"github.com/jtorz/phoenix-backend/app/shared/base"
	"github.com/jtorz/phoenix-backend/app/shared/baseerrors"

	//lint:ignore ST1001 dot import allowed only in dao packages for
	. "github.com/jtorz/phoenix-backend/app/shared/lex"
)

// DaoModule Data Access structure.
type DaoModule struct{}

// GetByID retrives the record information using its ID.
func (dao *DaoModule) GetByID(ctx context.Context, exe base.Executor,
	id string,
) (*fndmodel.Module, error) {
	rec := fndmodel.Module{}
	query := NewSelect(T.FndModule).
		Select(
			FndModule.ModName,
			FndModule.ModDescription,
			FndModule.ModOrder,
			FndModule.ModParentID,
			FndModule.ModCreatedAt,
			FndModule.ModUpdatedAt,
			FndModule.ModStatus,
		).
		From(goqu.T(T.FndModule)).
		Where(
			goqu.C(FndModule.ModID).Eq(id),
		)

	row, err := QueryRowContext(ctx, exe, query)
	if err != nil {
		DebugErr(ctx, err)
		return nil, err
	}
	var parentID ZeroString
	err = row.Scan(
		&rec.Name,
		&rec.Description,
		&rec.Order,
		&parentID,
		&rec.CreatedAt,
		&rec.UpdatedAt,
		&rec.Status,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("%s %w", T.FndModule, baseerrors.ErrNotFound)
		}
		DebugErr(ctx, err)
		return nil, err
	}
	if parentID != "" {
		rec.Parent = &fndmodel.Module{ID: string(parentID)}
	}
	rec.ID = id
	return &rec, nil
}

// List returns the list of records that can be filtered by the user.
func (dao *DaoModule) List(ctx context.Context, exe base.Executor,
	qry base.ClientQuery,
) (fndmodel.Modules, error) {
	res := make(fndmodel.Modules, 0)
	params, err := ParseClientFilter(qry, fndmodel.Module{})
	if err != nil {
		return nil, err
	}
	if qry.OnlyActive {
		params.FilterExp = params.FilterExp.Append(goqu.C(FndModule.ModStatus).Eq(base.StatusActive))
	}
	query := NewSelect(T.FndModule).
		Select(
			FndModule.ModID,
			FndModule.ModName,
			FndModule.ModDescription,
			FndModule.ModOrder,
			FndModule.ModParentID,
			FndModule.ModCreatedAt,
			FndModule.ModUpdatedAt,
			FndModule.ModStatus,
		).From(T.FndModule).
		Where(params.FilterExp).
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
		var parentID ZeroString
		rec := fndmodel.Module{}
		err := rows.Scan(
			&rec.ID,
			&rec.Name,
			&rec.Description,
			&rec.Order,
			&parentID,
			&rec.CreatedAt,
			&rec.UpdatedAt,
			&rec.Status,
		)
		if err != nil {
			DebugErr(ctx, err)
			return nil, err
		}
		if parentID != "" {
			rec.Parent = &fndmodel.Module{ID: string(parentID)}
		}
		res = append(res, rec)
	}
	return res, nil
}

// New creates a new record.
func (dao *DaoModule) New(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.Module,
) error {
	now := time.Now()
	record := goqu.Record{
		FndModule.ModID:          rec.ID,
		FndModule.ModName:        rec.Name,
		FndModule.ModDescription: rec.Description,
		FndModule.ModOrder:       rec.Order,
		FndModule.ModCreatedAt:   rec.CreatedAt,
		FndModule.ModUpdatedAt:   now,
		FndModule.ModStatus:      rec.Status,
	}
	if rec.Parent != nil {
		record[FndModule.ModParentID] = ZeroString(rec.Parent.ID)
	}
	ins := NewInsert(T.FndModule).Rows(record)
	_, err := DoInsert(ctx, tx, ins)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	rec.UpdatedAt = now
	return nil
}

// Edit edits the record.
func (dao *DaoModule) Edit(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.Module,
) error {
	now := time.Now()
	record := goqu.Record{
		FndModule.ModName:        rec.Name,
		FndModule.ModDescription: rec.Description,
		FndModule.ModOrder:       rec.Order,
		FndModule.ModCreatedAt:   rec.CreatedAt,
		FndModule.ModUpdatedAt:   now,
		FndModule.ModStatus:      rec.Status,
	}

	if rec.Parent != nil {
		record[FndModule.ModParentID] = ZeroString(rec.Parent.ID)
	}
	query := NewUpdate(T.FndModule).
		Set(record).
		Where(
			goqu.C(FndModule.ModID).Eq(rec.ID),

			goqu.C(FndModule.ModUpdatedAt).Eq(rec.UpdatedAt),
		)
	res, err := DoUpdate(ctx, tx, query)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	rec.UpdatedAt = now
	return CheckOneRowUpdated(ctx, T.FndModule, res)
}

// SetStatus updates the logical status of the record.
func (dao *DaoModule) SetStatus(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.Module,
) error {
	now := time.Now()
	query := NewUpdate(T.FndModule).
		Set(goqu.Record{
			FndModule.ModUpdatedAt: now,
			FndModule.ModStatus:    rec.Status,
		}).
		Where(
			goqu.C(FndModule.ModID).Eq(rec.ID),
			goqu.C(FndModule.ModUpdatedAt).Eq(rec.UpdatedAt),
		)
	res, err := DoUpdate(ctx, tx, query)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	rec.UpdatedAt = now
	return CheckOneRowUpdated(ctx, T.FndModule, res)
}

// Delete performs a physical delete of the record.
func (dao *DaoModule) Delete(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.Module,
) error {
	query := NewDelete(T.FndModule).
		Where(
			goqu.C(FndModule.ModID).Eq(rec.ID),
			goqu.C(FndModule.ModUpdatedAt).Eq(rec.UpdatedAt),
		)
	res, err := DoDelete(ctx, tx, query)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	return CheckOneRowUpdated(ctx, T.FndModule, res)
}

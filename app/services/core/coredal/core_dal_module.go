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

// DalModule Data Access structure.
type DalModule struct{}

// GetByID retrives the record information using its ID.
func (dal *DalModule) GetByID(ctx context.Context, exe base.Executor,
	id string,
) (*coremodel.Module, error) {
	rec := coremodel.Module{}
	query := NewSelect(
		CoreModule.ModName,
		CoreModule.ModDescription,
		CoreModule.ModOrder,
		CoreModule.ModParentID,
		CoreModule.ModCreatedAt,
		CoreModule.ModUpdatedAt,
		CoreModule.ModStatus,
	).
		From(goqu.T(T.CoreModule)).
		Where(
			goqu.C(CoreModule.ModID).Eq(id),
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
		DebugErr(ctx, err)
		return nil, WrapNotFound(ctx, T.CoreModule, err)
	}
	if parentID != "" {
		rec.Parent = &coremodel.Module{ID: string(parentID)}
	}
	rec.ID = id
	return &rec, nil
}

// List returns the list of records that can be filtered by the user.
func (dal *DalModule) List(ctx context.Context, exe base.Executor,
	qry base.ClientQuery,
) (coremodel.Modules, error) {
	res := make(coremodel.Modules, 0)
	params, err := ParseClientFilter(qry, coremodel.Module{})
	if err != nil {
		return nil, err
	}
	if qry.OnlyActive {
		params.FilterExp = params.FilterExp.Append(goqu.C(CoreModule.ModStatus).Eq(base.StatusActive))
	}
	query := NewSelect(
		CoreModule.ModID,
		CoreModule.ModName,
		CoreModule.ModDescription,
		CoreModule.ModOrder,
		CoreModule.ModParentID,
		CoreModule.ModCreatedAt,
		CoreModule.ModUpdatedAt,
		CoreModule.ModStatus,
	).From(T.CoreModule).
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
		rec := coremodel.Module{}
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
			rec.Parent = &coremodel.Module{ID: string(parentID)}
		}
		res = append(res, rec)
	}
	return res, nil
}

// New creates a new record.
func (dal *DalModule) New(ctx context.Context, tx *sql.Tx,
	rec *coremodel.Module,
) error {
	now := time.Now()
	record := goqu.Record{
		CoreModule.ModID:          rec.ID,
		CoreModule.ModName:        rec.Name,
		CoreModule.ModDescription: rec.Description,
		CoreModule.ModOrder:       rec.Order,
		CoreModule.ModCreatedAt:   rec.CreatedAt,
		CoreModule.ModUpdatedAt:   now,
		CoreModule.ModStatus:      rec.Status,
	}
	if rec.Parent != nil {
		record[CoreModule.ModParentID] = ZeroString(rec.Parent.ID)
	}
	ins := NewInsert(T.CoreModule).Rows(record)
	_, err := DoInsert(ctx, tx, ins)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	rec.UpdatedAt = now
	return nil
}

// Edit edits the record.
func (dal *DalModule) Edit(ctx context.Context, tx *sql.Tx,
	rec *coremodel.Module,
) error {
	now := time.Now()
	record := goqu.Record{
		CoreModule.ModName:        rec.Name,
		CoreModule.ModDescription: rec.Description,
		CoreModule.ModOrder:       rec.Order,
		CoreModule.ModUpdatedAt:   now,
		CoreModule.ModStatus:      rec.Status,
	}

	if rec.Parent != nil {
		record[CoreModule.ModParentID] = ZeroString(rec.Parent.ID)
	}
	query := NewUpdate(T.CoreModule).
		Set(record).
		Where(
			goqu.C(CoreModule.ModID).Eq(rec.ID),

			goqu.C(CoreModule.ModUpdatedAt).Eq(rec.UpdatedAt),
		)
	res, err := DoUpdate(ctx, tx, query)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	rec.UpdatedAt = now
	return CheckOneRowUpdated(ctx, T.CoreModule, res)
}

// SetStatus updates the logical status of the record.
func (dal *DalModule) SetStatus(ctx context.Context, tx *sql.Tx,
	rec *coremodel.Module,
) error {
	now := time.Now()
	query := NewUpdate(T.CoreModule).
		Set(goqu.Record{
			CoreModule.ModUpdatedAt: now,
			CoreModule.ModStatus:    rec.Status,
		}).
		Where(
			goqu.C(CoreModule.ModID).Eq(rec.ID),
			goqu.C(CoreModule.ModUpdatedAt).Eq(rec.UpdatedAt),
		)
	res, err := DoUpdate(ctx, tx, query)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	rec.UpdatedAt = now
	return CheckOneRowUpdated(ctx, T.CoreModule, res)
}

// Delete performs a physical delete of the record.
func (dal *DalModule) Delete(ctx context.Context, tx *sql.Tx,
	rec *coremodel.Module,
) error {
	query := NewDelete(T.CoreModule).
		Where(
			goqu.C(CoreModule.ModID).Eq(rec.ID),
			goqu.C(CoreModule.ModUpdatedAt).Eq(rec.UpdatedAt),
		)
	res, err := DoDelete(ctx, tx, query)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	return CheckOneRowUpdated(ctx, T.CoreModule, res)
}

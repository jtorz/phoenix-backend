package coredal

import (
	"context"
	"database/sql"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/jtorz/phoenix-backend/app/services/core/coremodel"
	"github.com/jtorz/phoenix-backend/app/shared/base"

	//lint:ignore ST1001 dot import allowed only in dal packages for
	. "github.com/jtorz/phoenix-backend/app/shared/lex"
)

// DalRole Data Access structure.
type DalRole struct{}

// GetByID retrives the record information using its ID.
func (dal *DalRole) GetByID(ctx context.Context, exe base.Executor,
	id string,
) (*coremodel.Role, error) {
	rec := coremodel.Role{}
	query := NewSelect(
		CoreRole.RolName,
		CoreRole.RolDescription,
		CoreRole.RolIcon,
		CoreRole.RolCreatedAt,
		CoreRole.RolUpdatedAt,
		CoreRole.RolStatus,
	).
		From(goqu.T(T.CoreRole)).
		Where(
			goqu.C(CoreRole.RolID).Eq(id),
		)

	row, err := QueryRowContext(ctx, exe, query)
	if err != nil {
		DebugErr(ctx, err)
		return nil, err
	}
	err = row.Scan(
		&rec.Name,
		&rec.Description,
		&rec.Icon,
		&rec.CreatedAt,
		&rec.UpdatedAt,
		&rec.Status,
	)
	if err != nil {
		DebugErr(ctx, err)
		return nil, WrapNotFound(ctx, T.CoreRole, err)
	}

	rec.ID = id

	return &rec, nil
}

// List returns the list of records that can be filtered by the user.
func (dal *DalRole) List(ctx context.Context, exe base.Executor,
	OnlyActive bool,
) (coremodel.Roles, error) {
	res := make(coremodel.Roles, 0)

	filterExp := exp.NewExpressionList(exp.AndType)
	if OnlyActive {
		filterExp = filterExp.Append(goqu.C(CoreRole.RolStatus).Eq(base.StatusActive))
	}
	query := NewSelect(
		CoreRole.RolID,
		CoreRole.RolName,
		CoreRole.RolDescription,
		CoreRole.RolIcon,
		CoreRole.RolCreatedAt,
		CoreRole.RolUpdatedAt,
		CoreRole.RolStatus,
	).From(T.CoreRole).
		Where(filterExp)

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
		rec := coremodel.Role{}
		err := rows.Scan(
			&rec.ID,
			&rec.Name,
			&rec.Description,
			&rec.Icon,
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
func (dal *DalRole) New(ctx context.Context, tx *sql.Tx,
	rec *coremodel.Role,
) error {
	now := time.Now()
	record := goqu.Record{
		CoreRole.RolName:        rec.Name,
		CoreRole.RolDescription: rec.Description,
		CoreRole.RolIcon:        rec.Icon,
		CoreRole.RolCreatedAt:   rec.CreatedAt,
		CoreRole.RolUpdatedAt:   now,
		CoreRole.RolStatus:      rec.Status,
	}
	ins := NewInsert(T.CoreRole).Rows(record)
	row, err := DoInsertReturning(ctx, tx, ins,
		CoreRole.RolID,
	)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	err = row.Scan(&rec.ID)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	rec.UpdatedAt = now
	return nil
}

// Edit edits the record.
func (dal *DalRole) Edit(ctx context.Context, tx *sql.Tx,
	rec *coremodel.Role,
) error {
	now := time.Now()
	record := goqu.Record{
		CoreRole.RolName:        rec.Name,
		CoreRole.RolDescription: rec.Description,
		CoreRole.RolIcon:        rec.Icon,
		CoreRole.RolUpdatedAt:   now,
		CoreRole.RolStatus:      rec.Status,
	}

	query := NewUpdate(T.CoreRole).
		Set(record).
		Where(
			goqu.C(CoreRole.RolID).Eq(rec.ID),

			goqu.C(CoreRole.RolUpdatedAt).Eq(rec.UpdatedAt),
		)
	res, err := DoUpdate(ctx, tx, query)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	rec.UpdatedAt = now
	return CheckOneRowUpdated(ctx, T.CoreRole, res)
}

// SetStatus updates the logical status of the record.
func (dal *DalRole) SetStatus(ctx context.Context, tx *sql.Tx,
	rec *coremodel.Role,
) error {
	now := time.Now()
	query := NewUpdate(T.CoreRole).
		Set(goqu.Record{
			CoreRole.RolUpdatedAt: now,
			CoreRole.RolStatus:    rec.Status,
		}).
		Where(
			goqu.C(CoreRole.RolID).Eq(rec.ID),
			goqu.C(CoreRole.RolUpdatedAt).Eq(rec.UpdatedAt),
		)
	res, err := DoUpdate(ctx, tx, query)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	rec.UpdatedAt = now
	return CheckOneRowUpdated(ctx, T.CoreRole, res)
}

// Delete performs a physical delete of the record.
func (dal *DalRole) Delete(ctx context.Context, tx *sql.Tx,
	rec *coremodel.Role,
) error {
	query := NewDelete(T.CoreRole).
		Where(
			goqu.C(CoreRole.RolID).Eq(rec.ID),
			goqu.C(CoreRole.RolUpdatedAt).Eq(rec.UpdatedAt),
		)
	res, err := DoDelete(ctx, tx, query)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	return CheckOneRowUpdated(ctx, T.CoreRole, res)
}

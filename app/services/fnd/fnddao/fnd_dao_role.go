package fnddao

import (
	"context"
	"database/sql"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/jtorz/phoenix-backend/app/services/fnd/fndmodel"
	"github.com/jtorz/phoenix-backend/app/shared/base"

	//lint:ignore ST1001 dot import allowed only in dao packages for
	. "github.com/jtorz/phoenix-backend/app/shared/lex"
)

// DaoRole Data Access structure.
type DaoRole struct{}

// GetByID retrives the record information using its ID.
func (dao *DaoRole) GetByID(ctx context.Context, exe base.Executor,
	id string,
) (*fndmodel.Role, error) {
	rec := fndmodel.Role{}
	query := NewSelect(
		FndRole.RolName,
		FndRole.RolDescription,
		FndRole.RolIcon,
		FndRole.RolCreatedAt,
		FndRole.RolUpdatedAt,
		FndRole.RolStatus,
	).
		From(goqu.T(T.FndRole)).
		Where(
			goqu.C(FndRole.RolID).Eq(id),
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
		return nil, err
	}

	rec.ID = id

	return &rec, nil
}

// List returns the list of records that can be filtered by the user.
func (dao *DaoRole) List(ctx context.Context, exe base.Executor,
	OnlyActive bool,
) (fndmodel.Roles, error) {
	res := make(fndmodel.Roles, 0)

	filterExp := exp.NewExpressionList(exp.AndType)
	if OnlyActive {
		filterExp = filterExp.Append(goqu.C(FndRole.RolStatus).Eq(base.StatusActive))
	}
	query := NewSelect(
		FndRole.RolID,
		FndRole.RolName,
		FndRole.RolDescription,
		FndRole.RolIcon,
		FndRole.RolCreatedAt,
		FndRole.RolUpdatedAt,
		FndRole.RolStatus,
	).From(T.FndRole).
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
		rec := fndmodel.Role{}
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
func (dao *DaoRole) New(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.Role,
) error {
	now := time.Now()
	record := goqu.Record{
		FndRole.RolName:        rec.Name,
		FndRole.RolDescription: rec.Description,
		FndRole.RolIcon:        rec.Icon,
		FndRole.RolCreatedAt:   rec.CreatedAt,
		FndRole.RolUpdatedAt:   now,
		FndRole.RolStatus:      rec.Status,
	}
	ins := NewInsert(T.FndRole).Rows(record)
	row, err := DoInsertReturning(ctx, tx, ins,
		FndRole.RolID,
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
func (dao *DaoRole) Edit(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.Role,
) error {
	now := time.Now()
	record := goqu.Record{
		FndRole.RolName:        rec.Name,
		FndRole.RolDescription: rec.Description,
		FndRole.RolIcon:        rec.Icon,
		FndRole.RolUpdatedAt:   now,
		FndRole.RolStatus:      rec.Status,
	}

	query := NewUpdate(T.FndRole).
		Set(record).
		Where(
			goqu.C(FndRole.RolID).Eq(rec.ID),

			goqu.C(FndRole.RolUpdatedAt).Eq(rec.UpdatedAt),
		)
	res, err := DoUpdate(ctx, tx, query)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	rec.UpdatedAt = now
	return CheckOneRowUpdated(ctx, T.FndRole, res)
}

// SetStatus updates the logical status of the record.
func (dao *DaoRole) SetStatus(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.Role,
) error {
	now := time.Now()
	query := NewUpdate(T.FndRole).
		Set(goqu.Record{
			FndRole.RolUpdatedAt: now,
			FndRole.RolStatus:    rec.Status,
		}).
		Where(
			goqu.C(FndRole.RolID).Eq(rec.ID),
			goqu.C(FndRole.RolUpdatedAt).Eq(rec.UpdatedAt),
		)
	res, err := DoUpdate(ctx, tx, query)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	rec.UpdatedAt = now
	return CheckOneRowUpdated(ctx, T.FndRole, res)
}

// Delete performs a physical delete of the record.
func (dao *DaoRole) Delete(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.Role,
) error {
	query := NewDelete(T.FndRole).
		Where(
			goqu.C(FndRole.RolID).Eq(rec.ID),
			goqu.C(FndRole.RolUpdatedAt).Eq(rec.UpdatedAt),
		)
	res, err := DoDelete(ctx, tx, query)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	return CheckOneRowUpdated(ctx, T.FndRole, res)
}

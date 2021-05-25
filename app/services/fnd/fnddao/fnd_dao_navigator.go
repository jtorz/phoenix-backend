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

// DaoNavigator Data Access structure.
type DaoNavigator struct{}

// GetByID retrives the record information using its ID.
func (dao *DaoNavigator) GetByID(ctx context.Context, exe base.Executor,
	id string,
) (*fndmodel.Navigator, error) {
	rec := fndmodel.Navigator{}
	query := NewSelect(T.FndNavigator).
		Select(
			FndNavigator.NavName,
			FndNavigator.NavDescription,
			FndNavigator.NavIcon,
			FndNavigator.NavOrder,
			FndNavigator.NavURL,
			FndNavigator.NavParentID,
			FndNavigator.NavCreatedAt,
			FndNavigator.NavUpdatedAt,
			FndNavigator.NavStatus,
		).
		From(goqu.T(T.FndNavigator)).
		Where(
			goqu.C(FndNavigator.NavID).Eq(id),
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
		&rec.Icon,
		&rec.Order,
		&rec.URL,
		&parentID,
		&rec.CreatedAt,
		&rec.UpdatedAt,
		&rec.Status,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("%s %w", T.FndNavigator, baseerrors.ErrNotFound)
		}
		DebugErr(ctx, err)
		return nil, err
	}
	if parentID != "" {
		rec.Parent = &fndmodel.Navigator{ID: string(parentID)}
	}

	rec.ID = id

	return &rec, nil
}

// List returns the list of records that can be filtered by the user.
func (dao *DaoNavigator) List(ctx context.Context, exe base.Executor,
	qry base.ClientQuery,
) (fndmodel.Navigators, error) {
	res := make(fndmodel.Navigators, 0)
	params, err := ParseClientFilter(qry, fndmodel.Navigator{})
	if err != nil {
		return nil, err
	}
	if qry.OnlyActive {
		params.FilterExp = params.FilterExp.Append(goqu.C(FndNavigator.NavStatus).Eq(base.StatusActive))
	}
	query := NewSelect(T.FndNavigator).
		Select(
			FndNavigator.NavID,
			FndNavigator.NavName,
			FndNavigator.NavDescription,
			FndNavigator.NavIcon,
			FndNavigator.NavOrder,
			FndNavigator.NavURL,
			FndNavigator.NavParentID,
			FndNavigator.NavCreatedAt,
			FndNavigator.NavUpdatedAt,
			FndNavigator.NavStatus,
		).From(T.FndNavigator).
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
		rec := fndmodel.Navigator{}
		err := rows.Scan(
			&rec.ID,
			&rec.Name,
			&rec.Description,
			&rec.Icon,
			&rec.Order,
			&rec.URL,
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
			rec.Parent = &fndmodel.Navigator{ID: string(parentID)}
		}
		res = append(res, rec)
	}
	return res, nil
}

// New creates a new record.
func (dao *DaoNavigator) New(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.Navigator,
) error {
	now := time.Now()
	record := goqu.Record{
		FndNavigator.NavID:          rec.ID,
		FndNavigator.NavName:        rec.Name,
		FndNavigator.NavDescription: rec.Description,
		FndNavigator.NavIcon:        rec.Icon,
		FndNavigator.NavOrder:       rec.Order,
		FndNavigator.NavURL:         rec.URL,
		FndNavigator.NavCreatedAt:   rec.CreatedAt,
		FndNavigator.NavUpdatedAt:   now,
		FndNavigator.NavStatus:      rec.Status,
	}
	if rec.Parent != nil {
		record[FndNavigator.NavParentID] = ZeroString(rec.Parent.ID)
	}
	ins := NewInsert(T.FndNavigator).Rows(record)
	row, err := DoInsertReturning(ctx, tx, ins,
		FndNavigator.NavID,
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
func (dao *DaoNavigator) Edit(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.Navigator,
) error {
	now := time.Now()
	record := goqu.Record{
		FndNavigator.NavName:        rec.Name,
		FndNavigator.NavDescription: rec.Description,
		FndNavigator.NavIcon:        rec.Icon,
		FndNavigator.NavOrder:       rec.Order,
		FndNavigator.NavURL:         rec.URL,
		FndNavigator.NavCreatedAt:   rec.CreatedAt,
		FndNavigator.NavUpdatedAt:   now,
		FndNavigator.NavStatus:      rec.Status,
	}
	if rec.Parent != nil {
		record[FndNavigator.NavParentID] = ZeroString(rec.Parent.ID)
	}

	query := NewUpdate(T.FndNavigator).
		Set(record).
		Where(
			goqu.C(FndNavigator.NavID).Eq(rec.ID),

			goqu.C(FndNavigator.NavUpdatedAt).Eq(rec.UpdatedAt),
		)
	res, err := DoUpdate(ctx, tx, query)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	rec.UpdatedAt = now
	return CheckOneRowUpdated(ctx, T.FndNavigator, res)
}

// SetStatus updates the logical status of the record.
func (dao *DaoNavigator) SetStatus(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.Navigator,
) error {
	now := time.Now()
	query := NewUpdate(T.FndNavigator).
		Set(goqu.Record{
			FndNavigator.NavUpdatedAt: now,
			FndNavigator.NavStatus:    rec.Status,
		}).
		Where(
			goqu.C(FndNavigator.NavID).Eq(rec.ID),
			goqu.C(FndNavigator.NavUpdatedAt).Eq(rec.UpdatedAt),
		)
	res, err := DoUpdate(ctx, tx, query)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	rec.UpdatedAt = now
	return CheckOneRowUpdated(ctx, T.FndNavigator, res)
}

// Delete performs a physical delete of the record.
func (dao *DaoNavigator) Delete(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.Navigator,
) error {
	query := NewDelete(T.FndNavigator).
		Where(
			goqu.C(FndNavigator.NavID).Eq(rec.ID),
			goqu.C(FndNavigator.NavUpdatedAt).Eq(rec.UpdatedAt),
		)
	res, err := DoDelete(ctx, tx, query)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	return CheckOneRowUpdated(ctx, T.FndNavigator, res)
}

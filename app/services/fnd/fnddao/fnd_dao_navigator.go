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

// DaoNavigator Data Access structure.
type DaoNavigator struct {
	Exe base.Executor
}

// GetByID retrives the record information using its ID.
func (dao *DaoNavigator) GetByID(ctx context.Context,
	id string,
) (*fndmodel.Navigator, error) {
	var rec *fndmodel.Navigator
	query := NewSelect(T.FndNavigator).
		Select(
			FndNavigator.NavName,
			FndNavigator.NavDescription,
			FndNavigator.NavIcon,
			FndNavigator.NavOrder,
			FndNavigator.NavURL,
			FndNavigator.NavStatus,
			FndNavigator.NavUpdatedAt,
		).
		Where(
			goqu.C(FndNavigator.NavID).Eq(id),
		)

	row, err := QueryRowContext(ctx, dao.Exe, query)
	if err != nil {
		return nil, WrapErr(ctx, err)
	}
	err = row.Scan(
		&rec.Name,
		&rec.Description,
		&rec.Icon,
		&rec.Order,
		&rec.URL,
		&rec.Status,
		&rec.UpdatedAt,
	)
	if err != nil {
		return nil, WrapErr(ctx, err)
	}
	return rec, nil
}

// All returns the the complete catalogue of records that can be filtered by the user.
func (dao *DaoNavigator) All(ctx context.Context,
	qry base.ClientQuery,
) ([]fndmodel.Navigator, error) {
	res := make([]fndmodel.Navigator, 0)
	params, err := FndRole.ParseFilter(qry)
	if err != nil {
		return nil, err
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
			FndNavigator.NavStatus,
			FndNavigator.NavUpdatedAt,
		).Where(params.FilterExp).
		Limit(params.Limit).
		Limit(params.Offset).
		Order(params.Sort...)

	rows, err := QueryContext(ctx, dao.Exe, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, nil
		}
		return nil, WrapErr(ctx, err)
	}
	defer rows.Close()
	for rows.Next() {
		var parentID string
		rec := fndmodel.Navigator{}
		err := rows.Scan(
			&rec.ID,
			&rec.Name,
			&rec.Description,
			&rec.Icon,
			&rec.Order,
			&rec.URL,
			&parentID,
			&rec.Status,
			&rec.UpdatedAt,
		)
		if err != nil {
			return nil, WrapErr(ctx, err)
		}
		if parentID != "" {
			rec.Parent = &fndmodel.Navigator{ID: parentID}
		}
		res = append(res, rec)
	}
	return res, nil
}

// List returns the list of records that can be filtered by the user.
func (dao *DaoNavigator) List(ctx context.Context,
	qry base.ClientQuery,
) ([]fndmodel.Navigator, error) {
	res := make([]fndmodel.Navigator, 0)
	params, err := FndRole.ParseFilter(qry)
	if err != nil {
		return nil, err
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
			FndNavigator.NavStatus,
			FndNavigator.NavUpdatedAt,
		).Where(params.FilterExp).
		Limit(params.Limit).
		Limit(params.Offset).
		Order(params.Sort...)

	rows, err := QueryContext(ctx, dao.Exe, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, nil
		}
		return nil, WrapErr(ctx, err)
	}
	defer rows.Close()
	for rows.Next() {
		var parentID string
		rec := fndmodel.Navigator{}
		err := rows.Scan(
			&rec.ID,
			&rec.Name,
			&rec.Description,
			&rec.Icon,
			&rec.Order,
			&rec.URL,
			&parentID,
			&rec.Status,
			&rec.UpdatedAt,
		)
		if err != nil {
			return nil, WrapErr(ctx, err)
		}
		if parentID != "" {
			rec.Parent = &fndmodel.Navigator{ID: parentID}
		}
		res = append(res, rec)
	}
	return res, nil
}

// New creates a new record.
func (dao *DaoNavigator) New(ctx context.Context,
	rec *fndmodel.Navigator,
) error {
	now := time.Now()
	record := goqu.Record{
		FndNavigator.NavName:        rec.Name,
		FndNavigator.NavDescription: rec.Description,
		FndNavigator.NavIcon:        rec.Icon,
		FndNavigator.NavOrder:       rec.Order,
		FndNavigator.NavURL:         rec.URL,
		FndNavigator.NavStatus:      rec.Status,
		FndNavigator.NavUpdatedAt:   now,
	}
	if rec.Parent != nil {
		if rec.Parent.ID != "" {
			record[FndNavigator.NavParentID] = rec.Parent.ID
		}
	}
	ins := NewInsert(T.FndNavigator).Rows(record)
	row, err := DoInsertReturning(ctx, dao.Exe, ins, FndNavigator.NavID)
	if err != nil {
		return WrapErr(ctx, err)
	}
	err = row.Scan(&rec.ID)
	if err != nil {
		return WrapErr(ctx, err)
	}
	rec.UpdatedAt = now
	return nil
}

// Edit edits the record.
func (dao *DaoNavigator) Edit(ctx context.Context,
	rec *fndmodel.Navigator,
) error {
	now := time.Now()

	record := goqu.Record{
		FndNavigator.NavName:        rec.Name,
		FndNavigator.NavDescription: rec.Description,
		FndNavigator.NavIcon:        rec.Icon,
		FndNavigator.NavOrder:       rec.Order,
		FndNavigator.NavURL:         rec.URL,
		FndNavigator.NavStatus:      rec.Status,
		FndNavigator.NavUpdatedAt:   now,
	}
	if rec.Parent != nil {
		if rec.Parent.ID != "" {
			record[FndNavigator.NavParentID] = rec.Parent.ID
		}
	}
	query := NewUpdate(T.FndNavigator).
		Set(record).
		Where(
			goqu.C(FndNavigator.NavID).Eq(rec.ID),
			goqu.C(FndNavigator.NavUpdatedAt).Eq(rec.UpdatedAt),
		)
	res, err := DoUpdate(ctx, dao.Exe, query)
	if err != nil {
		return WrapErr(ctx, err)
	}
	rec.UpdatedAt = now
	return CheckOneRowUpdated(ctx, res)
}

// SetStatus updates the logical status of the record.
func (dao *DaoNavigator) SetStatus(ctx context.Context,
	rec *fndmodel.Navigator,
) error {
	now := time.Now()
	query := NewUpdate(T.FndNavigator).
		Set(goqu.Record{
			FndNavigator.NavStatus:    rec.Status,
			FndNavigator.NavUpdatedAt: now,
		}).
		Where(
			goqu.C(FndNavigator.NavID).Eq(rec.ID),
			goqu.C(FndNavigator.NavUpdatedAt).Eq(rec.UpdatedAt),
		)
	res, err := DoUpdate(ctx, dao.Exe, query)
	if err != nil {
		return WrapErr(ctx, err)
	}
	rec.UpdatedAt = now
	return CheckOneRowUpdated(ctx, res)
}

// Delete physical delete of the record.
func (dao *DaoNavigator) Delete(ctx context.Context,
	rec *fndmodel.Navigator,
) error {
	query := NewDelete(T.FndNavigator).
		Where(
			goqu.C(FndNavigator.NavID).Eq(rec.ID),
			goqu.C(FndNavigator.NavUpdatedAt).Eq(rec.UpdatedAt),
		)
	res, err := DoDelete(ctx, dao.Exe, query)
	if err != nil {
		return WrapErr(ctx, err)
	}
	return CheckOneRowUpdated(ctx, res)
}

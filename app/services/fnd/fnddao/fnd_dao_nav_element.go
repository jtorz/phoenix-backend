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

// DaoNavElement Data Access structure.
type DaoNavElement struct{}

// GetByID retrives the record information using its ID.
func (dao *DaoNavElement) GetByID(ctx context.Context, exe base.Executor,
	id string,
) (*fndmodel.NavElement, error) {
	rec := fndmodel.NavElement{}
	query := NewSelect(T.FndNavElement).
		Select(
			FndNavElement.NaeName,
			FndNavElement.NaeDescription,
			FndNavElement.NaeIcon,
			FndNavElement.NaeOrder,
			FndNavElement.NaeURL,
			FndNavElement.NaeParentID,
			FndNavElement.NaeCreatedAt,
			FndNavElement.NaeUpdatedAt,
			FndNavElement.NaeStatus,
		).
		From(goqu.T(T.FndNavElement)).
		Where(
			goqu.C(FndNavElement.NaeID).Eq(id),
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
		DebugErr(ctx, err)
		return nil, err
	}
	rec.ParentID = string(parentID)

	rec.ID = id

	return &rec, nil
}

// ListAll returns the all the navigator records with the flag IsAssigned
// if the navigator element is assigned to the role.
func (dao *DaoNavElement) ListAll(ctx context.Context, exe base.Executor,
	rolID string,
) (fndmodel.Navigator, error) {
	res := make(fndmodel.Navigator, 0)

	query := NewSelect(
		FndNavElement.NaeID,
		FndNavElement.NaeName,
		FndNavElement.NaeDescription,
		FndNavElement.NaeIcon,
		FndNavElement.NaeOrder,
		FndNavElement.NaeURL,
		FndNavElement.NaeParentID,
		FndNavElement.NaeCreatedAt,
		FndNavElement.NaeUpdatedAt,
		FndNavElement.NaeStatus,
		goqu.C(FndNavElementRole.NerRoleID).IsNotNull(),
	).WithRecursive(
		"navigators",
		NewSelect(goqu.T(T.FndNavElement).All(), goqu.L("0").As("level")).
			From(T.FndNavElement).
			Where(goqu.C(FndNavElement.NaeParentID).IsNull()).
			Union(
				NewSelect(goqu.T(T.FndNavElement).As("child").All(), goqu.L("level+1")).
					From(goqu.T(T.FndNavElement).As("child")).
					InnerJoin(goqu.T("navigators"),
						goqu.On(goqu.C(FndNavElement.NaeID).Table("navigators").Eq(goqu.I(FndNavElement.NaeParentID).Table("child")))),
			),
	).
		From("navigators").
		LeftJoin(goqu.T(T.FndNavElementRole), FndNavElementRoleFkFndNavElement(
			goqu.C(FndNavElementRole.NerRoleID).Eq(rolID),
		)).Order(goqu.C("level").Asc())

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
		rec := fndmodel.NavElement{Children: fndmodel.Navigator{}}
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
			&rec.IsAssigned,
		)
		if err != nil {
			DebugErr(ctx, err)
			return nil, err
		}
		rec.ParentID = string(parentID)
		res = append(res, rec)
	}
	return res, nil
}

// New creates a new record.
func (dao *DaoNavElement) New(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.NavElement,
) error {
	now := time.Now()
	record := goqu.Record{
		FndNavElement.NaeID:          rec.ID,
		FndNavElement.NaeName:        rec.Name,
		FndNavElement.NaeDescription: rec.Description,
		FndNavElement.NaeIcon:        rec.Icon,
		FndNavElement.NaeOrder:       rec.Order,
		FndNavElement.NaeURL:         rec.URL,
		FndNavElement.NaeParentID:    ZeroString(rec.ParentID),
		FndNavElement.NaeCreatedAt:   rec.CreatedAt,
		FndNavElement.NaeUpdatedAt:   now,
		FndNavElement.NaeStatus:      rec.Status,
	}
	ins := NewInsert(T.FndNavElement).Rows(record)
	row, err := DoInsertReturning(ctx, tx, ins,
		FndNavElement.NaeID,
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
func (dao *DaoNavElement) Edit(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.NavElement,
) error {
	now := time.Now()
	record := goqu.Record{
		FndNavElement.NaeName:        rec.Name,
		FndNavElement.NaeDescription: rec.Description,
		FndNavElement.NaeIcon:        rec.Icon,
		FndNavElement.NaeOrder:       rec.Order,
		FndNavElement.NaeURL:         rec.URL,
		FndNavElement.NaeParentID:    ZeroString(rec.ParentID),
		FndNavElement.NaeUpdatedAt:   now,
		FndNavElement.NaeStatus:      rec.Status,
	}

	query := NewUpdate(T.FndNavElement).
		Set(record).
		Where(
			goqu.C(FndNavElement.NaeID).Eq(rec.ID),

			goqu.C(FndNavElement.NaeUpdatedAt).Eq(rec.UpdatedAt),
		)
	res, err := DoUpdate(ctx, tx, query)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	rec.UpdatedAt = now
	return CheckOneRowUpdated(ctx, T.FndNavElement, res)
}

// SetStatus updates the logical status of the record.
func (dao *DaoNavElement) SetStatus(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.NavElement,
) error {
	now := time.Now()
	query := NewUpdate(T.FndNavElement).
		Set(goqu.Record{
			FndNavElement.NaeUpdatedAt: now,
			FndNavElement.NaeStatus:    rec.Status,
		}).
		Where(
			goqu.C(FndNavElement.NaeID).Eq(rec.ID),
			goqu.C(FndNavElement.NaeUpdatedAt).Eq(rec.UpdatedAt),
		)
	res, err := DoUpdate(ctx, tx, query)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	rec.UpdatedAt = now
	return CheckOneRowUpdated(ctx, T.FndNavElement, res)
}

// Delete performs a physical delete of the record.
func (dao *DaoNavElement) Delete(ctx context.Context, tx *sql.Tx,
	rec *fndmodel.NavElement,
) error {
	query := NewDelete(T.FndNavElement).
		Where(
			goqu.C(FndNavElement.NaeID).Eq(rec.ID),
			goqu.C(FndNavElement.NaeUpdatedAt).Eq(rec.UpdatedAt),
		)
	res, err := DoDelete(ctx, tx, query)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	return CheckOneRowUpdated(ctx, T.FndNavElement, res)
}

// AssignToRole assigns the nav element to the role.
func (dao *DaoNavElement) AssignToRole(ctx context.Context, tx *sql.Tx,
	elementID, roleId string,
) error {
	record := goqu.Record{
		FndNavElementRole.NerRoleID:       elementID,
		FndNavElementRole.NerNavElementID: roleId,
	}
	ins := NewInsert(T.FndNavElementRole).Rows(record)
	_, err := DoInsert(ctx, tx, ins)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	return nil
}

// AssociateRole associates the nav element to the role.
func (dao *DaoNavElement) AssociateRole(ctx context.Context, tx *sql.Tx,
	elementID, roleID string,
) error {
	record := goqu.Record{
		FndNavElementRole.NerNavElementID: elementID,
		FndNavElementRole.NerRoleID:       roleID,
	}
	ins := NewInsert(T.FndNavElementRole).Rows(record)
	_, err := DoInsert(ctx, tx, ins)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	return nil
}

// DissociateRole dissociates the nav element from the role.
func (dao *DaoNavElement) DissociateRole(ctx context.Context, tx *sql.Tx,
	elementID, roleID string,
) error {
	query := NewDelete(T.FndNavElementRole).
		Where(
			goqu.C(FndNavElementRole.NerNavElementID).Eq(elementID),
			goqu.C(FndNavElementRole.NerRoleID).Eq(roleID),
		)
	res, err := DoDelete(ctx, tx, query)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	return CheckOneRowUpdated(ctx, T.FndNavElementRole, res)
}

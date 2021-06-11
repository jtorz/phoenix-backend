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

// DalNavElement Data Access structure.
type DalNavElement struct{}

// GetByID retrives the record information of the nav element using its ID.
func (dal *DalNavElement) GetByID(ctx context.Context, exe base.Executor,
	id string,
) (*coremodel.NavElement, error) {
	rec := coremodel.NavElement{}
	query := NewSelect(T.CoreNavElement).
		Select(
			CoreNavElement.NaeName,
			CoreNavElement.NaeDescription,
			CoreNavElement.NaeIcon,
			CoreNavElement.NaeOrder,
			CoreNavElement.NaeURL,
			CoreNavElement.NaeParentID,
			CoreNavElement.NaeCreatedAt,
			CoreNavElement.NaeUpdatedAt,
			CoreNavElement.NaeStatus,
		).
		From(goqu.T(T.CoreNavElement)).
		Where(
			goqu.C(CoreNavElement.NaeID).Eq(id),
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
		return nil, WrapNotFound(ctx, T.CoreNavElement, err)
	}
	rec.ParentID = string(parentID)

	rec.ID = id

	return &rec, nil
}

// ListAll returns the all the navigator records with the flag IsAssigned
// if the navigator element is assigned to the role.
func (dal *DalNavElement) ListAll(ctx context.Context, exe base.Executor,
	rolID string,
) (coremodel.Navigator, error) {
	res := make(coremodel.Navigator, 0)

	query := NewSelect(
		CoreNavElement.NaeID,
		CoreNavElement.NaeName,
		CoreNavElement.NaeDescription,
		CoreNavElement.NaeIcon,
		CoreNavElement.NaeOrder,
		CoreNavElement.NaeURL,
		CoreNavElement.NaeParentID,
		CoreNavElement.NaeCreatedAt,
		CoreNavElement.NaeUpdatedAt,
		CoreNavElement.NaeStatus,
		goqu.C(CoreNavElementRole.NerRoleID).IsNotNull(),
	).WithRecursive(
		"navigators",
		NewSelect(goqu.T(T.CoreNavElement).All(), goqu.L("0").As("level")).
			From(T.CoreNavElement).
			Where(goqu.C(CoreNavElement.NaeParentID).IsNull()).
			Union(
				NewSelect(goqu.T(T.CoreNavElement).As("child").All(), goqu.L("level+1")).
					From(goqu.T(T.CoreNavElement).As("child")).
					InnerJoin(goqu.T("navigators"),
						goqu.On(goqu.C(CoreNavElement.NaeID).Table("navigators").Eq(goqu.I(CoreNavElement.NaeParentID).Table("child")))),
			),
	).
		From("navigators").
		LeftJoin(goqu.T(T.CoreNavElementRole), CoreNavElementRoleFkCoreNavElement(
			goqu.C(CoreNavElementRole.NerRoleID).Eq(rolID),
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
		rec := coremodel.NavElement{Children: coremodel.Navigator{}}
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
func (dal *DalNavElement) New(ctx context.Context, tx *sql.Tx,
	rec *coremodel.NavElement,
) error {
	now := time.Now()
	record := goqu.Record{
		CoreNavElement.NaeID:          rec.ID,
		CoreNavElement.NaeName:        rec.Name,
		CoreNavElement.NaeDescription: rec.Description,
		CoreNavElement.NaeIcon:        rec.Icon,
		CoreNavElement.NaeOrder:       rec.Order,
		CoreNavElement.NaeURL:         rec.URL,
		CoreNavElement.NaeParentID:    ZeroString(rec.ParentID),
		CoreNavElement.NaeCreatedAt:   rec.CreatedAt,
		CoreNavElement.NaeUpdatedAt:   now,
		CoreNavElement.NaeStatus:      rec.Status,
	}
	ins := NewInsert(T.CoreNavElement).Rows(record)
	row, err := DoInsertReturning(ctx, tx, ins,
		CoreNavElement.NaeID,
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
func (dal *DalNavElement) Edit(ctx context.Context, tx *sql.Tx,
	rec *coremodel.NavElement,
) error {
	now := time.Now()
	record := goqu.Record{
		CoreNavElement.NaeName:        rec.Name,
		CoreNavElement.NaeDescription: rec.Description,
		CoreNavElement.NaeIcon:        rec.Icon,
		CoreNavElement.NaeOrder:       rec.Order,
		CoreNavElement.NaeURL:         rec.URL,
		CoreNavElement.NaeParentID:    ZeroString(rec.ParentID),
		CoreNavElement.NaeUpdatedAt:   now,
		CoreNavElement.NaeStatus:      rec.Status,
	}

	query := NewUpdate(T.CoreNavElement).
		Set(record).
		Where(
			goqu.C(CoreNavElement.NaeID).Eq(rec.ID),

			goqu.C(CoreNavElement.NaeUpdatedAt).Eq(rec.UpdatedAt),
		)
	res, err := DoUpdate(ctx, tx, query)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	rec.UpdatedAt = now
	return CheckOneRowUpdated(ctx, T.CoreNavElement, res)
}

// SetStatus updates the logical status of the record.
func (dal *DalNavElement) SetStatus(ctx context.Context, tx *sql.Tx,
	rec *coremodel.NavElement,
) error {
	now := time.Now()
	query := NewUpdate(T.CoreNavElement).
		Set(goqu.Record{
			CoreNavElement.NaeUpdatedAt: now,
			CoreNavElement.NaeStatus:    rec.Status,
		}).
		Where(
			goqu.C(CoreNavElement.NaeID).Eq(rec.ID),
			goqu.C(CoreNavElement.NaeUpdatedAt).Eq(rec.UpdatedAt),
		)
	res, err := DoUpdate(ctx, tx, query)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	rec.UpdatedAt = now
	return CheckOneRowUpdated(ctx, T.CoreNavElement, res)
}

// Delete performs a physical delete of the record.
func (dal *DalNavElement) Delete(ctx context.Context, tx *sql.Tx,
	rec *coremodel.NavElement,
) error {
	query := NewDelete(T.CoreNavElement).
		Where(
			goqu.C(CoreNavElement.NaeID).Eq(rec.ID),
			goqu.C(CoreNavElement.NaeUpdatedAt).Eq(rec.UpdatedAt),
		)
	res, err := DoDelete(ctx, tx, query)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	return CheckOneRowUpdated(ctx, T.CoreNavElement, res)
}

// AssignToRole assigns the nav element to the role.
func (dal *DalNavElement) AssignToRole(ctx context.Context, tx *sql.Tx,
	elementID, roleId string,
) error {
	record := goqu.Record{
		CoreNavElementRole.NerRoleID:       elementID,
		CoreNavElementRole.NerNavElementID: roleId,
	}
	ins := NewInsert(T.CoreNavElementRole).Rows(record)
	_, err := DoInsert(ctx, tx, ins)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	return nil
}

// AssociateRole associates the nav element to the role.
func (dal *DalNavElement) AssociateRole(ctx context.Context, tx *sql.Tx,
	elementID, roleID string,
) error {
	record := goqu.Record{
		CoreNavElementRole.NerNavElementID: elementID,
		CoreNavElementRole.NerRoleID:       roleID,
	}
	ins := NewInsert(T.CoreNavElementRole).Rows(record)
	_, err := DoInsert(ctx, tx, ins)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	return nil
}

// DissociateRole dissociates the nav element from the role.
func (dal *DalNavElement) DissociateRole(ctx context.Context, tx *sql.Tx,
	elementID, roleID string,
) error {
	query := NewDelete(T.CoreNavElementRole).
		Where(
			goqu.C(CoreNavElementRole.NerNavElementID).Eq(elementID),
			goqu.C(CoreNavElementRole.NerRoleID).Eq(roleID),
		)
	res, err := DoDelete(ctx, tx, query)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	return CheckOneRowUpdated(ctx, T.CoreNavElementRole, res)
}

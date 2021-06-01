package fnddal

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/jtorz/phoenix-backend/app/services/fnd/fndmodel"
	"github.com/jtorz/phoenix-backend/app/shared/base"
	"github.com/jtorz/phoenix-backend/app/shared/baseerrors"

	//lint:ignore ST1001 dot import allowed only in dal packages for lex.
	. "github.com/jtorz/phoenix-backend/app/shared/lex"
)

// DalAccountAccess Data Access structure.
type DalAccountAccess struct{}

func (dal *DalAccountAccess) Insert(ctx context.Context, exe *sql.Tx,
	ac *fndmodel.AccountAccess,
) error {
	ins := NewInsert(T.FndAccountAccess).Rows(goqu.Record{
		FndAccountAccess.AcaID:     ac.Key,
		FndAccountAccess.AcaType:   ac.Type,
		FndAccountAccess.AcaUserID: ac.User.ID,
		FndAccountAccess.AcaStatus: ac.Status,
	})
	_, err := DoInsert(ctx, exe, ins)
	DebugErr(ctx, err)
	return err
}

func (dal *DalAccountAccess) UseAccountAccess(ctx context.Context, exe *sql.Tx,
	key string, accType fndmodel.AccountAccessType,
) (string, error) {
	query := NewUpdate(T.FndAccountAccess).
		Set(goqu.Record{
			FndAccountAccess.AcaStatus: base.StatusInactive,
		}).
		Where(
			goqu.C(FndAccountAccess.AcaID).Eq(key),
			goqu.C(FndAccountAccess.AcaType).Eq(accType),
			goqu.C(FndAccountAccess.AcaStatus).Eq(base.StatusActive),
			goqu.C(FndAccountAccess.AcaExpirationDate).Gt(goqu.L("CURRENT_TIMESTAMP")),
		)

	row, err := DoUpdateReturningRow(ctx, exe, query, FndAccountAccess.AcaUserID)

	if err != nil {
		DebugErr(ctx, err)
		return "", err
	}
	var userID string
	err = row.Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("account access %w", baseerrors.ErrNotFound)
		}
	}
	DebugErr(ctx, err)
	return userID, err
}

func (dal *DalAccountAccess) GetAccessByUserID(ctx context.Context, exe base.Executor,
	userID string, accType fndmodel.AccountAccessType,
) (*fndmodel.AccountAccess, error) {
	res := fndmodel.AccountAccess{}
	query := NewSelect(FndAccountAccess.AcaID).
		From(T.FndAccountAccess).
		Where(
			goqu.C(FndAccountAccess.AcaUserID).Eq(userID),
			goqu.C(FndAccountAccess.AcaType).Eq(accType),
			goqu.C(FndAccountAccess.AcaStatus).Eq(base.StatusActive),
			goqu.C(FndAccountAccess.AcaExpirationDate).Lt(goqu.L("CURRENT_TIMESTAMP")),
		)

	row, err := QueryRowContext(ctx, exe, query)
	if err != nil {
		DebugErr(ctx, err)
		return nil, err
	}
	err = row.Scan(&res.Key)
	if err != nil {
		DebugErr(ctx, err)
		return nil, err
	}
	return &res, nil
}

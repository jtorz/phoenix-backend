package coredal

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/jtorz/phoenix-backend/app/services/core/coremodel"
	"github.com/jtorz/phoenix-backend/app/shared/base"
	"github.com/jtorz/phoenix-backend/app/shared/baseerrors"

	//lint:ignore ST1001 dot import allowed only in dal packages for lex.
	. "github.com/jtorz/phoenix-backend/app/shared/lex"
)

// DalAccountAccess Data Access structure.
type DalAccountAccess struct{}

func (dal *DalAccountAccess) Insert(ctx context.Context, exe *sql.Tx,
	ac *coremodel.AccountAccess,
) error {
	ins := NewInsert(T.CoreAccountAccess).Rows(goqu.Record{
		CoreAccountAccess.AcaID:     ac.Key,
		CoreAccountAccess.AcaType:   ac.Type,
		CoreAccountAccess.AcaUserID: ac.User.ID,
		CoreAccountAccess.AcaStatus: ac.Status,
	})
	_, err := DoInsert(ctx, exe, ins)
	DebugErr(ctx, err)
	return err
}

func (dal *DalAccountAccess) UseAccountAccess(ctx context.Context, exe *sql.Tx,
	key string, accType coremodel.AccountAccessType,
) (string, error) {
	query := NewUpdate(T.CoreAccountAccess).
		Set(goqu.Record{
			CoreAccountAccess.AcaStatus: base.StatusInactive,
		}).
		Where(
			goqu.C(CoreAccountAccess.AcaID).Eq(key),
			goqu.C(CoreAccountAccess.AcaType).Eq(accType),
			goqu.C(CoreAccountAccess.AcaStatus).Eq(base.StatusActive),
			goqu.C(CoreAccountAccess.AcaExpirationDate).Gt(goqu.L("CURRENT_TIMESTAMP")),
		)

	row, err := DoUpdateReturningRow(ctx, exe, query, CoreAccountAccess.AcaUserID)

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
	userID string, accType coremodel.AccountAccessType,
) (*coremodel.AccountAccess, error) {
	res := coremodel.AccountAccess{}
	query := NewSelect(CoreAccountAccess.AcaID).
		From(T.CoreAccountAccess).
		Where(
			goqu.C(CoreAccountAccess.AcaUserID).Eq(userID),
			goqu.C(CoreAccountAccess.AcaType).Eq(accType),
			goqu.C(CoreAccountAccess.AcaStatus).Eq(base.StatusActive),
			goqu.C(CoreAccountAccess.AcaExpirationDate).Lt(goqu.L("CURRENT_TIMESTAMP")),
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

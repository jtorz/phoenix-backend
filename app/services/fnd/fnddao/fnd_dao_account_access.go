package fnddao

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/jtorz/phoenix-backend/app/services/fnd/fndmodel"
	"github.com/jtorz/phoenix-backend/app/shared/base"
	"github.com/jtorz/phoenix-backend/app/shared/baseerrors"

	//lint:ignore ST1001  only lex package should be used as a odot import.
	. "github.com/jtorz/phoenix-backend/app/shared/lex"
)

// DaoAccountAccess Data Access structure.
type DaoAccountAccess struct {
	Exe base.Executor
}

func (dao *DaoAccountAccess) Insert(ctx context.Context,
	ac *fndmodel.AccountAccess,
) error {
	ins := NewInsert(T.FndAccountAccess).Rows(goqu.Record{
		FndAccountAccess.AcaID:     ac.Key,
		FndAccountAccess.AcaType:   ac.Type,
		FndAccountAccess.AcaUserID: ac.User.ID,
		FndAccountAccess.AcaStatus: ac.Status,
	})
	_, err := DoInsert(ctx, dao.Exe, ins)
	return WrapErr(ctx, err)
}

func (dao *DaoAccountAccess) UseAccountAccess(ctx context.Context,
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

	row, err := DoUpdateReturningRow(ctx, dao.Exe, query, FndAccountAccess.AcaUserID)

	if err != nil {
		return "", WrapErr(ctx, err)
	}
	var userID string
	err = row.Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("account access %w", baseerrors.ErrNotFound)
		}
	}
	return userID, WrapErr(ctx, err)
}

func (dao *DaoAccountAccess) GetAccessByUserID(ctx context.Context,
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

	row, err := QueryRowContext(ctx, dao.Exe, query)
	if err != nil {
		return nil, WrapErr(ctx, err)
	}
	err = row.Scan(&res.Key)
	if err != nil {
		return nil, WrapErr(ctx, err)
	}
	return &res, nil
}

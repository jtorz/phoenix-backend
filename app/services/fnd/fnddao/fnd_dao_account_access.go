package fnddao

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/jtorz/phoenix-backend/app/services/fnd/fndmodel"
	"github.com/jtorz/phoenix-backend/app/shared/base"
	"github.com/jtorz/phoenix-backend/app/shared/baseerrors"
	"github.com/jtorz/phoenix-backend/app/shared/daohelper"
	"github.com/jtorz/phoenix-backend/app/shared/lex"
)

// DaoAccountAccess Data Access structure.
type DaoAccountAccess struct {
	Exe base.Executor
	h   daohelper.QueryHelper
}

func (dao *DaoAccountAccess) Insert(ctx context.Context,
	ac *fndmodel.AccountAccess,
) error {
	ins := dao.h.NewInsert(lex.T.FndAccessAccount).Rows(goqu.Record{
		lex.FndtaccessAccount.AcaID:     ac.Key,
		lex.FndtaccessAccount.AcaType:   ac.Type,
		lex.FndtaccessAccount.AcaUserID: ac.User.ID,
		lex.FndtaccessAccount.AcaStatus: ac.Status,
	})
	_, err := dao.h.DoInsert(ctx, dao.Exe, ins)
	return dao.h.WrapErr(err)
}

func (dao *DaoAccountAccess) UseAccountAccess(ctx context.Context,
	accType fndmodel.AccountAccessType, key string,
) (string, error) {
	query := dao.h.NewUpdate(lex.T.FndAccessAccount).
		Set(goqu.Record{
			lex.FndtaccessAccount.AcaStatus: base.StatusInactive,
		}).
		Where(
			goqu.C(lex.FndtaccessAccount.AcaID).Eq(key),
			goqu.C(lex.FndtaccessAccount.AcaType).Eq(accType),
			goqu.C(lex.FndtaccessAccount.AcaStatus).Eq(base.StatusActive),
			goqu.C(lex.FndtaccessAccount.AcaExpirationDate).Gt(goqu.L("CURRENT_TIMESTAMP")),
		)

	row, err := dao.h.DoUpdateReturningRow(ctx, dao.Exe, query, lex.FndtaccessAccount.AcaUserID)

	if err != nil {
		return "", dao.h.WrapErr(err)
	}
	var userID string
	err = row.Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("account access %w", baseerrors.ErrNotFound)
		}
	}
	return userID, dao.h.WrapErr(err)
}

func (dao *DaoAccountAccess) GetAccessByUserID(ctx context.Context,
	accType fndmodel.AccountAccessType, userID string,
) (*fndmodel.AccountAccess, error) {
	res := fndmodel.AccountAccess{}
	query := dao.h.NewSelect(lex.T.FndAccessAccount).
		Select(lex.FndtaccessAccount.AcaID).
		Where(
			goqu.C(lex.FndtaccessAccount.AcaUserID).Eq(userID),
			goqu.C(lex.FndtaccessAccount.AcaType).Eq(accType),
			goqu.C(lex.FndtaccessAccount.AcaStatus).Eq(base.StatusActive),
			goqu.C(lex.FndtaccessAccount.AcaExpirationDate).Lt(goqu.L("CURRENT_TIMESTAMP")),
		)

	row, err := dao.h.QueryRowContext(ctx, dao.Exe, query)
	if err != nil {
		return nil, dao.h.WrapErr(err)
	}
	err = row.Scan(&res.Key)
	if err != nil {
		return nil, dao.h.WrapErr(err)
	}
	return &res, nil
}

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
	ins := dao.h.NewInsert(lex.T.FndAccountAccess).Rows(goqu.Record{
		lex.FndAccountAccess.AcaID:     ac.Key,
		lex.FndAccountAccess.AcaType:   ac.Type,
		lex.FndAccountAccess.AcaUserID: ac.User.ID,
		lex.FndAccountAccess.AcaStatus: ac.Status,
	})
	_, err := dao.h.DoInsert(ctx, dao.Exe, ins)
	return dao.h.WrapErr(err)
}

func (dao *DaoAccountAccess) UseAccountAccess(ctx context.Context,
	key string, accType fndmodel.AccountAccessType,
) (string, error) {
	query := dao.h.NewUpdate(lex.T.FndAccountAccess).
		Set(goqu.Record{
			lex.FndAccountAccess.AcaStatus: base.StatusInactive,
		}).
		Where(
			goqu.C(lex.FndAccountAccess.AcaID).Eq(key),
			goqu.C(lex.FndAccountAccess.AcaType).Eq(accType),
			goqu.C(lex.FndAccountAccess.AcaStatus).Eq(base.StatusActive),
			goqu.C(lex.FndAccountAccess.AcaExpirationDate).Gt(goqu.L("CURRENT_TIMESTAMP")),
		)

	row, err := dao.h.DoUpdateReturningRow(ctx, dao.Exe, query, lex.FndAccountAccess.AcaUserID)

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
	userID string, accType fndmodel.AccountAccessType,
) (*fndmodel.AccountAccess, error) {
	res := fndmodel.AccountAccess{}
	query := dao.h.NewSelect(lex.T.FndAccountAccess).
		Select(lex.FndAccountAccess.AcaID).
		Where(
			goqu.C(lex.FndAccountAccess.AcaUserID).Eq(userID),
			goqu.C(lex.FndAccountAccess.AcaType).Eq(accType),
			goqu.C(lex.FndAccountAccess.AcaStatus).Eq(base.StatusActive),
			goqu.C(lex.FndAccountAccess.AcaExpirationDate).Lt(goqu.L("CURRENT_TIMESTAMP")),
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

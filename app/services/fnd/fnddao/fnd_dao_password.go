package fnddao

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/jtorz/phoenix-backend/app/services/fnd/fndmodel"
	"github.com/jtorz/phoenix-backend/app/shared/base"
	"github.com/jtorz/phoenix-backend/app/shared/daohelper"
	"github.com/jtorz/phoenix-backend/app/shared/lex"
)

// DaoPassword Data Access structure.
type DaoPassword struct {
	Exe base.Executor
	h   daohelper.QueryHelper
}

/// New saves the password record in the database.
func (dao *DaoPassword) New(ctx context.Context, usuarioID string, p *fndmodel.Password) error {
	jsonData, err := json.Marshal(p.Data)
	if err != nil {
		return err
	}
	now := time.Now()
	ins := dao.h.NewInsert(lex.T.FndPassword).Rows(goqu.Record{
		lex.FndPassword.PasData:      jsonData,
		lex.FndPassword.PasUserID:    usuarioID,
		lex.FndPassword.PasType:      p.Type,
		lex.FndPassword.PasStatus:    p.Status,
		lex.FndPassword.PasUpdatedAt: now,
	})
	row, err := dao.h.DoInsertReturning(ctx, dao.Exe, ins, lex.FndPassword.PasID)
	if err != nil {
		return dao.h.WrapErr(err)
	}
	if err := row.Scan(&p.ID); err != nil {
		return err
	}
	p.UpdatedAt = now
	return nil
}

func (dao *DaoPassword) asd(ctx context.Context) ([]fndmodel.User, error) {
	res := make([]fndmodel.User, 0)
	query := dao.h.NewSelect(lex.T.FndUser).
		Select(
			lex.FndUser.UseEmail,
		).
		Where(
			goqu.C(lex.FndUser.UseID).Eq(2),
		)

	rows, err := dao.h.QueryContext(ctx, dao.Exe, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, nil
		}
		return nil, dao.h.WrapErr(err)
	}
	defer rows.Close()
	for rows.Next() {
		rec := fndmodel.User{}
		err := rows.Scan(
			&rec.Email,
		)
		if err != nil {
			return nil, dao.h.WrapErr(err)
		}
		res = append(res, rec)
	}
	return res, nil

}

package fnddao

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/jtorz/phoenix-backend/app/services/fnd/fndmodel"
	"github.com/jtorz/phoenix-backend/app/shared/base"
	"github.com/jtorz/phoenix-backend/app/shared/baseerrors"
	"github.com/jtorz/phoenix-backend/app/shared/daohelper"
	"github.com/jtorz/phoenix-backend/app/shared/lex"
)

// DaoUser Data Access structure.
type DaoUser struct {
	h daohelper.QueryHelper
}

// Login retrieves the necessary data to login a user given its email/username.
func (dao DaoUser) Login(ctx context.Context, exe base.Executor,
	user string,
) (*fndmodel.User, error) {
	res := &fndmodel.User{Password: &fndmodel.Password{}}
	query := dao.h.NewSelect(lex.T.Fndtuser).
		Select(
			lex.Fndtuser.UseID,
			lex.Fndtuser.UseName,
			lex.Fndtuser.UseMiddleName,
			lex.Fndtuser.UseLastName,
			lex.Fndtuser.UseEmail,
			lex.Fndtuser.UseUsername,
			lex.Fndtuser.UseStatus,
			lex.Fndtpassword.PasData,
			lex.Fndtpassword.PasType,
		).
		InnerJoin(goqu.T(lex.T.Fndtpassword), goqu.On(goqu.Ex{lex.Fndtpassword.PasUserID: goqu.I(lex.Fndtuser.UseID)})).
		Where(
			goqu.ExOr{
				lex.Fndtuser.UseUsername: user,
				lex.Fndtuser.UseEmail:    user,
			},
			goqu.C(lex.Fndtpassword.PasStatus).Eq(base.StatusActive),
		)

	row, err := dao.h.QueryRowContext(ctx, exe, query)
	if err != nil {
		return nil, dao.h.WrapErr(err)
	}
	err = row.Scan(
		&res.ID,
		&res.Name,
		&res.MiddleName,
		&res.LastName,
		&res.Email,
		&res.Username,
		&res.Status,
		&res.Password.Data,
		&res.Password.Type,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, baseerrors.ErrAuth
		}
		return nil, dao.h.WrapErr(err)
	}
	return res, nil
}

package fnddao

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/jtorz/phoenix-backend/app/services/fnd/fndmodel"
	"github.com/jtorz/phoenix-backend/app/shared/base"
	"github.com/jtorz/phoenix-backend/app/shared/baseerrors"
	"github.com/jtorz/phoenix-backend/app/shared/daohelper"
	"github.com/jtorz/phoenix-backend/app/shared/lex"
)

// DaoUser Data Access structure.
type DaoUser struct {
	Exe base.Executor
	h   daohelper.QueryHelper
}

func (dao *DaoUser) New(ctx context.Context, u *fndmodel.User) error {
	ins := dao.h.NewInsert(lex.T.FndUser).Rows(goqu.Record{
		lex.Fndtuser.UseName:       u.Name,
		lex.Fndtuser.UseMiddleName: u.MiddleName,
		lex.Fndtuser.UseLastName:   u.LastName,
		lex.Fndtuser.UseEmail:      u.Email,
		lex.Fndtuser.UseUsername:   u.Username,
		lex.Fndtuser.UseStatus:     u.Status,
	})
	r, err := dao.h.DoInsertReturning(ctx, dao.Exe, ins, lex.Fndtuser.UseID, lex.Fndtuser.UseUpdatedAt)
	if err != nil {
		return dao.h.WrapErr(err)
	}
	err = r.Scan(&u.ID, &u.UpdatedAt)
	return err
}

// Login retrieves the necessary data to login a user given its email/username.
func (dao *DaoUser) Login(ctx context.Context,
	user string,
) (*fndmodel.User, error) {
	res := &fndmodel.User{Password: &fndmodel.Password{}}
	query := dao.h.NewSelect(lex.T.FndUser).
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
		InnerJoin(goqu.T(lex.T.FndPassword), goqu.On(goqu.Ex{lex.Fndtpassword.PasUserID: goqu.I(lex.Fndtuser.UseID)})).
		Where(
			goqu.ExOr{
				lex.Fndtuser.UseUsername: user,
				lex.Fndtuser.UseEmail:    user,
			},
			goqu.C(lex.Fndtpassword.PasStatus).Eq(base.StatusActive),
		)

	row, err := dao.h.QueryRowContext(ctx, dao.Exe, query)
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

// GetUserByMail returns a user given its email.
func (dao *DaoUser) GetUserByMail(ctx context.Context,
	email string,
) (*fndmodel.User, error) {
	return dao.getUser(ctx, goqu.C(lex.Fndtuser.UseEmail).Eq(email))
}

// GetUserByID retrives the record information using its ID.
func (dao *DaoUser) GetUserByID(ctx context.Context,
	userID string,
) (*fndmodel.User, error) {
	return dao.getUser(ctx, goqu.C(lex.Fndtuser.UseID).Eq(userID))
}

// getUser searchs the user with the given filters.
func (dao *DaoUser) getUser(ctx context.Context,
	filter ...exp.Expression,
) (*fndmodel.User, error) {
	query := dao.h.NewSelect(lex.T.FndUser).
		Select(
			lex.Fndtuser.UseID,
			lex.Fndtuser.UseEmail,
			lex.Fndtuser.UseUsername,
			lex.Fndtuser.UseName,
			lex.Fndtuser.UseMiddleName,
			lex.Fndtuser.UseLastName,
			lex.Fndtuser.UseStatus,
			lex.Fndtuser.UseUpdatedAt,
		).
		Where(filter...)

	row, err := dao.h.QueryRowContext(ctx, dao.Exe, query)
	if err != nil {
		return nil, dao.h.WrapErr(err)
	}
	rec := fndmodel.User{}
	err = row.Scan(
		&rec.ID,
		&rec.Email,
		&rec.Username,
		&rec.Name,
		&rec.MiddleName,
		&rec.LastName,
		&rec.Status,
		&rec.UpdatedAt,
	)
	if err != nil {
		return nil, dao.h.WrapErr(err)
	}
	return &rec, nil
}

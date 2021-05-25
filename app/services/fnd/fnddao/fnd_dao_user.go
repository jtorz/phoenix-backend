package fnddao

import (
	"context"
	"database/sql"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/jtorz/phoenix-backend/app/services/fnd/fndmodel"
	"github.com/jtorz/phoenix-backend/app/shared/base"
	"github.com/jtorz/phoenix-backend/app/shared/baseerrors"

	//lint:ignore ST1001 dot import allowed only in dao packages for lex.
	. "github.com/jtorz/phoenix-backend/app/shared/lex"
)

// DaoUser Data Access structure.
type DaoUser struct{}

func (dao *DaoUser) New(ctx context.Context, tx *sql.Tx,
	u *fndmodel.User,
) error {
	ins := NewInsert(T.FndUser).Rows(goqu.Record{
		FndUser.UseName:       u.Name,
		FndUser.UseMiddleName: u.MiddleName,
		FndUser.UseLastName:   u.LastName,
		FndUser.UseEmail:      u.Email,
		FndUser.UseUsername:   u.Username,
		FndUser.UseStatus:     u.Status,
	})
	r, err := DoInsertReturning(ctx, tx, ins, FndUser.UseID, FndUser.UseUpdatedAt)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	err = r.Scan(&u.ID, &u.UpdatedAt)
	return err
}

// Login retrieves the necessary data to login a user given its email/username.
func (dao *DaoUser) Login(ctx context.Context, exe base.Executor,
	user string,
) (*fndmodel.User, error) {
	res := &fndmodel.User{Password: &fndmodel.Password{}}
	query := NewSelect(
		FndUser.UseID,
		FndUser.UseName,
		FndUser.UseMiddleName,
		FndUser.UseLastName,
		FndUser.UseEmail,
		FndUser.UseUsername,
		FndUser.UseStatus,
		FndPassword.PasData,
		FndPassword.PasType,
	).
		From(T.FndUser).
		InnerJoin(goqu.T(T.FndPassword), FndPasswordFkFndUser()).
		Where(
			goqu.ExOr{
				FndUser.UseUsername: user,
				FndUser.UseEmail:    user,
			},
			goqu.C(FndPassword.PasStatus).Eq(base.StatusActive),
		)

	row, err := QueryRowContext(ctx, exe, query)
	if err != nil {
		DebugErr(ctx, err)
		return nil, err
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
		DebugErr(ctx, err)
		return nil, err
	}
	return res, nil
}

// GetUserByMail returns a user given its email.
func (dao *DaoUser) GetUserByMail(ctx context.Context, exe base.Executor,
	email string,
) (*fndmodel.User, error) {
	return dao.getUser(ctx, exe, goqu.C(FndUser.UseEmail).Eq(email))
}

// GetUserByID retrives the record information using its ID.
func (dao *DaoUser) GetUserByID(ctx context.Context, exe base.Executor,
	userID string,
) (*fndmodel.User, error) {
	return dao.getUser(ctx, exe, goqu.C(FndUser.UseID).Eq(userID))
}

// getUser searchs the user with the given filters.
func (dao *DaoUser) getUser(ctx context.Context, exe base.Executor,
	filter ...exp.Expression,
) (*fndmodel.User, error) {
	query := NewSelect(T.FndUser).
		Select(
			FndUser.UseID,
			FndUser.UseEmail,
			FndUser.UseUsername,
			FndUser.UseName,
			FndUser.UseMiddleName,
			FndUser.UseLastName,
			FndUser.UseStatus,
			FndUser.UseUpdatedAt,
		).
		Where(filter...)

	row, err := QueryRowContext(ctx, exe, query)
	if err != nil {
		DebugErr(ctx, err)
		return nil, err
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
		DebugErr(ctx, err)
		return nil, err
	}
	return &rec, nil
}

// SetStatus changes the logical status of the user.
func (dao *DaoUser) SetStatus(ctx context.Context, tx *sql.Tx,
	u *fndmodel.User,
) error {
	now := time.Now()
	query := NewUpdate(T.FndUser).
		Set(goqu.Record{
			FndUser.UseStatus:    u.ID,
			FndUser.UseUpdatedAt: now,
		}).
		Where(
			goqu.C(FndUser.UseID).Eq(u.ID),
			goqu.C(FndUser.UseUpdatedAt).Eq(u.UpdatedAt),
		)
	res, err := DoUpdate(ctx, tx, query)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	if err = CheckOneRowUpdated(ctx, T.FndUser, res); err != nil {
		return err
	}
	u.UpdatedAt = now
	return nil
}

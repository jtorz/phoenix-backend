package coredal

import (
	"context"
	"database/sql"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/jtorz/phoenix-backend/app/services/core/coremodel"
	"github.com/jtorz/phoenix-backend/app/shared/base"
	"github.com/jtorz/phoenix-backend/app/shared/baseerrors"

	//lint:ignore ST1001 dot import allowed only in dal packages for dalhelper.
	. "github.com/jtorz/phoenix-backend/app/shared/dalhelper"
)

// DalUser Data Access structure.
type DalUser struct{}

func (dal *DalUser) New(ctx context.Context, tx *sql.Tx,
	u *coremodel.User,
) error {
	ins := NewInsert(T.CoreUser).Rows(goqu.Record{
		CoreUser.UseName:       u.Name,
		CoreUser.UseMiddleName: u.MiddleName,
		CoreUser.UseLastName:   u.LastName,
		CoreUser.UseEmail:      u.Email,
		CoreUser.UseUsername:   u.Username,
		CoreUser.UseStatus:     u.Status,
	})
	r, err := DoInsertReturning(ctx, tx, ins, CoreUser.UseID, CoreUser.UseUpdatedAt)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	err = r.Scan(&u.ID, &u.UpdatedAt)
	return err
}

// Login retrieves the necessary data to login a user given its email/username.
func (dal *DalUser) Login(ctx context.Context, exe base.Executor,
	user string,
) (*coremodel.User, error) {
	res := &coremodel.User{Password: &coremodel.Password{}}
	query := NewSelect(
		CoreUser.UseID,
		CoreUser.UseName,
		CoreUser.UseMiddleName,
		CoreUser.UseLastName,
		CoreUser.UseEmail,
		CoreUser.UseUsername,
		CoreUser.UseStatus,
		CorePassword.PasData,
		CorePassword.PasType,
	).
		From(T.CoreUser).
		InnerJoin(goqu.T(T.CorePassword), CorePasswordFkCoreUser()).
		Where(
			goqu.ExOr{
				CoreUser.UseUsername: user,
				CoreUser.UseEmail:    user,
			},
			goqu.C(CorePassword.PasStatus).Eq(base.StatusActive),
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
func (dal *DalUser) GetUserByMail(ctx context.Context, exe base.Executor,
	email string,
) (*coremodel.User, error) {
	return dal.getUser(ctx, exe, goqu.C(CoreUser.UseEmail).Eq(email))
}

// GetUserByID retrives the record information using its ID.
func (dal *DalUser) GetUserByID(ctx context.Context, exe base.Executor,
	userID string,
) (*coremodel.User, error) {
	return dal.getUser(ctx, exe, goqu.C(CoreUser.UseID).Eq(userID))
}

// getUser searchs the user with the given filters.
func (dal *DalUser) getUser(ctx context.Context, exe base.Executor,
	filter ...exp.Expression,
) (*coremodel.User, error) {
	query := NewSelect(
		CoreUser.UseID,
		CoreUser.UseEmail,
		CoreUser.UseUsername,
		CoreUser.UseName,
		CoreUser.UseMiddleName,
		CoreUser.UseLastName,
		CoreUser.UseStatus,
		CoreUser.UseUpdatedAt,
	).
		From(T.CoreUser).
		Where(filter...)

	row, err := QueryRowContext(ctx, exe, query)
	if err != nil {
		DebugErr(ctx, err)
		return nil, err
	}
	rec := coremodel.User{}
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
func (dal *DalUser) SetStatus(ctx context.Context, tx *sql.Tx,
	u *coremodel.User,
) error {
	now := time.Now()
	query := NewUpdate(T.CoreUser).
		Set(goqu.Record{
			CoreUser.UseStatus:    u.ID,
			CoreUser.UseUpdatedAt: now,
		}).
		Where(
			goqu.C(CoreUser.UseID).Eq(u.ID),
			goqu.C(CoreUser.UseUpdatedAt).Eq(u.UpdatedAt),
		)
	res, err := DoUpdate(ctx, tx, query)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	if err = CheckOneRowUpdated(ctx, T.CoreUser, res); err != nil {
		return err
	}
	u.UpdatedAt = now
	return nil
}

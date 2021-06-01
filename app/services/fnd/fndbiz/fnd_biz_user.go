package fndbiz

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"

	"github.com/jtorz/phoenix-backend/app/services/fnd/fnddal"
	"github.com/jtorz/phoenix-backend/app/services/fnd/fndmodel"
	"github.com/jtorz/phoenix-backend/app/shared/base"
	"github.com/jtorz/phoenix-backend/app/shared/baseerrors"
	"github.com/jtorz/phoenix-backend/app/shared/baseservice"
)

// BizUser business component.
type BizUser struct {
	dal *fnddal.DalUser
}

func NewBizUser() BizUser {
	return BizUser{
		dal: &fnddal.DalUser{},
	}
}

// Login retrieves the necessary data to log in a user given its email/username.
func (biz *BizUser) Login(ctx context.Context, exe base.Executor,
	user, pass string,
) (*fndmodel.User, error) {
	u, err := biz.dal.Login(ctx, exe, user)
	if err != nil {
		return nil, err
	}
	if u.Status != base.StatusActive {
		return nil, fmt.Errorf("can't login user in status %s: %w", u.Status, baseerrors.ErrActionNotAllowedStatus)
	}
	if err = u.Password.ComparePassword(pass); err != nil {
		return nil, err
	}
	return u, nil
}

var usernameStyle = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9\._]*$`)

// New creates a new user anf sends the email to activate the account.
func (biz *BizUser) New(ctx context.Context, tx *sql.Tx, senderSvc baseservice.MailSenderSvc,
	u *fndmodel.User,
) error {
	if !usernameStyle.MatchString(u.Username) {
		return fmt.Errorf("%w: username must start with a letter and contain only letter, numbers, dots or underscores", baseerrors.ErrInvalidData)
	}
	u.Status = base.StatusCaptured
	if err := biz.dal.New(ctx, tx, u); err != nil {
		return err
	}
	bizAcc := NewBizAccountAccess()
	if err := bizAcc.NewAccessRestore(ctx, tx, senderSvc, *u, true); err != nil {
		return err
	}
	biz.setRecordActions(ctx, u)
	return nil
}

// GetUserByMail returns a user given its email.
func (biz *BizUser) GetUserByMail(ctx context.Context, exe base.Executor,
	email string,
) (*fndmodel.User, error) {
	u, err := biz.dal.GetUserByMail(ctx, exe, email)
	if err != nil {
		return nil, err
	}
	biz.setRecordActions(ctx, u)
	return u, nil
}

// GetUserByID retrives the user information using its ID.
func (biz *BizUser) GetUserByID(ctx context.Context, exe base.Executor,
	userID string,
) (*fndmodel.User, error) {
	u, err := biz.dal.GetUserByID(ctx, exe, userID)
	if err != nil {
		return nil, err
	}
	biz.setRecordActions(ctx, u)
	return u, nil
}

// RequestRestoreAccount creates an account access to allow the user change the password their own.
func (biz *BizUser) RequestRestoreAccount(ctx context.Context, tx *sql.Tx, senderSvc baseservice.MailSenderSvc,
	email string,
) (*fndmodel.User, error) {
	u, err := biz.GetUserByMail(ctx, tx, email)
	if err != nil {
		if baseerrors.IsErrNotFound(err) {
			return nil, baseerrors.ErrAuth
		}
		return nil, err
	}
	if u.Status == base.StatusInactive {
		return nil, fmt.Errorf("can't activate user account on status inactive: %w", baseerrors.ErrActionNotAllowedStatus)
	}

	bizAcc := NewBizAccountAccess()
	if err = bizAcc.NewAccessRestore(ctx, tx, senderSvc, *u, false); err != nil {
		return nil, err
	}
	return u, nil
}

// RestoreAccount activates the user account sets their password,
// marks the restore account access as active,
// and sends the information to the email service to notify the user.
func (biz *BizUser) RestoreAccount(ctx context.Context, tx *sql.Tx, senderSvc baseservice.MailSenderSvc,
	key, password string,
) (*fndmodel.User, error) {
	bizAcc := NewBizAccountAccess()
	userID, err := bizAcc.UseAccountAccess(ctx, tx, key, fndmodel.AccAccRestoreAccount)
	if err != nil {
		return nil, err
	}
	u, err := biz.GetUserByID(ctx, tx, userID)
	if err != nil {
		return nil, err
	}
	if u.Status == base.StatusInactive {
		return nil, fmt.Errorf("can't activate user account on status inactive: %w", baseerrors.ErrActionNotAllowedStatus)
	}

	u.Status = base.StatusActive
	if err = biz.dal.SetStatus(ctx, tx, u); err != nil {
		return nil, err
	}
	bizPass := NewBizPassword()
	if err = bizPass.ChangeUserPassword(ctx, tx, senderSvc, *u, password); err != nil {
		return nil, err
	}
	return u, nil
}

// setRecordActionsUsers sets the record actiosn to every element in the Users slice.
func (biz *BizUser) setRecordActionsUsers(ctx context.Context,
	recs fndmodel.Users,
) {
	for i := range recs {
		biz.setRecordActions(ctx, &recs[i])
	}
}

// setRecordActions sets the record action sto User record.
func (biz *BizUser) setRecordActions(ctx context.Context,
	rec *fndmodel.User,
) {
	rec.RecordActions = base.NewRecordActionsCommon(rec.Status)
}

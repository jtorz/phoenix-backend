package fndbiz

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"

	"github.com/jtorz/phoenix-backend/app/services/fnd/fnddao"
	"github.com/jtorz/phoenix-backend/app/services/fnd/fndmodel"
	"github.com/jtorz/phoenix-backend/app/shared/base"
	"github.com/jtorz/phoenix-backend/app/shared/baseerrors"
	"github.com/jtorz/phoenix-backend/app/shared/baseservice"
)

// BizUser business component.
type BizUser struct {
	dao *fnddao.DaoUser
}

func NewBizUser() BizUser {
	return BizUser{
		dao: &fnddao.DaoUser{},
	}
}

// Login retrieves the necessary data to log in a user given its email/username.
func (biz *BizUser) Login(ctx context.Context, exe base.Executor,
	user, pass string,
) (*fndmodel.User, error) {
	u, err := biz.dao.Login(ctx, exe, user)
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
	if err := biz.dao.New(ctx, tx, u); err != nil {
		return err
	}
	bizAcc := NewBizAccountAccess()
	if err := bizAcc.NewAccessRestore(ctx, tx, senderSvc, *u, true); err != nil {
		return err
	}
	u.SimpleActions(u.Status)
	return nil
}

// GetUserByMail returns a user given its email.
func (biz *BizUser) GetUserByMail(ctx context.Context, exe base.Executor,
	email string,
) (*fndmodel.User, error) {
	u, err := biz.dao.GetUserByMail(ctx, exe, email)
	if err != nil {
		return nil, err
	}
	u.SimpleActions(u.Status)
	return u, nil
}

// GetUserByID retrives the user information using its ID.
func (biz *BizUser) GetUserByID(ctx context.Context, exe base.Executor,
	userID string,
) (*fndmodel.User, error) {
	u, err := biz.dao.GetUserByID(ctx, exe, userID)
	if err != nil {
		return nil, err
	}
	u.SimpleActions(u.Status)
	return u, nil
}

// RequestRestore creates an account access to allow the user change the password their own.
func (biz *BizUser) RequestRestore(ctx context.Context, tx *sql.Tx, senderSvc baseservice.MailSenderSvc,
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

// Restore activates the user account sets their password,
// marks the restore account access as active,
// and sends the information to the email service to notify the user.
func (biz *BizUser) Restore(ctx context.Context, tx *sql.Tx, senderSvc baseservice.MailSenderSvc,
	key string,
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
	if err = biz.dao.SetStatus(ctx, tx, u); err != nil {
		return nil, err
	}
	// TODO:
	/*

	 */

	data := map[string]interface{}{
		"user": u,
	}
	err = senderSvc.SendMail(ctx, baseservice.MailTemplate{
		SenderUserID: u.ID,
		Type:         baseservice.MailTypeRestoreAccount,
		Data:         data,
		To:           []string{u.Email},
	})
	if err != nil {
		return nil, err
	}
	return u, nil
}

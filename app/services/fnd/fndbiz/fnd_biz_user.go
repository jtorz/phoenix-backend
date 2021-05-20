package fndbiz

import (
	"context"
	"fmt"
	"net/url"

	"github.com/jtorz/phoenix-backend/app/services/fnd/fnddao"
	"github.com/jtorz/phoenix-backend/app/services/fnd/fndmodel"
	"github.com/jtorz/phoenix-backend/app/shared/base"
	"github.com/jtorz/phoenix-backend/app/shared/baseerrors"
	"github.com/jtorz/phoenix-backend/app/shared/baseservice"
)

// BizUser business component.
type BizUser struct {
	Exe base.Executor
	dao DaoUser
}

func NewBizUser(exe base.Executor) BizUser {
	return BizUser{
		Exe: exe,
		dao: &fnddao.DaoUser{Exe: exe},
	}
}

type DaoUser interface {
	Login(_ context.Context, userNameOrEmail string) (*fndmodel.User, error)
	GetUserByMail(ctx context.Context, _ string) (*fndmodel.User, error)
	GetUserByID(ctx context.Context, userID string) (*fndmodel.User, error)
	New(context.Context, *fndmodel.User) error
}

// Login retrieves the necessary data to log in a user given its email/username.
func (biz *BizUser) Login(ctx context.Context,
	user, pass string,
) (*fndmodel.User, error) {

	u, err := biz.dao.Login(ctx, user)
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

// New creates a new user anf sends the email to activate the account.
func (biz *BizUser) New(ctx context.Context, senderSvc baseservice.MailSenderSvc,
	u *fndmodel.User,
) error {
	u.Status = base.StatusCaptured
	if err := biz.dao.New(ctx, u); err != nil {
		return err
	}
	if err := biz.generateAccess(ctx, senderSvc, *u); err != nil {
		return err
	}
	u.SimpleActions(u.Status)
	return nil
}

// GetUserByMail returns a user given its email.
func (biz *BizUser) GetUserByMail(ctx context.Context,
	email string,
) (*fndmodel.User, error) {
	u, err := biz.dao.GetUserByMail(ctx, email)
	if err != nil {
		return nil, err
	}
	u.SimpleActions(u.Status)
	return u, nil
}

// GetUserByID retrives the user information using its ID.
func (biz *BizUser) GetUserByID(ctx context.Context,
	userID string,
) (*fndmodel.User, error) {
	u, err := biz.dao.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	u.SimpleActions(u.Status)
	return u, nil
}

// RequestRestore creates an account access to allow the user change the password their own.
func (biz *BizUser) RequestRestore(ctx context.Context, senderSvc baseservice.MailSenderSvc,
	email string,
) (*fndmodel.User, error) {
	u, err := biz.GetUserByMail(ctx, email)
	if err != nil {
		if baseerrors.IsErrNotFound(err) {
			return nil, baseerrors.ErrAuth
		}
		return nil, err
	}
	if u.Status == base.StatusInactive {
		return nil, fmt.Errorf("can't activate user account on status inactive: %w", baseerrors.ErrActionNotAllowedStatus)
	}

	if err = biz.generateAccess(ctx, senderSvc, *u); err != nil {
		return nil, err
	}
	return u, nil
}

func (biz *BizUser) generateAccess(ctx context.Context, senderSvc baseservice.MailSenderSvc,
	u fndmodel.User,
) (err error) {
	bizAcc := NewBizAccountAccess(biz.Exe)
	ac, err := bizAcc.GetAccessByUserID(ctx, fndmodel.AccAccRestoreAccount, u.ID)

	if err != nil {
		if baseerrors.IsErrNotFound(err) {
			ac.User = u
			return biz.sendEmailAccount(ctx, senderSvc, *ac, true)
		}
		return err
	}

	ac, err = bizAcc.NewAccountAccess(ctx, u)
	if err != nil {
		return
	}
	return biz.sendEmailAccount(ctx, senderSvc, *ac, true)
}

func (biz *BizUser) sendEmailAccount(ctx context.Context, senderSvc baseservice.MailSenderSvc,
	ac fndmodel.AccountAccess, isActivation bool,
) (err error) {
	data := map[string]interface{}{
		"user": ac.User,
		"url":  restoreAccountURLResource(ac.Key, isActivation),
	}

	return senderSvc.SendMail(ctx, baseservice.MailTemplate{
		SenderUserID: ac.User.ID,
		Type:         baseservice.MailTypeRestoreAccount,
		Data:         data,
		To:           []string{ac.User.Email},
	})
}

func restoreAccountURLResource(key string, isActivation bool) string {
	var r = "/restore-account/"
	if isActivation {
		r += "activate/"
	} else {
		r += "restore/"
	}
	return r + url.QueryEscape(key)
}

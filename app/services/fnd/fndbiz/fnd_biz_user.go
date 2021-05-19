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
	dao DaoUser
}

func NewBizUser(exe base.Executor) BizUser {
	return BizUser{dao: &fnddao.DaoUser{Exe: exe}}
}

type DaoUser interface {
	Login(_ context.Context, userNameOrEmail string) (*fndmodel.User, error)
	GetUserByMail(ctx context.Context, _ string) (*fndmodel.User, error)
	GetUserByID(ctx context.Context, userID string) (*fndmodel.User, error)
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

// RequestRestore creates an account access to allow the user change the password their own.
func (biz *BizUser) RequestRestore(ctx context.Context, senderSvc baseservice.MailSenderSvc,
	email string,
) error {
	u := &fndmodel.User{Email: email}
	/* err := biz.GetUserByMail(u)
	if err != nil {
		if fnderrors.IsErrNotFound(err) {
			return fnderrors.ErrAuth
		}
		return err
	}
	if u.Status == basemodel.StatusInactive {
		return fmt.Errorf("can't activate user account on status inactive: %w", fnderrors.ErrActionNotAllowedStatus)
	}
	return biz.generateSendEmailAccount(sender, u) */
	return biz.generateAccess(ctx, senderSvc, u)
}

func (biz *BizUser) generateAccess(ctx context.Context, senderSvc baseservice.MailSenderSvc,
	u *fndmodel.User,
) (err error) {
	/* bizAcc := NewBizAccountAccess(biz.TCtx)
	ac, err := bizAcc.GetAccessByUserID(fndmodel.AccAccRestoreAccount, u.ID)
	if err == nil {
		ac.User = u
		return biz.sendEmailAccount(mail, *ac, true)
	}

	ac, err = fndmodel.NewAccountAccess(u, fndmodel.AccAccRestoreAccount)
	if err != nil {
		return
	}
	err = bizAcc.Insert(ac)
	if err != nil {
		return
	}*/
	ac := &fndmodel.AccountAccess{}
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

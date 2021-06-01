package fndbiz

import (
	"context"
	"database/sql"
	"net/url"

	"github.com/jtorz/phoenix-backend/app/services/fnd/fnddal"
	"github.com/jtorz/phoenix-backend/app/services/fnd/fndmodel"
	"github.com/jtorz/phoenix-backend/app/shared/baseerrors"
	"github.com/jtorz/phoenix-backend/app/shared/baseservice"
	"github.com/jtorz/phoenix-backend/app/shared/ctxinfo"
)

// BizAccountAccess business component.
type BizAccountAccess struct {
	dal *fnddal.DalAccountAccess
}

// NewBizAccountAccess creates business component.
func NewBizAccountAccess() BizAccountAccess {
	return BizAccountAccess{
		dal: &fnddal.DalAccountAccess{},
	}
}

// NewAccessRestore creates the account access for the user to allow them change their password
// and sends the information to the email service.
func (biz *BizAccountAccess) NewAccessRestore(ctx context.Context, tx *sql.Tx, senderSvc baseservice.MailSenderSvc,
	u fndmodel.User, isActivation bool,
) error {
	ac, err := biz.GetOrCreate(ctx, tx, u, fndmodel.AccAccRestoreAccount)
	if err != nil {
		return err
	}

	data := map[string]interface{}{
		"user": ac.User,
		"url":  biz.restoreAccountURLResource(ac.Key, isActivation),
	}
	sender := ac.User.ID
	agent := ctxinfo.GetAgent(ctx)
	if agent != nil && agent.UserID != "" {
		sender = agent.UserID
	}
	return senderSvc.SendMail(ctx, baseservice.MailTemplate{
		SenderUserID: sender,
		Type:         baseservice.MailTypeRestoreAccount,
		Data:         data,
		To:           []string{ac.User.Email},
	})
}

func (biz *BizAccountAccess) restoreAccountURLResource(key string, isActivation bool) string {
	var r = "/restore-account/"
	if isActivation {
		r += "activate/"
	} else {
		r += "restore/"
	}
	return r + url.QueryEscape(key)
}

// GetOrCreate returns the user's active access.
// If there are no active accesses a new acces is created.
func (biz *BizAccountAccess) GetOrCreate(ctx context.Context, exe *sql.Tx,
	u fndmodel.User, accType fndmodel.AccountAccessType,
) (*fndmodel.AccountAccess, error) {
	ac, err := biz.dal.GetAccessByUserID(ctx, exe, u.ID, accType)
	if err == nil {
		return ac, nil
	}
	if !baseerrors.IsErrNotFound(err) {
		return nil, err
	}
	return biz.newAccountAccess(ctx, exe, u, accType)
}

// newAccountAccess creates the new account access for the user.
func (biz *BizAccountAccess) newAccountAccess(ctx context.Context, tx *sql.Tx,
	u fndmodel.User, accType fndmodel.AccountAccessType,
) (*fndmodel.AccountAccess, error) {
	ac, err := fndmodel.NewAccountAccess(u, accType)
	if err != nil {
		return nil, err
	}
	if err = biz.dal.Insert(ctx, tx, ac); err != nil {
		return nil, err
	}
	return ac, nil
}

// UseAccountAccess returns the user ID for the account access and inactivates the access.
func (biz *BizAccountAccess) UseAccountAccess(ctx context.Context, tx *sql.Tx,
	key string, accType fndmodel.AccountAccessType,
) (string, error) {
	return biz.dal.UseAccountAccess(ctx, tx, key, accType)
}

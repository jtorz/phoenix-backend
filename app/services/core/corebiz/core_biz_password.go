package corebiz

import (
	"context"
	"database/sql"

	"github.com/jtorz/phoenix-backend/app/services/core/coredal"
	"github.com/jtorz/phoenix-backend/app/services/core/coremodel"
	"github.com/jtorz/phoenix-backend/app/shared/baseservice"
)

// BizPassword business component.
type BizPassword struct {
	dal *coredal.DalPassword
}

func NewBizPassword() BizPassword {
	return BizPassword{
		dal: &coredal.DalPassword{},
	}
}

// ChangeUserPassword invalidates all the active passwords of the user and creates a new one.
func (biz BizPassword) ChangeUserPassword(ctx context.Context, tx *sql.Tx, senderSvc baseservice.MailSenderSvc,
	u coremodel.User, password string,
) error {
	p, err := coremodel.NewPassword(coremodel.PassTypeScrypt2017, password)
	if err != nil {
		return err
	}
	if err := biz.dal.InvalidateForUser(ctx, tx, u.ID); err != nil {
		return err
	}
	if err := biz.dal.New(ctx, tx, u.ID, p); err != nil {
		return err
	}

	data := map[string]interface{}{
		"user": u,
	}
	err = senderSvc.SendMail(ctx, baseservice.MailTemplate{
		SenderUserID: u.ID,
		Type:         baseservice.MailTypePasswordChanged,
		Data:         data,
		To:           []string{u.Email},
	})
	if err != nil {
		return err
	}
	return nil
}

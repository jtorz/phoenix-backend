package fndbiz

import (
	"context"
	"database/sql"

	"github.com/jtorz/phoenix-backend/app/services/fnd/fnddao"
	"github.com/jtorz/phoenix-backend/app/services/fnd/fndmodel"
	"github.com/jtorz/phoenix-backend/app/shared/baseservice"
)

// BizPassword business component.
type BizPassword struct {
	dao *fnddao.DaoPassword
}

func NewBizPassword() BizPassword {
	return BizPassword{
		dao: &fnddao.DaoPassword{},
	}
}

// ChangeUserPassword invalidates all the active passwords of the user and creates a new one.
func (biz BizPassword) ChangeUserPassword(ctx context.Context, tx *sql.Tx, senderSvc baseservice.MailSenderSvc,
	u fndmodel.User, password string,
) error {
	p, err := fndmodel.NewPassword(fndmodel.PassTypeScrypt2017, password)
	if err != nil {
		return err
	}
	if err := biz.dao.InvalidateForUser(ctx, tx, u.ID); err != nil {
		return err
	}
	if err := biz.dao.New(ctx, tx, u.ID, p); err != nil {
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

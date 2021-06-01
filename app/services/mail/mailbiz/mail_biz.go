package mailbiz

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jtorz/phoenix-backend/app/services/mail/maildal"
	"github.com/jtorz/phoenix-backend/app/shared/base"
	"github.com/jtorz/phoenix-backend/app/shared/baseservice"
)

// BizMail business component.
type BizMail struct {
	Exe    base.Executor
	AppURL string
	dal    DalMail
}

func NewBizMail(exe base.Executor, appURL string) BizMail {
	return BizMail{
		Exe:    exe,
		AppURL: appURL,
		dal:    &maildal.DalMail{Exe: exe},
	}
}

type DalMail interface {
}

// SendMail sends an email using a specific template registered in the database
// according to the MailTemplate.Type.
func (biz BizMail) SendMail(ctx context.Context, mail baseservice.MailTemplate) error {
	bytez, err := json.MarshalIndent(mail, "", "    ")
	if err != nil {
		return err
	}
	fmt.Println(string(bytez))
	fmt.Println("TODO: BizMail.SendMail", strings.Repeat("*", 100))
	return nil
}

// SendMailGeneral sends a simple email with the MailGeneral data .
func (biz BizMail) SendMailGeneral(ctx context.Context, mail baseservice.MailGeneral) error {
	bytez, err := json.MarshalIndent(mail, "", "    ")
	if err != nil {
		return err
	}
	fmt.Println(string(bytez))
	fmt.Println("TODO: BizMail.SendMailGeneral", strings.Repeat("*", 100))
	return nil
}

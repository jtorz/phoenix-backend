package mail

import (
	"context"
	"database/sql"

	"github.com/jtorz/phoenix-backend/app/services/mail/mailbiz"
	"github.com/jtorz/phoenix-backend/app/shared/baseservice"
)

// Service is used to send emails, either by using a defined template or using simple string data.
//
// implements the github.com/jtorz/phoenix-backend/app/shared/baseservice.MailSender interface.
type Service struct {
	DB     *sql.DB
	AppURL string
}

func NewService(db *sql.DB, appURL string) *Service {
	return &Service{DB: db, AppURL: appURL}
}

// SendMail sends an email using a specific template registered in the database
// according to the MailTemplate.Type.
func (senderSvc Service) SendMail(ctx context.Context, mail baseservice.MailTemplate) error {
	return mailbiz.NewBizMail(senderSvc.DB, senderSvc.AppURL).SendMail(ctx, mail)
}

// SendMailGeneral sends a simple email with the MailGeneral data.
func (senderSvc Service) SendMailGeneral(ctx context.Context, mail baseservice.MailGeneral) error {
	return mailbiz.NewBizMail(senderSvc.DB, senderSvc.AppURL).SendMailGeneral(ctx, mail)
}

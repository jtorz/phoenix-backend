package mail

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jtorz/phoenix-backend/app/shared/baseservice"
)

// Service is used to send emails, either by using a defined template or using simple string data.
//
// implements the github.com/jtorz/phoenix-backend/app/shared/baseservice.MailSender interface.
type Service struct {
	DB     *sql.DB
	Config ServiceConfig
	AppURL string
}

func NewService(db *sql.DB, config ServiceConfig, appURL string) *Service {
	return &Service{DB: db, Config: config, AppURL: appURL}
}

// SendMail sends an email using a specific template registered in the database
// according to the MailTemplate.Type.
func (sender Service) SendMail(context.Context, baseservice.MailTemplate) error {
	return errors.New("TODO: not implemented")
}

// SendMailGeneral sends a simple email with the MailGeneral data .
func (sender Service) SendMailGeneral(context.Context, baseservice.MailGeneral) error {
	return errors.New("TODO: not implemented")
}

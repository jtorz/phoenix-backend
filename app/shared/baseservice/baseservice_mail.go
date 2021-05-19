// Package baseservice defines general services that can be used accross the system.
package baseservice

import "context"

// MailSenderSvc is used to send emails, either by using a defined template or using simple string data.
type MailSenderSvc interface {
	// SendMail sends an email using a specific template
	// according to the MailTemplate.Type.
	SendMail(context.Context, MailTemplate) error

	// SendMailGeneral sends a simple email with the MailGeneral data .
	SendMailGeneral(context.Context, MailGeneral) error
}

// MailType catalog of emails send by the system.
type MailType string

const (
	// MailTypeRestoreAccount mail that sends the account access to change the user password.
	MailTypeRestoreAccount MailType = "REST_ACC"
	// MailTypePasswordChanged mail that notifies that the user password changed.
	MailTypePasswordChanged MailType = "PASS_CHG"
	// MailTypeMailChanged mail that notifies that the user email changed.
	MailTypeMailChanged MailType = "MAIL_CHG"
)

// MailFile represents the file that can be added to the email message.
// You can add attachment from file in path, from base64 string or from []byte.
// You can define if attachment is inline or not.
// Only one, Data, B64Data or FilePath is supported. If multiple are set, then
// the first in that order is used.
type MailFile struct {
	// FilePath is the path of the file to attach.
	FilePath string
	// Name is the name of file in attachment. Required for Data and B64Data. Optional for FilePath.
	Name string
	// MimeType of attachment. If empty then is obtained from Name (if not empty) or FilePath. If cannot obtained, application/octet-stream is set.
	MimeType string
	// B64Data is the base64 string to attach.
	B64Data string
	// Data is the []byte of file to attach.
	Data []byte
	// Inline defines if attachment is inline or not.
	Inline bool
}

// MailTemplate holds the data for emails generated from a template.
type MailTemplate struct {
	SenderUserID string
	Type         MailType
	Data         map[string]interface{}
	To           []string
	Cc           []string
	Bcc          []string
	Attachment   []MailFile
}

// MimeText is the mime type text.
//
// From mozilla.org: The type represents the general category into which the data type falls, such as video or text.
type MimeText string

const (
	// MimeTextPlain mime type text/plain
	//
	// From mozilla.org: This is the default for textual files. Even if it really means "unknown textual file," browsers assume they can display it.
	MimeTextPlain MimeText = "text/plain"

	// MimeTextHTML mime type text/html
	//
	// From mozilla.org: All HTML content should be served with this type
	MimeTextHTML MimeText = "text/html"
)

// MailGeneral holds the data for simple email with html content.
type MailGeneral struct {
	SenderUserID string
	Subject      string
	Content      string
	From         string
	To           []string
	Cc           []string
	Bcc          []string
	Attachment   []MailFile
	MimeText     MimeText
}

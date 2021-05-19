package main

import (
	"log"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
)

const htmlBody = `<html>
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
		<title>Hello Gophers!</title>
	</head>
	<body>
		<p>This is the <b>Go gopher</b>.</p>
		<p><img src="cid:Gopher.png" alt="Go gopher" /></p>
		<p>Image created by Renee French</p>
	</body>
</html>`

func main() {
	client := mail.NewSMTPClient()

	//SMTP Client
	client.Host = "smtp.gmail.com"
	client.Port = 465
	client.Username = "example@gmail.comm"
	client.Password = "asdasd"
	client.Encryption = mail.EncryptionSSLTLS
	client.ConnectTimeout = 10 * time.Second
	client.SendTimeout = 10 * time.Second

	//KeepAlive is not settted because by default is false

	//Connect to client
	smtpClient, err := client.Connect()

	if err != nil {
		log.Fatal(err)
		return
	}

	err = sendEmail(htmlBody, "email.com", smtpClient)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func sendEmail(htmlBody string, to string, smtpClient *mail.SMTPClient) error {
	//Create the email message
	email := mail.NewMSG()

	email.SetFrom("From Example <from.email@example.com>").
		AddTo(to).
		Attach(&mail.File{
			Name:     "test.json",
			MimeType: "application/json",
			Data:     []byte(`{"hello":World}`),
		}).
		SetSubject("New Go Email")

	//Get from each mail
	email.GetFrom()
	email.SetBody(mail.TextHTML, htmlBody)

	//Send with high priority
	email.SetPriority(mail.PriorityHigh)

	// always check error after send
	if email.Error != nil {
		return email.Error
	}

	//Pass the client to the email message to send it
	return email.Send(smtpClient)
}

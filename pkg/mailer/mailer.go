package mailer

import gomail "gopkg.in/gomail.v2"

type AppMailer struct {
	Gmail       string
	dialer      *gomail.Dialer
	AppPassword string
}

func NewAppMailer(gmail, appPassword string) *AppMailer {
	return &AppMailer{
		Gmail: gmail,
		dialer: gomail.NewDialer(
			"smtp.gmail.com",
			587,
			gmail,
			appPassword,
		),
		AppPassword: appPassword,
	}
}

func (am *AppMailer) SendGmail(receiver, subject, body string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", am.Gmail)
	msg.SetHeader("To", receiver)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)
	return am.dialer.DialAndSend(msg)
}

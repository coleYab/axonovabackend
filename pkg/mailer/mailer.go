package mailer

import (
	"fmt"
	"log"

	"github.com/resend/resend-go/v3"
)

type AppMailer struct {
	sender string
	client *resend.Client
}

func NewAppMailer(gmail, apiKey string) *AppMailer {
	return &AppMailer{
		sender: gmail,
		client: resend.NewClient(apiKey),
	}
}

func (am *AppMailer) SendGmail(receiver, subject, body string) error {
	params := &resend.SendEmailRequest{From: am.sender, To: []string{receiver}, Subject: subject, Html: body}

	sent, err := am.client.Emails.Send(params)
	if err != nil {
		return fmt.Errorf("failed to send an email %s", err.Error())
	}

	log.Println("Email was sent successfully, id=", sent.Id)
	return nil
}

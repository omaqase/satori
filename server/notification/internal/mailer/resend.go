package mailer

import (
	"github.com/omaqase/satori/notification/internal/config"
	protobuf "github.com/omaqase/satori/notification/protobuf/gen"
	"github.com/resend/resend-go/v2"
)

type Resend struct {
	Client *resend.Client
}

func NewResendClient(config config.Config) *Resend {
	return &Resend{
		Client: resend.NewClient(config.Resend.ApiKey),
	}
}

func (r *Resend) Send(in *protobuf.SendNotificationRequest) error {
	params := &resend.SendEmailRequest{
		From:    "satori@kiteo.app",
		To:      []string{in.Receiver},
		Subject: "Satori Notification",
		Html:    in.Content,
	}

	_, err := r.Client.Emails.Send(params)

	return err
}

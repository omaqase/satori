package mailer

import (
	"context"
	"github.com/omaqase/satori/notification/pkg/validators"
	protobuf "github.com/omaqase/satori/notification/protobuf/gen"
	"google.golang.org/grpc"
)

type Mailer struct {
	Resend *Resend
}

func NewMailer(resend *Resend) *Mailer {
	return &Mailer{
		Resend: resend,
	}
}

func (m *Mailer) SendNotification(ctx context.Context, in *protobuf.SendNotificationRequest, opts ...grpc.CallOption) (*protobuf.SendNotificationResponse, error) {
	out := &protobuf.SendNotificationResponse{}

	if err := validators.ValidateNotificationRequest(in); err != nil {
		return out, err
	}

	if err := m.Resend.Send(in); err != nil {
		return out, err
	}

	return out, nil
}

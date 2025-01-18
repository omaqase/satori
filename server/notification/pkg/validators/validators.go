package validators

import (
	"errors"
	protobuf "github.com/omaqase/satori/notification/protobuf/gen"
	"net/mail"
)

var ErrInvalidBody = errors.New("invalid body: provided content are not valid")
var ErrInvalidMeta = errors.New("invalid metadata: provided meta content are not valid")

func ValidateNotificationRequest(in *protobuf.SendNotificationRequest) error {
	if len(in.Content) == 0 {
		return ErrInvalidBody
	}

	if !ValidateEmail(in.Content) {
		return ErrInvalidMeta
	}

	return nil
}

func ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err != nil
}

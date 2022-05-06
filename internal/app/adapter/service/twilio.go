package service

import (
	"ddd-boilerplate/internal/app/domain/vo"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Twilio struct{}

func (t Twilio) CreateSession(serviceID string, paticipants ...vo.Paticipant) (string, error) {
	sessionID := uuid.NewString()
	err := httpCall(serviceID, paticipants)
	if err != nil {
		return "", errors.Wrap(err, "http call error on twilio.CreateSession")
	}
	return sessionID, nil
}

func httpCall(args ...interface{}) error {
	return nil
}

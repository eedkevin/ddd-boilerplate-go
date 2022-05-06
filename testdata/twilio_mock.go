package testdata

import (
	"ddd-boilerplate/internal/app/domain/vo"

	"github.com/google/uuid"
)

type MTwilio struct{}

func (t MTwilio) CreateSession(serviceID string, paticipants ...vo.Paticipant) (string, error) {
	return uuid.NewString(), nil
}

package usecase

import (
	"ddd-boilerplate/internal/app/application/service"
	"ddd-boilerplate/internal/app/domain"
	"ddd-boilerplate/internal/app/domain/repository"
	"ddd-boilerplate/internal/app/domain/vo"
)

// CreateSession creates a session
func CreateSession(telecom service.ITelecom, servicePool repository.IServicePool, paticipants ...vo.Paticipant) (domain.Session, error) {
	service := servicePool.GetOne()
	sessionID, err := telecom.CreateSession(service.ID, paticipants...)
	if err != nil {
		return domain.Session{}, err
	}

	return domain.Session{
		ID:          sessionID,
		Paticipants: paticipants,
		Expiry:      "2021-12-30T20:00:00Z",
	}, nil
}

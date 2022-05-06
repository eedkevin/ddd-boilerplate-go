package service

import "ddd-boilerplate/internal/app/domain/vo"

type ITelecom interface {
	CreateSession(serviceID string, paticipants ...vo.Paticipant) (string, error)
}

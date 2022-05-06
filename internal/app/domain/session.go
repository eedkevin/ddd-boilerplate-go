package domain

import "ddd-boilerplate/internal/app/domain/vo"

type Session struct {
	ID          string
	Paticipants []vo.Paticipant
	Expiry      string
}

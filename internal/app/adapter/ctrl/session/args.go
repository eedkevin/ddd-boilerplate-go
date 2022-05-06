package session

import (
	"ddd-boilerplate/internal/app/domain/vo"
)

// CreateSessionArgs defines the args for CreateSession
type CreateSessionArgs struct {
	Paticipants []vo.Paticipant `json:"paticipants" binding:"required" validate:"required"`
}

// CloseSessionArgs defines the args for CloseSession
type CloseSessionArgs struct {
	SessionID string `json:"session_id" binding:"required" validate:"required"`
}

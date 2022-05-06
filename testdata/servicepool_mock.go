package testdata

import (
	"ddd-boilerplate/internal/app/domain"

	"github.com/google/uuid"
)

type MTwilioServicePool struct{}

func (t MTwilioServicePool) GetOne() domain.Service {
	return domain.Service{
		ID: uuid.NewString(),
	}
}

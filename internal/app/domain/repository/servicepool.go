package repository

import "ddd-boilerplate/internal/app/domain"

type IServicePool interface {
	GetOne() domain.Service
}

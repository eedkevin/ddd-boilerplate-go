package repository

import (
	"ddd-boilerplate/internal/app/domain"
	"ddd-boilerplate/internal/app/infrastructure/redis"
)

//TwilioServicePool is the twilio service pool
type TwilioServicePool struct {
	Redis        *redis.Client
	RedisCluster *redis.ClusterClient
}

//GetOne returns a service
func (t TwilioServicePool) GetOne() domain.Service {
	return domain.Service{
		ID: redis.Pop(),
	}
}

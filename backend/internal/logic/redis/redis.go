package redis

import "backend/internal/service"

type sRedis struct {
}

func init() {
	service.RegisterRedis(&sRedis{})
}

package middleware

import "backend/internal/service"

type sMiddleware struct {
}

func init() {
	service.RegisterMiddleware(&sMiddleware{})
}

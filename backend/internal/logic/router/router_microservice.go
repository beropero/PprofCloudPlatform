package router

import (
	"backend/internal/controller/microservice"
	"backend/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (s *sRouter) RegisterMicroservice(g *ghttp.RouterGroup) {
	g.Group("/microservice", func(g *ghttp.RouterGroup) {
		g.Middleware(
			service.Middleware().JwtInterceptor,
		)
		g.Bind(
			microservice.NewV1(),
		)
	})
}

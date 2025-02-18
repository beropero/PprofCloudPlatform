package router

import (
	"backend/internal/controller/profile"
	"backend/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (s *sRouter) RegisterProfile(g *ghttp.RouterGroup) {
	g.Group("/profile", func(g *ghttp.RouterGroup) {
		g.Middleware(
			service.Middleware().JwtInterceptor,
		)
		g.Bind(
			profile.NewV1(),
		)
	})
}

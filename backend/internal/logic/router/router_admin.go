package router

import (
	"backend/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (s *sRouter) RegisterAdmin(g *ghttp.RouterGroup) {
	g.Group("/admin", func(g *ghttp.RouterGroup) {
		g.Middleware(
			service.Middleware().JwtInterceptor,
			service.Middleware().RoleInterceptor([]string{"admin"}),
		)
		g.Bind()
	})
}

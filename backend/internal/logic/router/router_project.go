package router

import (
	"backend/internal/controller/project"
	"backend/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (s *sRouter) RegisterProject(g *ghttp.RouterGroup) {
	g.Group("/project", func(g *ghttp.RouterGroup) {
		g.Middleware(
			service.Middleware().JwtInterceptor,
		)
		g.Bind(
			project.NewV1(),
		)
	})
}

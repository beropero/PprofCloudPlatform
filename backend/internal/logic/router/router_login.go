package router

import (
	"backend/internal/controller/login"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (s *sRouter) RegisterLogin(g *ghttp.RouterGroup) {
	g.Group("/login", func(g *ghttp.RouterGroup) {
		g.Bind(
			login.NewV1(),
		)
	})
}

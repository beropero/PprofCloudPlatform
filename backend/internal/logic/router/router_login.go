package router

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

func (s *sRouter) RegisterLogin(g *ghttp.RouterGroup) {
	g.Group("/login", func(g *ghttp.RouterGroup) {
		g.Middleware()
		g.Bind()
	})
}

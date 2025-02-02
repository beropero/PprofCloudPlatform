package router

import (
	"backend/internal/controller/upload"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (s *sRouter) RegisterUpload(g *ghttp.RouterGroup) {
	g.Group("/upload", func(g *ghttp.RouterGroup) {
		g.Middleware()
		g.Bind(
			upload.NewV1(),
		)
	})
}

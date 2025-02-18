package router

import (
	"backend/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

type sRouter struct {
}

func init() {
	service.RegisterRouter(&sRouter{})
}

func (r *sRouter) Register(s *ghttp.Server) {
	var (
		router = service.Router()
	)

	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(
			service.Middleware().Ctx,
			ghttp.MiddlewareHandlerResponse,
			service.Middleware().CORS,
		)
		router.RegisterLogin(group)
		router.RegisterAdmin(group)
		router.RegisterUpload(group)
		router.RegisterMicroservice(group)
		router.RegisterProject(group)
		router.RegisterProfile(group)
	})
}

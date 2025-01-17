package middleware

import (
	"backend/internal/model"
	"backend/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (s *sMiddleware) Ctx(r *ghttp.Request) {
	customCtx := &model.Context{}
	service.BizCtx().Init(r, customCtx)
	r.Middleware.Next()
}

func (s *sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

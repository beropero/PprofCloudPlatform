package middleware

import (
	"backend/internal/dao"
	"backend/internal/model"
	"backend/internal/model/entity"
	"backend/internal/service"
	"fmt"
	"net/http"
	"sort"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// JWT校验拦截器 (需要在ctx中间件之后绑定)
func (s *sMiddleware) JwtInterceptor(r *ghttp.Request) {
	var (
		ctx         = r.Context()
		tokenString = r.Header.Get("Authorization")
	)

	// 解析JWT令牌
	decode, err := service.Jwt().DecodeJwt(ctx, model.DecodeJwtInput{Token: tokenString})
	if err != nil {
		r.Response.WriteStatus(http.StatusForbidden)
		r.Exit()
	}
	userid := decode.Userid
	// userrole := decode.UserRole
	rediskey := fmt.Sprintf("token:user:%s", userid)
	// Redis中是否存在JWT
	search, _ := g.Redis().Do(ctx, "GET", rediskey)
	if search.String() != tokenString {
		r.Response.WriteStatus(http.StatusForbidden)
		r.Exit()
	}

	// 拉取用户信息
	var user = entity.Users{}
	err = dao.Users.Ctx(ctx).Where("user_id", userid).Scan(&user)
	if err != nil || user.UserId == 0 {
		r.Response.WriteStatus(http.StatusForbidden)
		r.Exit()
	}
	service.BizCtx().SetUser(ctx, &model.ContextUser{
		UserId:         user.UserId,
		UserPermission: user.Permission,
		UserEmail:      user.Email,
	})
	// 账号未激活
	if user.Permission == "inactive" {
		r.Response.WriteStatus(http.StatusForbidden)
		r.Exit()
	}
	r.Middleware.Next()
}

// 权限拦截器
func (s *sMiddleware) RoleInterceptor(role []string) func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		var (
			ctx  = r.Context()
			user = service.BizCtx().Get(ctx).User
		)
		sort.Strings(role)
		index := sort.SearchStrings(role, user.UserPermission)
		// 权限不匹配
		if index >= len(role) || role[index] != user.UserPermission {
			r.Response.WriteStatus(http.StatusForbidden)
			r.Exit()
		}
		r.Middleware.Next()
	}
}

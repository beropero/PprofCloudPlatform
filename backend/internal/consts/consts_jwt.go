package consts

import "github.com/gogf/gf/v2/frame/g"

var (
	JwtCfg = struct {
		Secret  string
		Expired int
	}{
		Secret:  g.Cfg().MustGet(ctx, "jwt.secret").String(),
		Expired: g.Cfg().MustGet(ctx, "jwt.expired").Int() * 3600 * 24,
	}
)

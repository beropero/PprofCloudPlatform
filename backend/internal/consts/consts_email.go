package consts

import (
	"github.com/gogf/gf/v2/frame/g"
)

var (
	EmailCfg = struct {
		Host     string
		Port     int
		Username string
		Password string
		Expired  int
	}{
		Host:     g.Cfg().MustGet(ctx, "email.host").String(),
		Port:     g.Cfg().MustGet(ctx, "email.port").Int(),
		Username: g.Cfg().MustGet(ctx, "email.username").String(),
		Password: g.Cfg().MustGet(ctx, "email.password").String(),
		Expired:  g.Cfg().MustGet(ctx, "email.expired").Int(),
	}
)

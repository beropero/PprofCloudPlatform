package v1

import "github.com/gogf/gf/v2/frame/g"

type LoginReq struct {
	g.Meta   `path:"/login" method:"post" tags:"Login" summary:"登录"`
	Email    string `json:"email"    v:"required,email"`
	Password string `json:"password" v:"required,length:6,20"`
}
type LoginRes struct {
	Token string `json:"token"`
}

type RegisterReq struct {
	g.Meta   `path:"/register" method:"post" tags:"Login" summary:"注册"`
	Email    string `json:"email"    v:"required,email"`
	Password string `json:"password" v:"required,length:6,20"`
	Code     string `json:"code"`
}

type RegisterRes struct {
	Token string `json:"token"`
}

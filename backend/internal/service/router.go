// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IRouter interface {
		Register(s *ghttp.Server)
		RegisterAdmin(g *ghttp.RouterGroup)
		RegisterLogin(g *ghttp.RouterGroup)
		RegisterMicroservice(g *ghttp.RouterGroup)
		RegisterProfile(g *ghttp.RouterGroup)
		RegisterProject(g *ghttp.RouterGroup)
		RegisterUpload(g *ghttp.RouterGroup)
	}
)

var (
	localRouter IRouter
)

func Router() IRouter {
	if localRouter == nil {
		panic("implement not found for interface IRouter, forgot register?")
	}
	return localRouter
}

func RegisterRouter(i IRouter) {
	localRouter = i
}

package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UploadTotalReq struct {
	g.Meta    `path:"/uploadtotal" method:"post" tags:"Upload" mime:"multipart/form-data" summary:"上传所有性能数据"`
	Cpu       *ghttp.UploadFile `form:"cpu" json:"cpu"`
	Memory    *ghttp.UploadFile `form:"memory" json:"memory"`
	Mutex     *ghttp.UploadFile `form:"mutex" json:"mutex"`
	Block     *ghttp.UploadFile `form:"block" json:"block"`
	Goroutine *ghttp.UploadFile `form:"goroutine" json:"goroutine"`
}

type UploadTotalRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

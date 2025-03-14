package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UploadReq struct {
	g.Meta  `path:"/uploadfile" method:"post" tags:"Upload" mime:"multipart/form-data" summary:"上传所有性能数据"`
	Type    string            `form:"type" json:"type" v:"required#类型不能为空"`
	File    *ghttp.UploadFile `form:"file"       json:"file" v:"required#类型不能为空"`
	Comment string            `form:"comment"   json:"comment" `
}

type UploadRes struct {
	Msg string `json:"msg"`
}

package model

import "github.com/gogf/gf/v2/net/ghttp"

type UploadInput struct {
	Type    string            `form:"type"      json:"type"`
	File    *ghttp.UploadFile `form:"file"       json:"file"`
	Token   string            `form:"token"     json:"token"`
	Comment string            `form:"comment"   json:"comment"`
}

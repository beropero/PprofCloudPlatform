package model

import "github.com/gogf/gf/v2/net/ghttp"

type UploadInput struct {
	Types     []string          `form:"types"      json:"types"`
	Cpu       *ghttp.UploadFile `form:"cpu"       json:"cpu"`
	Memory    *ghttp.UploadFile `form:"memory"    json:"memory"`
	Mutex     *ghttp.UploadFile `form:"mutex"     json:"mutex"`
	Block     *ghttp.UploadFile `form:"block"     json:"block"`
	Goroutine *ghttp.UploadFile `form:"goroutine" json:"goroutine"`
}

package model

import (
	"regexp"

	"github.com/gogf/gf/v2/net/ghttp"
)

type AnalysisInput struct {
	Types     string            `form:"type"      json:"types"`
	Cpu       *ghttp.UploadFile `form:"cpu"       json:"cpu"`
	Memory    *ghttp.UploadFile `form:"memory"    json:"memory"`
	Mutex     *ghttp.UploadFile `form:"mutex"     json:"mutex"`
	Block     *ghttp.UploadFile `form:"block"     json:"block"`
	Goroutine *ghttp.UploadFile `form:"goroutine" json:"goroutine"`
}

// 可配置的过滤规则
type FilterConfig struct {
	ExcludeFunctionPrefixes []string       // 需要过滤的函数名前缀（如 ["runtime.", "reflect."]）
	ExcludeFunctionRegex    *regexp.Regexp // 正则过滤函数名
	KeepRuntime             bool           // 是否保留部分 runtime 关键函数（如 runtime.main）
	MinSampleValue          int64          // 过滤低于此值的样本（如 CPU 时间 < 1ms）
}

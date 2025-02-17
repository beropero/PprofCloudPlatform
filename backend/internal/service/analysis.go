// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"backend/internal/model"
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/google/pprof/profile"
)

type (
	IAnalysis interface {
		// Analysis 解析所有性能数据
		Analysis(ctx context.Context, in model.AnalysisInput) (err error)
		AnalysisCpu(ctx context.Context, file *ghttp.UploadFile) error
		AnalysisMemory(ctx context.Context, file *ghttp.UploadFile) (err error)
		AnalysisMutex(ctx context.Context, file *ghttp.UploadFile) (err error)
		AnalysisBlock(ctx context.Context, file *ghttp.UploadFile) (err error)
		AnalysisGoroutine(ctx context.Context, file *ghttp.UploadFile) (err error)
		// 高性能样本过滤与合并
		FilterSamplesConcurrent(samples []*profile.Sample, config *model.FilterConfig) []*profile.Sample
		ProcessProfile(ctx context.Context, p *profile.Profile) (err error)
		// InsertProfile 插入 profile 主表数据
		InsertProfile(ctx context.Context, p *profile.Profile) (profileId int64, err error)
		// InsertSampleTypes 插入 SampleType
		InsertSampleTypes(ctx context.Context, types []*profile.ValueType, profileId int64) (err error)
		// InsertMappings 插入 mapping 信息
		InsertMappings(ctx context.Context, mappings []*profile.Mapping, profileId int64) (mappingIDMap map[uint64]int64, err error)
		// insertFunctions 插入 Function 并返回 ID 映射表
		InsertFunctions(ctx context.Context, profileID int64, funcs []*profile.Function) (map[uint64]int64, error)
		// insertLocations 插入 Location 并返回 ID 映射表
		InsertLocations(ctx context.Context, profileID int64, locs []*profile.Location, mappingIDMap map[uint64]int64) (map[uint64]int64, error)
		// insertLineEntries 插入 LineEntry
		InsertLineEntries(ctx context.Context, locs []*profile.Location, locationIDMap map[uint64]int64, funcIDMap map[uint64]int64) error
		// insertSamples 插入 Sample 及关联数据
		InsertSamples(ctx context.Context, profileID int64, samples []*profile.Sample, locationIDMap map[uint64]int64) error
		// insertComments 插入注释
		InsertComments(ctx context.Context, profileID int64, comments []string) error
	}
)

var (
	localAnalysis IAnalysis
)

func Analysis() IAnalysis {
	if localAnalysis == nil {
		panic("implement not found for interface IAnalysis, forgot register?")
	}
	return localAnalysis
}

func RegisterAnalysis(i IAnalysis) {
	localAnalysis = i
}

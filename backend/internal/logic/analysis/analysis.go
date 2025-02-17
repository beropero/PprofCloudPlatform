package analysis

// import (
// 	"backend/internal/model"
// 	"backend/internal/service"
// 	"bytes"
// 	"context"
// 	"fmt"
// 	"io"
// 	"strings"

// 	"github.com/gogf/gf/v2/frame/g"
// 	"github.com/gogf/gf/v2/net/ghttp"
// 	"github.com/google/pprof/profile"
// )

// type sAnalysis struct {
// }

// func init() {
// 	service.RegisterAnalysis(&sAnalysis{})
// }

// // 预校验所有类型合法性,并去除重复
// func validateTypes(types []string) (dedupedTypes []string, err error) {

// 	uniqueTypes := make(map[string]struct{})
// 	dedupedTypes = make([]string, 0, len(types))
// 	for _, t := range types {
// 		if _, exists := uniqueTypes[t]; exists {
// 			continue
// 		}
// 		uniqueTypes[t] = struct{}{}
// 		dedupedTypes = append(dedupedTypes, t)
// 	}
// 	return
// }

// // Analysis 解析所有性能数据
// func (s *sAnalysis) Analysis(ctx context.Context, in model.AnalysisInput) (err error) {
// 	// 解析请求参数
// 	types, err := validateTypes(strings.Split(in.Types, ","))
// 	if err != nil {
// 		return err
// 	}
// 	// 捕获指定类型
// 	for _, t := range types {
// 		switch t {
// 		case "block":
// 			if in.Block != nil {
// 				// 处理 block 类型
// 				p, err := GetPprofObjFormFile(in.Block)
// 				if err != nil {
// 					return fmt.Errorf("解析 Profile 数据失败: %w", err)
// 				}
// 				service.Analysis().ProcessProfile(ctx, p)
// 			}
// 		case "cpu":
// 			if in.Cpu != nil {
// 				// 处理 cpu 类型
// 				p, err := GetPprofObjFormFile(in.Cpu)
// 				if err != nil {
// 					return fmt.Errorf("解析 Profile 数据失败: %w", err)
// 				}
// 				service.Analysis().ProcessProfile(ctx, p)
// 			}
// 		case "goroutine":
// 			if in.Goroutine != nil {
// 				// 处理 goroutine 类型
// 				p, err := GetPprofObjFormFile(in.Goroutine)
// 				if err != nil {
// 					return fmt.Errorf("解析 Profile 数据失败: %w", err)
// 				}
// 				service.Analysis().ProcessProfile(ctx, p)
// 			}

// 		case "memory":
// 			if in.Memory != nil {
// 				// 处理 memory 类型
// 				p, err := GetPprofObjFormFile(in.Memory)
// 				if err != nil {
// 					return fmt.Errorf("解析 Profile 数据失败: %w", err)
// 				}
// 				g.Dump(p)
// 				fmt.Println(p.String())
// 				//service.Analysis().ProcessProfile(ctx, p)
// 			}
// 		case "mutex":
// 			if in.Mutex != nil {
// 				// 处理 mutex 类型
// 				p, err := GetPprofObjFormFile(in.Mutex)
// 				if err != nil {
// 					return fmt.Errorf("解析 Profile 数据失败: %w", err)
// 				}
// 				service.Analysis().ProcessProfile(ctx, p)
// 			}
// 		default:
// 			return fmt.Errorf("unknown profile type: %q", t)
// 		}

// 	}

// 	return
// }

// // 从文件中获取 Profile 对象
// func GetPprofObjFormFile(file *ghttp.UploadFile) (prof *profile.Profile, err error) {
// 	// 解析 Profile 数据
// 	fileData, _ := file.Open()
// 	// 读取文件内容
// 	data, err := io.ReadAll(fileData)
// 	if err != nil {
// 		return nil, fmt.Errorf("读取文件内容失败: %w", err)
// 	}
// 	prof, err = profile.Parse(bytes.NewReader(data))
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to parse profile: %v", err)
// 	}

// 	return
// }

// func (s *sAnalysis) AnalysisCpu(ctx context.Context, file *ghttp.UploadFile) error {
// 	prof, err := GetPprofObjFormFile(file)
// 	if err != nil {
// 		return fmt.Errorf("解析 Profile 数据失败: %w", err)
// 	}
// 	service.Analysis().FilterSamplesConcurrent(prof.Sample, nil)
// 	return nil
// }

// func (s *sAnalysis) AnalysisMemory(ctx context.Context, file *ghttp.UploadFile) (err error) {
// 	prof, err := GetPprofObjFormFile(file)
// 	if err != nil {
// 		return fmt.Errorf("解析 Profile 数据失败: %w", err)
// 	}
// 	service.Analysis().FilterSamplesConcurrent(prof.Sample, nil)
// 	return
// }

// func (s *sAnalysis) AnalysisMutex(ctx context.Context, file *ghttp.UploadFile) (err error) {
// 	return
// }

// func (s *sAnalysis) AnalysisBlock(ctx context.Context, file *ghttp.UploadFile) (err error) {
// 	return
// }

// func (s *sAnalysis) AnalysisGoroutine(ctx context.Context, file *ghttp.UploadFile) (err error) {
// 	return
// }

// // func (s *sAnalysis)extractSamples(p *profile.Profile) []Sample {
// // 	var samples []Sample
// // 	for _, s := range p.Sample {
// // 		stacktrace := make([]string, len(s.Location))
// // 		for i, loc := range s.Location {
// // 			line := loc.Line[0] // 取最顶层的调用位置
// // 			funcName := line.Function.Name
// // 			file := line.Function.Filename
// // 			lineNum := int(line.Line)
// // 			stacktrace[i] = fmt.Sprintf("%s:%s:%d", funcName, file, lineNum)
// // 		}

// // 		sample := Sample{
// // 			ProfileType:  strings.ToLower(p.PeriodType.Type),
// // 			Value:        s.Value[0], // 根据 profile 类型确定具体含义
// // 			Stacktrace:   stacktrace,
// // 			FunctionName: stacktrace[0],
// // 			Timestamp:    time.Unix(0, p.TimeNanos),
// // 			Labels:       parseLabels(s.Label),
// // 		}
// // 		samples = append(samples, sample)
// // 	}
// // 	return samples
// // }

// //	数据过滤
// // 移除低价值样本（如 runtime.mallocgc 等底层函数）
// // 合并相同调用路径的重复样本（减少存储冗余）
// // func isLowValueSample(sample *profile.Sample) bool {
// // 	// 过滤掉 runtime.mallocgc 等底层函数
// // 	for _, loc := range sample.Location {
// // 		for _, line := range loc.Line {
// // 			if strings.HasPrefix(line.Function.Name, "runtime.") {
// // 				return true
// // 			}
// // 		}
// // 	}
// // 	return false
// // }

// // // 判断两个样本是否具有相同的调用路径
// // func isSameCallStack(sample1, sample2 *profile.Sample) bool {
// // 	if len(sample1.Location) != len(sample2.Location) {
// // 		return false
// // 	}
// // 	for i, loc := range sample1.Location {
// // 		if loc.ID != sample2.Location[i].ID {
// // 			return false
// // 		}
// // 	}
// // 	return true
// // }

// // // 判断样本是否为重复样本
// // func isDuplicateSample(sample *profile.Sample, samples []*profile.Sample) bool {
// // 	for _, s := range samples {
// // 		if isSameCallStack(sample, s) {
// // 			return true
// // 		}
// // 	}
// // 	return false
// // }
// // func (s *sAnalysis) filterSamples(samples []*profile.Sample) []*profile.Sample {
// // 	// 移除低价值样本
// // 	var filteredSamples []*profile.Sample
// // 	for _, sample := range samples {
// // 		if isLowValueSample(sample) {
// // 			continue
// // 		}
// // 		filteredSamples = append(filteredSamples, sample)
// // 	}
// // 	// 合并相同调用路径的重复样本
// // 	var mergedSamples []*profile.Sample
// // 	for _, sample := range filteredSamples {
// // 		if isDuplicateSample(sample, mergedSamples) {
// // 			continue
// // 		}
// // 		mergedSamples = append(mergedSamples, sample)
// // 	}
// // 	return mergedSamples
// // }

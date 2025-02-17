package analysis

// import (
// 	"backend/internal/model"
// 	"hash/fnv"
// 	"runtime"
// 	"strings"
// 	"sync"

// 	"github.com/google/pprof/profile"
// )

// // 预编译常用 runtime 函数前缀（减少重复字符串操作）
// var (
// 	// runtimePrefixes = []string{"runtime.", "reflect.", "syscall."}
// 	runtimeMainFunc = "runtime.main"
// )

// // 判断是否低价值样本（支持动态配置）
// func isLowValueSample(sample *profile.Sample, config *model.FilterConfig) bool {
// 	// 按样本值过滤
// 	if len(sample.Value) > 0 && sample.Value[0] < config.MinSampleValue {
// 		return true
// 	}

// 	// 遍历调用栈
// 	for _, loc := range sample.Location {
// 		if loc == nil {
// 			continue
// 		}
// 		for _, line := range loc.Line {
// 			if line.Function == nil {
// 				continue
// 			}
// 			funcName := line.Function.Name

// 			// 保留特定 runtime 函数（如 runtime.main）
// 			if config.KeepRuntime && funcName == runtimeMainFunc {
// 				return false
// 			}

// 			// 前缀匹配过滤
// 			for _, prefix := range config.ExcludeFunctionPrefixes {
// 				if strings.HasPrefix(funcName, prefix) {
// 					return true
// 				}
// 			}

// 			// 正则匹配过滤
// 			if config.ExcludeFunctionRegex != nil &&
// 				config.ExcludeFunctionRegex.MatchString(funcName) {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// // 生成样本的调用栈哈希（用于快速判重）
// func generateStackHash(sample *profile.Sample) uint64 {
// 	h := fnv.New64a()
// 	for _, loc := range sample.Location {
// 		if loc == nil || len(loc.Line) == 0 {
// 			continue
// 		}
// 		// 取顶层调用位置
// 		line := loc.Line[0]
// 		if line.Function != nil {
// 			h.Write([]byte(line.Function.Name))
// 		}
// 		h.Write([]byte{0}) // 分隔符
// 	}
// 	return h.Sum64()
// }

// // 高性能样本过滤与合并
// func (s *sAnalysis) FilterSamplesConcurrent(samples []*profile.Sample, config *model.FilterConfig) []*profile.Sample {
// 	// 预分配内存（减少 append 的扩容次数）
// 	filtered := make([]*profile.Sample, 0, len(samples)/2) // 假设过滤一半
// 	hashMap := make(map[uint64]struct{}, len(samples)/2)

// 	// 并行过滤（利用多核）
// 	var mu sync.Mutex
// 	sem := make(chan struct{}, runtime.NumCPU()) // 并发度控制

// 	for _, sample := range samples {
// 		sem <- struct{}{}
// 		go func(s *profile.Sample) {
// 			defer func() { <-sem }()

// 			// 过滤低价值样本
// 			if isLowValueSample(s, config) {
// 				return
// 			}

// 			// 生成哈希并判重
// 			hash := generateStackHash(s)
// 			mu.Lock()
// 			defer mu.Unlock()
// 			if _, exists := hashMap[hash]; !exists {
// 				hashMap[hash] = struct{}{}
// 				filtered = append(filtered, s)
// 			}
// 		}(sample)
// 	}

// 	// 等待所有 goroutine 完成
// 	for i := 0; i < cap(sem); i++ {
// 		sem <- struct{}{}
// 	}
// 	return filtered
// }

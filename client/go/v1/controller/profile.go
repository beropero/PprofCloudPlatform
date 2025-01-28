package controller

import (
	"bytes"
	"fmt"
	"time"

	"github.com/beropero/PprofCloudPlatform/client/go/v1/capture"
	"github.com/beropero/PprofCloudPlatform/client/go/v1/uploader"
)

type ProfileData struct {
	CPU       *bytes.Buffer `json:"cpu"`
	Memory    *bytes.Buffer `json:"memory"`
	Mutex     *bytes.Buffer `json:"mutex"`
	Block     *bytes.Buffer `json:"block"`
	Goroutine *bytes.Buffer `json:"goroutine"`
}

func (c *Controller) CaptureProfileDataAndUpload(types []string) error {
	data := &ProfileData{} // 初始化所有字段为 nil

	// 去重处理，避免重复捕获相同类型
	uniqueTypes := make(map[string]struct{})
	dedupedTypes := make([]string, 0, len(types))
	for _, t := range types {
		if _, exists := uniqueTypes[t]; exists {
			continue
		}
		uniqueTypes[t] = struct{}{}
		dedupedTypes = append(dedupedTypes, t)
	}

	// 预校验所有类型合法性
	for _, t := range dedupedTypes {
		switch t {
		case "cpu", "memory", "mutex", "block", "goroutine":
			// 合法类型
		default:
			return fmt.Errorf("unknown profile type: %q", t)
		}
	}
	// 捕获所有类型
	for _, t := range dedupedTypes {
		switch t {
		case "cpu":
			cpudata, err := capture.CaptureCPUProfile(5 * time.Second)
			if err != nil {
				return fmt.Errorf("failed to capture CPU profile: %w", err)
			}
			data.CPU = bytes.NewBuffer(cpudata)
		case "memory":
			memdata, err := capture.CaptureMemoryProfile()
			if err != nil {
				return fmt.Errorf("failed to capture memory profile: %w", err)
			}
			data.Memory = bytes.NewBuffer(memdata)
		case "mutex":
			mutexdata, err := capture.CaptureMutexProfile()
			if err != nil {
				return fmt.Errorf("failed to capture mutex profile: %w", err)
			}
			data.Mutex = bytes.NewBuffer(mutexdata)
		case "block":
			blockdata, err := capture.CaptureBlockProfile()
			if err != nil {
				return fmt.Errorf("failed to capture block profile: %w", err)
			}
			data.Block = bytes.NewBuffer(blockdata)
		case "goroutine":
			goroutinedata, err := capture.CaptureGoroutineProfile()
			if err != nil {
				return fmt.Errorf("failed to capture goroutine profile: %w", err)
			}
			data.Goroutine = bytes.NewBuffer(goroutinedata)
		default:
			return fmt.Errorf("unknown profile type: %s", t)
		}
	}
	// 上传数据
	err := uploader.UploadJSON(data, c.Config.UploadUrl)

	return fmt.Errorf("upload failed: %w", err)
}

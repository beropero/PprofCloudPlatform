package controller

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/beropero/PprofCloudPlatform/client/go/v1/capture"
	"github.com/beropero/PprofCloudPlatform/client/go/v1/uploader"
)

type ProfileData struct {
	CPU       []byte `json:"cpu"`
	Memory    []byte `json:"memory"`
	Mutex     []byte `json:"mutex"`
	Block     []byte `json:"block"`
	Goroutine []byte `json:"goroutine"`
}

func (c *Controller) CaptureProfileDataAndUpload(types []string) error {
	// 创建请求体缓冲区
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	// 添加类型作为表单字段
	err := writer.WriteField("types", strings.Join(types, ","))
	if err != nil {
		return fmt.Errorf("failed to write types field: %w", err)
	}
	// 添加服务名称作为表单字段
	if err := writer.WriteField("service", "pprofprofile"); err != nil {
		return fmt.Errorf("failed to write service field: %w", err)
	}
	data := &ProfileData{} // 初始化所有字段为 nil
	fmt.Println("types:", types)
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
	// 捕获指定类型
	for _, t := range dedupedTypes {
		switch t {
		case "cpu":
			cpudata, err := capture.CaptureCPUProfile(5 * time.Second)
			if err != nil {
				return fmt.Errorf("failed to capture CPU profile: %w", err)
			}
			data.CPU = cpudata
			if err := addFilePart(t, data.CPU, writer); err != nil {
				return fmt.Errorf("failed to add CPU profile: %w", err)
			}
		case "memory":
			memdata, err := capture.CaptureMemoryProfile()
			if err != nil {
				return fmt.Errorf("failed to capture memory profile: %w", err)
			}
			data.Memory = memdata
			if err := addFilePart(t, data.Memory, writer); err != nil {
				return fmt.Errorf("failed to add CPU profile: %w", err)
			}
		case "mutex":
			mutexdata, err := capture.CaptureMutexProfile()
			if err != nil {
				return fmt.Errorf("failed to capture mutex profile: %w", err)
			}
			data.Mutex = mutexdata
			if err := addFilePart(t, data.Mutex, writer); err != nil {
				return fmt.Errorf("failed to add CPU profile: %w", err)
			}
		case "block":
			blockdata, err := capture.CaptureBlockProfile()
			if err != nil {
				return fmt.Errorf("failed to capture block profile: %w", err)
			}
			data.Block = blockdata
			if err := addFilePart(t, data.Block, writer); err != nil {
				return fmt.Errorf("failed to add CPU profile: %w", err)
			}
		case "goroutine":
			goroutinedata, err := capture.CaptureGoroutineProfile()
			if err != nil {
				return fmt.Errorf("failed to capture goroutine profile: %w", err)
			}
			data.Goroutine = goroutinedata
			if err := addFilePart(t, data.Goroutine, writer); err != nil {
				return fmt.Errorf("failed to add CPU profile: %w", err)
			}
		default:
			return fmt.Errorf("unknown profile type: %s", t)
		}
	}
	// 上传数据
	err = uploader.UploadFormData(body, writer, c.Config.UploadUrl)
	if err != nil {
		return fmt.Errorf("failed to upload profile data: %w", err)
	}
	return nil
}

// 辅助函数：添加二进制数据到表单
func addFilePart(fieldName string, data []byte, writer *multipart.Writer) error {
	// 使用系统临时目录
	tempDir := os.TempDir()

	// 生成唯一的临时文件名
	tempFile := filepath.Join(tempDir, fmt.Sprintf("%s-%d.pprof", fieldName, time.Now().UnixNano()))

	// 写入临时文件
	if err := os.WriteFile(tempFile, data, 0644); err != nil {
		return fmt.Errorf("写入临时文件失败: %w", err)
	}

	// 确保临时文件最后被删除
	defer os.Remove(tempFile)

	// 创建表单文件字段
	part, err := writer.CreateFormFile(fieldName, filepath.Base(tempFile))
	if err != nil {
		return fmt.Errorf("创建表单文件失败: %w", err)
	}

	// 写入数据
	if _, err := part.Write(data); err != nil {
		return fmt.Errorf("写入表单数据失败: %w", err)
	}

	return nil
}

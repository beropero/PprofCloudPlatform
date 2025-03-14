package controller

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/beropero/PprofCloudPlatform/client/go/v1/capture"
	"github.com/beropero/PprofCloudPlatform/client/go/v1/config"
	"github.com/beropero/PprofCloudPlatform/client/go/v1/uploader"
)

type ProfileInput struct {
	Comment string `json:"comment"`
	Sec     int    `json:"sec"`
	Type    string `json:"type"`
	Gc      bool   `json:"gc"`
}

func (c *Controller) CaptureProfileDataAndUpload(ctx context.Context, in ProfileInput) error {
	// 创建请求体缓冲区
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	// 添加类型作为表单字段
	err := writer.WriteField("type", in.Type)
	if err != nil {
		return fmt.Errorf("failed to write types field: %w", err)
	}
	// 添加服务名称作为表单字段
	if err := writer.WriteField("service", "pprofprofile"); err != nil {
		return fmt.Errorf("failed to write service field: %w", err)
	}

	// 捕获指定类型
	var buf []byte
	var sec = in.Sec
	switch in.Type {
	case "cpu":
		if sec == 0 {
			sec = 5
		}
		buf, err = capture.CaptureCPUProfile(time.Duration(sec) * time.Second)
		if err != nil {
			return fmt.Errorf("failed to capture CPU profile: %w", err)
		}
	case "trace":
		if sec == 0 {
			sec = 5
		}
		buf, err = capture.CaptureTraceProfile(time.Duration(sec) * time.Second)
		if err != nil {
			return fmt.Errorf("failed to capture trace profile: %w", err)
		}
	default:
		c := capture.Capture(in.Type)
		buf, err = c.Collect(ctx, capture.CollectInput{})
		if err != nil {
			return fmt.Errorf("failed to capture delta profile: %w", err)
		}

	}
	// 添加二进制数据到表单
	if err := addFilePart(in.Type, buf, writer); err != nil {
		return fmt.Errorf("failed to add delta profile: %w", err)
	}
	// 添加注释作为表单字段
	if err := writer.WriteField("comment", in.Comment); err != nil {
		return fmt.Errorf("failed to write comment field: %w", err)
	}
	// 上传数据
	err = uploader.UploadFormData(body, writer, config.DefaultConfig().UploadUrl, c.Config)
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
	part, err := writer.CreateFormFile("file", filepath.Base(tempFile))
	if err != nil {
		return fmt.Errorf("创建表单文件失败: %w", err)
	}

	// 写入数据
	if _, err := part.Write(data); err != nil {
		return fmt.Errorf("写入表单数据失败: %w", err)
	}

	return nil
}
